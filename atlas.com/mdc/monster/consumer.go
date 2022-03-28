package monster

import (
	"atlas-mdc/kafka"
	"atlas-mdc/monster/drop"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

const (
	consumerNameKilled = "monster_killed_event"
	topicTokenKilled   = "TOPIC_MONSTER_KILLED_EVENT"
)

func DeathConsumer(groupId string) kafka.ConsumerConfig {
	return kafka.NewConsumerConfig[killedEvent](consumerNameKilled, topicTokenKilled, groupId, HandleMonsterKilledEvent())
}

type damageEntry struct {
	CharacterId uint32 `json:"character"`
	Damage      uint64 `json:"damage"`
}

type killedEvent struct {
	WorldId       byte          `json:"worldId"`
	ChannelId     byte          `json:"channelId"`
	MapId         uint32        `json:"mapId"`
	UniqueId      uint32        `json:"uniqueId"`
	MonsterId     uint32        `json:"monsterId"`
	X             int16         `json:"x"`
	Y             int16         `json:"y"`
	KillerId      uint32        `json:"killerId"`
	DamageEntries []damageEntry `json:"damageEntries"`
}

func HandleMonsterKilledEvent() kafka.HandlerFunc[killedEvent] {
	return func(l logrus.FieldLogger, span opentracing.Span, event killedEvent) {
		if m, ok := GetMonster(l, span)(event.MonsterId); ok {
			var damageEntries = make([]*DamageEntry, 0)
			for _, entry := range event.DamageEntries {
				damageEntries = append(damageEntries, NewDamageEntry(entry.CharacterId, entry.Damage))
			}

			drop.CreateDrops(l, span)(event.WorldId, event.ChannelId, event.MapId, event.UniqueId, event.MonsterId, event.X, event.Y, event.KillerId)
			DistributeExperience(l, span)(event.WorldId, event.ChannelId, event.MapId, m, damageEntries)
		}
	}
}
