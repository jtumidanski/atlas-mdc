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

func requestById(monsterId uint32) requests.Request[attributes] {
	return requests.MakeGetRequest[attributes](fmt.Sprintf(monsterResource, monsterId))
}
