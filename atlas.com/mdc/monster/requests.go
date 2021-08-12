package monster

import (
	"atlas-mdc/rest/requests"
	"fmt"
	"github.com/sirupsen/logrus"
)

const (
	monsterServicePrefix string = "/ms/mis/"
	monsterService              = requests.BaseRequest + monsterServicePrefix
	monstersResource            = monsterService + "monsters"
	monsterResource             = monstersResource + "/%d"
)

func requestById(l logrus.FieldLogger) func(monsterId uint32) (*MonsterDataContainer, error) {
	return func(monsterId uint32) (*MonsterDataContainer, error) {
		ar := &MonsterDataContainer{}
		err := requests.Get(l)(fmt.Sprintf(monsterResource, monsterId), ar)
		if err != nil {
			return nil, err
		}
		return ar, nil
	}
}
