package monster

type MonsterDataContainer struct {
	Data MonsterData `json:"data"`
}

type MonsterDataListContainer struct {
	Data []MonsterData `json:"data"`
}

type MonsterData struct {
	Id         string            `json:"id"`
	Type       string            `json:"type"`
	Attributes MonsterAttributes `json:"attributes"`
}

type MonsterAttributes struct {
	Name                string `json:"name"`
	HP                  uint32 `json:"hp"`
	MP                  uint32 `json:"mp"`
	Experience          uint32 `json:"experience"`
	Level               uint32 `json:"level"`
	WeaponAttackDamage  uint32 `json:"paDamage"`
	WeaponDefenseDamage uint32 `json:"pdDamage"`
	MagicAttackDamage   uint32 `json:"maDamage"`
	MagicDefenseDamage  uint32 `json:"mdDamage"`
	Friendly            bool   `json:"friendly"`
	RemoveAfter         uint32 `json:"removeAfter"`
	Boss                bool   `json:"boss"`
	ExplosiveReward     bool   `json:"explosiveReward"`
	FFALoot             bool   `json:"FFALoot"`
	Undead              bool   `json:"undead"`
	BuffToGive          int32  `json:"buffToGive"`
	CarnivalPoint       uint32 `json:"carnivalPoint"`
	RemoveOnMiss        bool   `json:"removeOnMiss"`
	Changeable          bool   `json:"changeable"`
	TagColor            byte   `json:"tagColor"`
	TagBackgroundColor  byte   `json:"TagBackgroundColor"`
	FixedStance         uint32 `json:"fixedStance"`
	FirstAttack         bool   `json:"firstAttack"`
	DropPeriod          uint32 `json:"dropPeriod"`
}
