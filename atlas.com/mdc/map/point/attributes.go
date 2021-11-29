package point

import (
	"atlas-mdc/rest/response"
	"encoding/json"
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

func (c *MapPointDataContainer) MarshalJSON() ([]byte, error) {
	t := struct {
		Data     interface{} `json:"data"`
		Included interface{} `json:"included"`
	}{}
	if len(c.data) == 1 {
		t.Data = c.data[0]
	} else {
		t.Data = c.data
	}
	return json.Marshal(t)
}

func (c *MapPointDataContainer) UnmarshalJSON(data []byte) error {
	d, i, err := response.UnmarshalRoot(data, response.MapperFunc(EmptyMapPointData))
	if err != nil {
		return err
	}

	c.data = d
	c.included = i
	return nil
}

func (c *MapPointDataContainer) Data() *MapPointData {
	if len(c.data) >= 1 {
		return c.data[0].(*MapPointData)
	}
	return nil
}

func (c *MapPointDataContainer) DataList() []MapPointData {
	var r = make([]MapPointData, 0)
	for _, x := range c.data {
		r = append(r, *x.(*MapPointData))
	}
	return r
}

func EmptyMapPointData() interface{} {
	return &MapPointData{}
}
