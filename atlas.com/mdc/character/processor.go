package character

import (
	"atlas-mdc/model"
	"atlas-mdc/rest/requests"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"strconv"
)

func ByIdModelProvider(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32) model.Provider[Model] {
	return func(characterId uint32) model.Provider[Model] {
		return requests.Provider[attributes, Model](l, span)(requestCharacter(characterId), makeModel)
	}
}

func GetById(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32) (Model, error) {
	return func(characterId uint32) (Model, error) {
		return ByIdModelProvider(l, span)(characterId)()
	}
}

func InMap(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32, mapId uint32) bool {
	return func(characterId uint32, mapId uint32) bool {
		c, err := GetById(l, span)(characterId)
		if err != nil {
			l.WithError(err).Errorf("Unable to retrieve character for map check, assuming false.")
			return false
		}
		return c.MapId() == mapId
	}
}

func makeModel(ca requests.DataBody[attributes]) (Model, error) {
	cid, err := strconv.ParseUint(ca.Id, 10, 32)
	if err != nil {
		return Model{}, err
	}
	att := ca.Attributes
	r := Model{
		id:    uint32(cid),
		hp:    att.Hp,
		level: att.Level,
		mapId: att.MapId,
	}
	return r, nil
}
