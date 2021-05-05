package death

import (
	"atlas-mdc/kafka/handler"
	"atlas-mdc/monster/drop"
	"github.com/sirupsen/logrus"
)

type DamageEntry struct {
	CharacterId uint32 `json:"characterId"`
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
	return func(l logrus.FieldLogger, e interface{}) {
		if event, ok := e.(*MonsterKilledEvent); ok {
			drop.CreateDrops(l)(event.WorldId, event.ChannelId, event.MapId, event.UniqueId, event.MonsterId, event.X, event.Y, event.KillerId)
		} else {
		}
	}
}
