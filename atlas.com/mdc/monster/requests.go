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

func requestById(monsterId uint32) requests.Request[MonsterAttributes] {
	return requests.MakeGetRequest[MonsterAttributes](fmt.Sprintf(monsterResource, monsterId))
}
