package monster

import (
	"atlas-mdc/rest/requests"
	"fmt"
)

const (
	monsterServicePrefix string = "/ms/mis/"
	monsterService              = requests.BaseRequest + monsterServicePrefix
	monstersResource            = monsterService + "monsters"
	monsterResource             = monstersResource + "/%d"
)

var Monster = func() *monster {
	return &monster{}
}

type monster struct {
}

func (m *monster) GetById(monsterId uint32) (*MonsterDataContainer, error) {
	ar := &MonsterDataContainer{}
	err := requests.Get(fmt.Sprintf(monsterResource, monsterId), ar)
	if err != nil {
		return nil, err
	}
	return ar, nil
}
