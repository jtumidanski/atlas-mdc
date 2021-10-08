package consumers

import (
	"atlas-mdc/kafka/handler"
	"atlas-mdc/monster/death"
	"context"
	"github.com/sirupsen/logrus"
	"sync"
)

const MonsterKilledEvent = "monster_killed_event"

func CreateEventConsumers(l *logrus.Logger, ctx context.Context, wg *sync.WaitGroup) {
	cec := func(topicToken string, name string, emptyEventCreator handler.EmptyEventCreator, processor handler.EventHandler) {
		createEventConsumer(l, ctx, wg, name, topicToken, emptyEventCreator, processor)
	}
	cec("TOPIC_MONSTER_KILLED_EVENT", MonsterKilledEvent, death.MonsterKilledEventCreator(), death.HandleMonsterKilledEvent())
}

func createEventConsumer(l *logrus.Logger, ctx context.Context, wg *sync.WaitGroup, name string, topicToken string, emptyEventCreator handler.EmptyEventCreator, processor handler.EventHandler) {
	wg.Add(1)
	go NewConsumer(l, ctx, wg, name, topicToken, "Monster Death Coordinator", emptyEventCreator, processor)
}
