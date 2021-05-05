package monster

type DamageEntry struct {
	characterId uint32
	damage      uint64
}

func (e DamageEntry) CharacterId() uint32 {
	return e.characterId
}

func (e DamageEntry) Damage() uint64 {
	return e.damage
}

func NewDamageEntry(characterId uint32, damage uint64) *DamageEntry {
	return &DamageEntry{
		characterId: characterId,
		damage:      damage,
	}
}

type Model struct {
	experience uint32
	hp         uint32
}

func (m Model) HP() uint32 {
	return m.hp
}

func (m Model) Experience() uint32 {
	return m.experience
}
