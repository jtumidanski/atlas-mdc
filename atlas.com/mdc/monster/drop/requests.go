package drop

import (
	"atlas-mdc/rest/requests"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

const (
	dropInformationServicePrefix string = "/ms/dis/"
	dropInformationService              = requests.BaseRequest + dropInformationServicePrefix
	monsterDropResource                 = dropInformationService + "monsters/drops?monsterId=%d"
)

func requestByMonsterId(l logrus.FieldLogger, span opentracing.Span) func(monsterId uint32) (*MonsterDropDataContainer, error) {
	return func(monsterId uint32) (*MonsterDropDataContainer, error) {
		ar := &MonsterDropDataContainer{}
		err := requests.Get(l, span)(fmt.Sprintf(monsterDropResource, monsterId), ar)
		if err != nil {
			return nil, err
		}
		return ar, nil
	}
}
