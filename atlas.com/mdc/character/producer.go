package character

import (
	"atlas-mdc/kafka"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type characterExperienceGainEvent struct {
	CharacterId  uint32 `json:"characterId"`
	PersonalGain uint32 `json:"personalGain"`
	PartyGain    uint32 `json:"partyGain"`
	Show         bool   `json:"show"`
	Chat         bool   `json:"chat"`
	White        bool   `json:"white"`
}

func GiveExperience(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32, personalGain uint32, partyGain uint32, show bool, chat bool, white bool) {
	producer := kafka.ProduceEvent(l, span, "TOPIC_CHARACTER_EXPERIENCE_EVENT")
	return func(characterId uint32, personalGain uint32, partyGain uint32, show bool, chat bool, white bool) {
		event := &characterExperienceGainEvent{characterId, personalGain, partyGain, show, chat, white}
		producer(kafka.CreateKey(int(characterId)), event)
	}
}
