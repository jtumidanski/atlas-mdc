package drop

import (
	drop2 "atlas-mdc/drop"
	"atlas-mdc/map/point"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"math"
	"math/rand"
	"strconv"
)

func CreateDrops(l logrus.FieldLogger, span opentracing.Span) func(worldId byte, channelId byte, mapId uint32, monsterUniqueId uint32, monsterId uint32, x int16, y int16, killerId uint32) {
	return func(worldId byte, channelId byte, mapId uint32, monsterUniqueId uint32, monsterId uint32, x int16, y int16, killerId uint32) {
		// TODO determine type of drop
		//    monster is explosive? 3
		//    monster has ffa loot? 2
		//    killer is in party? 1
		dropType := byte(0)

		ns, err := GetDropsForMonster(l, span)(monsterId)
		if err != nil {
			return
		}

		l.Debugf("Successfully found %d drops to evaluate.", len(ns))

		ns = getSuccessfulDrops(ns, killerId)

		ids := ""
		for _, n := range ns {
			if len(ids) != 0 {
				ids += ","
			}
			ids += fmt.Sprintf("[%d, %d]", n.ItemId(), n.Chance())
		}
		l.Debugf("Successfully found %d drops to emit. %s", len(ns), ids)

		for i, drop := range ns {
			createDrop(l, span)(worldId, channelId, mapId, i+1, monsterUniqueId, x, y, killerId, dropType, drop)
		}
	}
}

func GetDropsForMonster(l logrus.FieldLogger, span opentracing.Span) func(monsterId uint32) ([]Model, error) {
	return func(monsterId uint32) ([]Model, error) {
		rest, err := requestByMonsterId(l, span)(monsterId)
		if err != nil {
			return nil, err
		}

		ns := make([]Model, 0)
		for _, drop := range rest.DataList() {
			id, err := strconv.ParseUint(drop.Id, 10, 32)
			if err != nil {
				break
			}
			n := makeDrop(uint32(id), drop.Attributes)
			ns = append(ns, n)
		}
		return ns, nil
	}
}

func makeDrop(id uint32, att MonsterDropAttributes) Model {
	return Model{
		monsterId:       att.MonsterId,
		itemId:          att.ItemId,
		minimumQuantity: att.MinimumQuantity,
		maximumQuantity: att.MaximumQuantity,
		chance:          att.Chance,
	}
}

func getSuccessfulDrops(ns []Model, killerId uint32) []Model {
	rs := make([]Model, 0)
	for _, drop := range ns {
		if evaluateSuccess(killerId, drop) {
			rs = append(rs, drop)
		}
	}
	return rs
}

func evaluateSuccess(killerId uint32, drop Model) bool {
	//TODO evaluate rates
	// channel rate
	// buff rate  (cards)

	//TODO evaluate card rate for killer, whether it's meso or drop.
	chance := int32(math.Min(float64(drop.Chance()), math.MaxUint32))
	chance *= 1000
	return rand.Int31n(999999) < chance
}

func createDrop(l logrus.FieldLogger, span opentracing.Span) func(worldId byte, channelId byte, mapId uint32, index int, uniqueId uint32, x int16, y int16, killerId uint32, dropType byte, drop Model) {
	return func(worldId byte, channelId byte, mapId uint32, index int, uniqueId uint32, x int16, y int16, killerId uint32, dropType byte, drop Model) {
		factor := 0
		if dropType == 3 {
			factor = 40
		} else {
			factor = 25
		}
		newX := x
		if index%2 == 0 {
			newX += int16(factor * ((index + 1) / 2))
		} else {
			newX += int16(-(factor * (index / 2)))
		}
		if drop.ItemId() == 0 {
			spawnMeso(l, span)(worldId, channelId, mapId, uniqueId, x, y, killerId, dropType, drop, newX, y)
		} else {
			spawnItem(l, span)(worldId, channelId, mapId, drop.ItemId(), uniqueId, x, y, killerId, dropType, drop, newX, y)
		}
	}
}

func spawnItem(l logrus.FieldLogger, span opentracing.Span) func(worldId byte, channelId byte, mapId uint32, itemId uint32, uniqueId uint32, x int16, y int16, killerId uint32, dropType byte, drop Model, posX int16, posY int16) {
	return func(worldId byte, channelId byte, mapId uint32, itemId uint32, uniqueId uint32, x int16, y int16, killerId uint32, dropType byte, drop Model, posX int16, posY int16) {
		quantity := uint32(1)
		if drop.MaximumQuantity() != 1 {
			quantity = uint32(rand.Int31n(int32(drop.MaximumQuantity()-drop.MinimumQuantity()))) + drop.MinimumQuantity()
		}
		spawnDrop(l, span)(worldId, channelId, mapId, itemId, quantity, 0, posX, posY, x, y, uniqueId, killerId, false, dropType)
	}
}

func spawnMeso(l logrus.FieldLogger, span opentracing.Span) func(worldId byte, channelId byte, mapId uint32, uniqueId uint32, x int16, y int16, killerId uint32, dropType byte, drop Model, posX int16, posY int16) {
	return func(worldId byte, channelId byte, mapId uint32, uniqueId uint32, x int16, y int16, killerId uint32, dropType byte, drop Model, posX int16, posY int16) {
		mesos := uint32(rand.Int31n(int32(drop.MaximumQuantity()-drop.MinimumQuantity()))) + drop.MinimumQuantity()
		//TODO apply characters meso buff.
		mesos *= 20
		spawnDrop(l, span)(worldId, channelId, mapId, 0, 0, mesos, posX, posY, x, y, uniqueId, killerId, false, dropType)
	}
}

func spawnDrop(l logrus.FieldLogger, span opentracing.Span) func(worldId byte, channelId byte, mapId uint32, itemId uint32, quantity uint32, mesos uint32, posX int16, posY int16, monsterX int16, monsterY int16, uniqueId uint32, killerId uint32, playerDrop bool, dropType byte) {
	return func(worldId byte, channelId byte, mapId uint32, itemId uint32, quantity uint32, mesos uint32, posX int16, posY int16, monsterX int16, monsterY int16, uniqueId uint32, killerId uint32, playerDrop bool, dropType byte) {
		tempX, tempY := calculateDropPosition(l, span)(mapId, posX, posY, monsterX, monsterY)
		tempX, tempY = calculateDropPosition(l, span)(mapId, tempX, tempY, tempX, tempY)
		drop2.Spawn(l, span)(worldId, channelId, mapId, itemId, quantity, mesos, dropType, tempX, tempY, killerId, 0, uniqueId, monsterX, monsterY, playerDrop, byte(1))
	}
}

func calculateDropPosition(l logrus.FieldLogger, span opentracing.Span) func(mapId uint32, initialX int16, initialY int16, fallbackX int16, fallbackY int16) (int16, int16) {
	return func(mapId uint32, initialX int16, initialY int16, fallbackX int16, fallbackY int16) (int16, int16) {
		resp, err := point.CalculateDropPosition(l, span)(mapId, initialX, initialY, fallbackX, fallbackY)
		if err != nil {
			return fallbackX, fallbackY
		} else {
			return resp.Data().Attributes.X, resp.Data().Attributes.Y
		}
	}
}
