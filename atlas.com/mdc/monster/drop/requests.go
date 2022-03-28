package drop

import (
	"atlas-mdc/rest/requests"
	"fmt"
)

const (
	dropInformationServicePrefix string = "/ms/dis/"
	dropInformationService              = requests.BaseRequest + dropInformationServicePrefix
	monsterDropResource                 = dropInformationService + "monsters/drops?monsterId=%d"
)

func requestByMonsterId(monsterId uint32) requests.Request[attributes] {
	return requests.MakeGetRequest[attributes](fmt.Sprintf(monsterDropResource, monsterId))
}
