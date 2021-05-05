package character

type Model struct {
	id    uint32
	hp    uint16
	level byte
	mapId uint32
}

func (a Model) Id() uint32 {
	return a.id
}

func (a Model) HP() uint16 {
	return a.hp
}

func (a Model) Level() byte {
	return a.level
}

func (a Model) MapId() uint32 {
	return a.mapId
}
