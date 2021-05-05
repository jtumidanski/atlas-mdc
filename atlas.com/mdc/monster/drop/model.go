package drop

type Model struct {
	monsterId       uint32
	itemId          uint32
	minimumQuantity uint32
	maximumQuantity uint32
	chance          uint32
}

func (d Model) MonsterId() uint32 {
	return d.monsterId
}

func (d Model) ItemId() uint32 {
	return d.itemId
}

func (d Model) MinimumQuantity() uint32 {
	return d.minimumQuantity
}

func (d Model) MaximumQuantity() uint32 {
	return d.maximumQuantity
}

func (d Model) Chance() uint32 {
	return d.chance
}

type modelBuilder struct {
	monsterId       uint32
	itemId          uint32
	minimumQuantity uint32
	maximumQuantity uint32
	chance          uint32
}

func NewBuilder() *modelBuilder {
	return &modelBuilder{}
}

func (m *modelBuilder) SetMonsterId(monsterId uint32) *modelBuilder {
	m.monsterId = monsterId
	return m
}

func (m *modelBuilder) SetItemId(itemId uint32) *modelBuilder {
	m.itemId = itemId
	return m
}

func (m *modelBuilder) SetMinimumQuantity(minimumQuantity uint32) *modelBuilder {
	m.minimumQuantity = minimumQuantity
	return m
}

func (m *modelBuilder) SetMaximumQuantity(maximumQuantity uint32) *modelBuilder {
	m.maximumQuantity = maximumQuantity
	return m
}

func (m *modelBuilder) SetChance(chance uint32) *modelBuilder {
	m.chance = chance
	return m
}

func (m *modelBuilder) Build() Model {
	return Model{
		monsterId:       m.monsterId,
		itemId:          m.itemId,
		minimumQuantity: m.minimumQuantity,
		maximumQuantity: m.maximumQuantity,
		chance:          m.chance,
	}
}
