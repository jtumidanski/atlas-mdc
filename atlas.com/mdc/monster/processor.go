package monster

import (
	"atlas-mdc/character"
	"atlas-mdc/configuration"
	"github.com/sirupsen/logrus"
	"math"
)

func GetMonster(l logrus.FieldLogger) func(monsterId uint32) (*Model, bool) {
	return func(monsterId uint32) (*Model, bool) {
		resp, err := requestById(l)(monsterId)
		if err != nil {
			l.WithError(err).Errorf("Retrieving monster %d information.", monsterId)
			return nil, false
		}
		return makeMonster(resp), true
	}
}

func makeMonster(resp *MonsterDataContainer) *Model {
	return &Model{
		experience: resp.Data.Attributes.Experience,
		hp:         resp.Data.Attributes.HP,
	}
}

func DistributeExperience(l logrus.FieldLogger) func(worldId byte, channelId byte, mapId uint32, m *Model, entries []*DamageEntry) {
	return func(worldId byte, channelId byte, mapId uint32, m *Model, entries []*DamageEntry) {
		d := produceDistribution(l)(mapId, m, entries)
		for k, v := range d.Solo() {
			experience := float64(v) * d.ExperiencePerDamage()
			c, err := character.GetCharacterById(l)(k)
			if err != nil {
				l.WithError(err).Errorf("Unable to locate character %d whose for distributing experience from monster death.", k)
			} else {
				whiteExperienceGain := isWhiteExperienceGain(c.Id(), d.PersonalRatio(), d.StandardDeviationRatio())
				distributeCharacterExperience(l)(c.Id(), c.Level(), experience, 0.0, c.Level(), true, whiteExperienceGain, false)
			}
		}
	}
}

type distribution struct {
	solo                   map[uint32]uint64
	party                  map[uint32]map[uint32]uint64
	personalRatio          map[uint32]float64
	experiencePerDamage    float64
	standardDeviationRatio float64
}

func (d distribution) Solo() map[uint32]uint64 {
	return d.solo
}

func (d distribution) ExperiencePerDamage() float64 {
	return d.experiencePerDamage
}

func (d distribution) PersonalRatio() map[uint32]float64 {
	return d.personalRatio
}

func (d distribution) StandardDeviationRatio() float64 {
	return d.standardDeviationRatio
}

func produceDistribution(l logrus.FieldLogger) func(mapId uint32, monster *Model, entries []*DamageEntry) distribution {
	return func(mapId uint32, monster *Model, entries []*DamageEntry) distribution {

		totalEntries := 0
		//TODO incorporate party distribution.
		partyDistribution := make(map[uint32]map[uint32]uint64)
		soloDistribution := make(map[uint32]uint64)

		for _, entry := range entries {
			if character.InMap(l)(entry.CharacterId(), mapId) {
				soloDistribution[entry.CharacterId()] = entry.Damage()
			}
			totalEntries += 1
		}

		//TODO account for healing
		totalDamage := monster.HP()
		epd := float64(monster.Experience()*20) / float64(totalDamage)

		personalRatio := make(map[uint32]float64)
		entryExperienceRatio := make([]float64, 0)

		for k, v := range soloDistribution {
			ratio := float64(v) / float64(totalDamage)
			personalRatio[k] = ratio
			entryExperienceRatio = append(entryExperienceRatio, ratio)
		}

		for _, party := range partyDistribution {
			ratio := 0.0
			for k, v := range party {
				cr := float64(v) / float64(totalDamage)
				personalRatio[k] = cr
				ratio += cr
			}
			entryExperienceRatio = append(entryExperienceRatio, ratio)
		}

		stdr := calculateExperienceStandardDeviationThreshold(entryExperienceRatio, totalEntries)
		return distribution{
			solo:                   soloDistribution,
			party:                  partyDistribution,
			personalRatio:          personalRatio,
			experiencePerDamage:    epd,
			standardDeviationRatio: stdr,
		}
	}
}

func isWhiteExperienceGain(characterId uint32, personalRatio map[uint32]float64, standardDeviationRatio float64) bool {
	if val, ok := personalRatio[characterId]; ok {
		return val >= standardDeviationRatio
	} else {
		return false
	}
}

func calculateExperienceStandardDeviationThreshold(entryExperienceRatio []float64, totalEntries int) float64 {
	averageExperienceReward := 0.0
	for _, v := range entryExperienceRatio {
		averageExperienceReward += v
	}
	averageExperienceReward /= float64(totalEntries)

	varExperienceReward := 0.0
	for _, v := range entryExperienceRatio {
		varExperienceReward += math.Pow(v-averageExperienceReward, 2)
	}
	varExperienceReward /= float64(len(entryExperienceRatio))

	return averageExperienceReward + math.Sqrt(varExperienceReward)
}

func distributeCharacterExperience(l logrus.FieldLogger) func(characterId uint32, level byte, experience float64, partyBonusMod float64, totalPartyLevel byte, hightestPartyDamage bool, whiteExperienceGain bool, hasPartySharers bool) {
	return func(characterId uint32, level byte, experience float64, partyBonusMod float64, totalPartyLevel byte, hightestPartyDamage bool, whiteExperienceGain bool, hasPartySharers bool) {
		expSplitCommonMod := configuration.Get().ExpSplitCommonMod
		characterExperience := (float64(expSplitCommonMod) * float64(level)) / float64(totalPartyLevel)
		if hightestPartyDamage {
			characterExperience += float64(configuration.Get().ExpSplitMvpMod)
		}
		characterExperience *= experience
		bonusExperience := partyBonusMod * characterExperience

		giveExperienceToCharacter(l)(characterId, characterExperience, bonusExperience, whiteExperienceGain, hasPartySharers)
	}
}

func giveExperienceToCharacter(l logrus.FieldLogger) func(characterId uint32, experience float64, bonus float64, whiteExperienceGain bool, hasPartySharers bool) {
	return func(characterId uint32, experience float64, bonus float64, whiteExperienceGain bool, hasPartySharers bool) {
		correctedPersonal := experienceValueToInteger(experience)
		correctedParty := experienceValueToInteger(bonus)
		character.GiveExperience(l)(characterId, correctedPersonal, correctedParty, true, false, whiteExperienceGain)
	}
}

func experienceValueToInteger(experience float64) uint32 {
	if experience > math.MaxInt32 {
		experience = math.MaxInt32
	} else if experience < math.MinInt32 {
		experience = math.MinInt32
	}
	return uint32(math.RoundToEven(experience))
}
