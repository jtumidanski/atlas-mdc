package drop

import (
	"atlas-mdc/kafka"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type command struct {
	WorldId      byte   `json:"worldId"`
	ChannelId    byte   `json:"channelId"`
	MapId        uint32 `json:"mapId"`
	ItemId       uint32 `json:"itemId"`
	Quantity     uint32 `json:"quantity"`
	Mesos        uint32 `json:"mesos"`
	DropType     byte   `json:"dropType"`
	X            int16  `json:"x"`
	Y            int16  `json:"y"`
	OwnerId      uint32 `json:"ownerId"`
	OwnerPartyId uint32 `json:"ownerPartyId"`
	DropperId    uint32 `json:"dropperId"`
	DropperX     int16  `json:"dropperX"`
	DropperY     int16  `json:"dropperY"`
	PlayerDrop   bool   `json:"playerDrop"`
	Mod          byte   `json:"mod"`
}

func Spawn(l logrus.FieldLogger, span opentracing.Span) func(worldId byte, channelId byte, mapId uint32, itemId uint32, quantity uint32,
	mesos uint32, dropType byte, x int16, y int16, ownerId uint32, ownerPartyId uint32, dropperId uint32,
	dropperX int16, dropperY int16, playerDrop bool, mod byte) {
	producer := kafka.ProduceEvent(l, span, "TOPIC_SPAWN_DROP_COMMAND")
	return func(worldId byte, channelId byte, mapId uint32, itemId uint32, quantity uint32, mesos uint32, dropType byte,
		x int16, y int16, ownerId uint32, ownerPartyId uint32, dropperId uint32, dropperX int16, dropperY int16,
		playerDrop bool, mod byte) {
		c := &command{
			WorldId:      worldId,
			ChannelId:    channelId,
			MapId:        mapId,
			ItemId:       itemId,
			Quantity:     quantity,
			Mesos:        mesos,
			DropType:     dropType,
			X:            x,
			Y:            y,
			OwnerId:      ownerId,
			OwnerPartyId: ownerPartyId,
			DropperId:    dropperId,
			DropperX:     dropperX,
			DropperY:     dropperY,
			PlayerDrop:   playerDrop,
			Mod:          mod,
		}
		producer(kafka.CreateKey(0), c)
	}
}
