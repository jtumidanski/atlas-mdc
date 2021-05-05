package point

import (
	"atlas-mdc/rest/response"
)

type MapPointDataContainer struct {
	data     response.DataSegment
	included response.DataSegment
}

type MapPointData struct {
	Id         string             `json:"id"`
	Type       string             `json:"type"`
	Attributes MapPointAttributes `json:"attributes"`
}

type MapPointAttributes struct {
	X int16 `json:"x"`
	Y int16 `json:"y"`
}

func (a *MapPointDataContainer) UnmarshalJSON(data []byte) error {
	d, i, err := response.UnmarshalRoot(data, response.MapperFunc(EmptyMapPointData))
	if err != nil {
		return err
	}

	a.data = d
	a.included = i
	return nil
}

func (a *MapPointDataContainer) Data() *MapPointData {
	if len(a.data) >= 1 {
		return a.data[0].(*MapPointData)
	}
	return nil
}

func (a *MapPointDataContainer) DataList() []MapPointData {
	var r = make([]MapPointData, 0)
	for _, x := range a.data {
		r = append(r, *x.(*MapPointData))
	}
	return r
}

func EmptyMapPointData() interface{} {
	return &MapPointData{}
}
