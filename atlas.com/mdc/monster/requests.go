package monster

import (
	"atlas-mdc/rest/requests"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

const (
	monsterServicePrefix string = "/ms/mis/"
	monsterService              = requests.BaseRequest + monsterServicePrefix
	monstersResource            = monsterService + "monsters"
	monsterResource             = monstersResource + "/%d"
)

func requestById(l logrus.FieldLogger, span opentracing.Span) func(monsterId uint32) (*MonsterDataContainer, error) {
	return func(monsterId uint32) (*MonsterDataContainer, error) {
		ar := &MonsterDataContainer{}
		err := requests.Get(l, span)(fmt.Sprintf(monsterResource, monsterId), ar)
		if err != nil {
			return nil, err
		}
		return ar, nil
	}
}
