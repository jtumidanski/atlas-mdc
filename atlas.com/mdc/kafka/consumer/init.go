package consumer

import (
	"atlas-mdc/kafka/handler"
	"atlas-mdc/monster/death"
	"github.com/sirupsen/logrus"
)

func CreateEventConsumers(l *logrus.Logger) {
	cec := func(topicToken string, emptyEventCreator handler.EmptyEventCreator, processor handler.EventHandler) {
		createEventConsumer(l, topicToken, emptyEventCreator, processor)
	}
	cec("TOPIC_MONSTER_KILLED_EVENT", death.MonsterKilledEventCreator(), death.HandleMonsterKilledEvent())
}

func createEventConsumer(l *logrus.Logger, topicToken string, emptyEventCreator handler.EmptyEventCreator, processor handler.EventHandler) {
	go NewConsumer(l, topicToken, "Monster Death Coordinator", emptyEventCreator, processor)
}
