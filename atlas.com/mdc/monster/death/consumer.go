package death

import (
	"atlas-mdc/kafka/handler"
	"atlas-mdc/monster"
	"atlas-mdc/monster/drop"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type DamageEntry struct {
	CharacterId uint32 `json:"character"`
	Damage      uint64 `json:"damage"`
}

type MonsterKilledEvent struct {
	WorldId       byte          `json:"worldId"`
	ChannelId     byte          `json:"channelId"`
	MapId         uint32        `json:"mapId"`
	UniqueId      uint32        `json:"uniqueId"`
	MonsterId     uint32        `json:"monsterId"`
	X             int16         `json:"x"`
	Y             int16         `json:"y"`
	KillerId      uint32        `json:"killerId"`
	DamageEntries []DamageEntry `json:"damageEntries"`
}

func MonsterKilledEventCreator() handler.EmptyEventCreator {
	return func() interface{} {
		return &MonsterKilledEvent{}
	}
}

func HandleMonsterKilledEvent() handler.EventHandler {
	return func(l logrus.FieldLogger, span opentracing.Span, e interface{}) {
		if event, ok := e.(*MonsterKilledEvent); ok {
			if m, ok := monster.GetMonster(l, span)(event.MonsterId); ok {
				var damageEntries = make([]*monster.DamageEntry, 0)
				for _, entry := range event.DamageEntries {
					damageEntries = append(damageEntries, monster.NewDamageEntry(entry.CharacterId, entry.Damage))
				}

				drop.CreateDrops(l, span)(event.WorldId, event.ChannelId, event.MapId, event.UniqueId, event.MonsterId, event.X, event.Y, event.KillerId)
				monster.DistributeExperience(l, span)(event.WorldId, event.ChannelId, event.MapId, m, damageEntries)
			}
		} else {
			l.Errorf("Unable to cast event provided to handler")
		}
	}
}
