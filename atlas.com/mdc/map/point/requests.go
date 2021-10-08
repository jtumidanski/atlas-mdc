package point

import (
	"atlas-mdc/monster/drop/position"
	"atlas-mdc/rest/requests"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

const (
	mapInformationServicePrefix string = "/ms/mis/"
	mapInformationService              = requests.BaseRequest + mapInformationServicePrefix
	mapResource                        = mapInformationService + "maps/%d"
	dropPosition                       = mapResource + "/dropPosition"
)

func CalculateDropPosition(l logrus.FieldLogger, span opentracing.Span) func(mapId uint32, initialX int16, initialY int16, fallbackX int16, fallbackY int16) (*MapPointDataContainer, error) {
	return func(mapId uint32, initialX int16, initialY int16, fallbackX int16, fallbackY int16) (*MapPointDataContainer, error) {
		input := &position.DropPositionInputDataContainer{Data: position.DropPositionData{
			Id:   "0",
			Type: "com.atlas.mis.attribute.DropPositionInputAttributes",
			Attributes: position.DropPositionAttributes{
				InitialX:  initialX,
				InitialY:  initialY,
				FallbackX: fallbackX,
				FallbackY: fallbackY,
			},
		}}
		resp, err := requests.Post(l, span)(fmt.Sprintf(dropPosition, mapId), input)
		if err != nil {
			return nil, err
		}

		ar := &MapPointDataContainer{}
		err = requests.ProcessResponse(resp, ar)
		if err != nil {
			return nil, err
		}

		return ar, nil
	}
}
