package drop

import (
	"atlas-mdc/rest/requests"
	"fmt"
	"github.com/sirupsen/logrus"
)

const (
	dropInformationServicePrefix string = "/ms/dis/"
	dropInformationService              = requests.BaseRequest + dropInformationServicePrefix
	monsterDropResource                 = dropInformationService + "monsters/drops?monsterId=%d"
)

func requestByMonsterId(l logrus.FieldLogger) func(monsterId uint32) (*MonsterDropDataContainer, error) {
	return func(monsterId uint32) (*MonsterDropDataContainer, error) {
		ar := &MonsterDropDataContainer{}
		err := requests.Get(l)(fmt.Sprintf(monsterDropResource, monsterId), ar)
		if err != nil {
			return nil, err
		}
		return ar, nil
	}
}
