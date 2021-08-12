package character

import (
	"errors"
	"github.com/sirupsen/logrus"
	"strconv"
)

func GetCharacterById(l logrus.FieldLogger) func(characterId uint32) (*Model, error) {
	return func(characterId uint32) (*Model, error) {
		cs, err := requestCharacter(l)(characterId)
		if err != nil {
			return nil, err
		}
		ca := makeCharacterAttributes(cs.Data())
		if ca == nil {
			return nil, errors.New("unable to make character attributes")
		}
		return ca, nil
	}
}

func InMap(l logrus.FieldLogger) func(characterId uint32, mapId uint32) bool {
	return func(characterId uint32, mapId uint32) bool {
		c, err := GetCharacterById(l)(characterId)
		if err != nil {
			l.WithError(err).Errorf("Unable to retrieve character for map check, assuming false.")
			return false
		}
		return c.MapId() == mapId
	}
}

func makeCharacterAttributes(ca *dataBody) *Model {
	cid, err := strconv.ParseUint(ca.Id, 10, 32)
	if err != nil {
		return nil
	}
	att := ca.Attributes
	r := Model{
		id:    uint32(cid),
		hp:    att.Hp,
		level: att.Level,
		mapId: att.MapId,
	}
	return &r
}
