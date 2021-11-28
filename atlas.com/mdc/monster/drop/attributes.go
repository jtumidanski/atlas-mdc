package drop

import (
	"atlas-mdc/rest/response"
	"encoding/json"
)

type MonsterDropDataContainer struct {
	data     response.DataSegment
	included response.DataSegment
}

type MonsterDropData struct {
	Id         string                `json:"id"`
	Type       string                `json:"type"`
	Attributes MonsterDropAttributes `json:"attributes"`
}

type MonsterDropAttributes struct {
	MonsterId       uint32 `json:"monsterId"`
	ItemId          uint32 `json:"itemId"`
	MaximumQuantity uint32 `json:"maximumQuantity"`
	MinimumQuantity uint32 `json:"minimumQuantity"`
	Chance          uint32 `json:"chance"`
}

func (c *MonsterDropDataContainer) MarshalJSON() ([]byte, error) {
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

func (c *MonsterDropDataContainer) UnmarshalJSON(data []byte) error {
	d, i, err := response.UnmarshalRoot(data, response.MapperFunc(EmptyMonsterDropData))
	if err != nil {
		return err
	}

	c.data = d
	c.included = i
	return nil
}

func (c *MonsterDropDataContainer) Data() *MonsterDropData {
	if len(c.data) >= 1 {
		return c.data[0].(*MonsterDropData)
	}
	return nil
}

func (c *MonsterDropDataContainer) DataList() []MonsterDropData {
	var r = make([]MonsterDropData, 0)
	for _, x := range c.data {
		r = append(r, *x.(*MonsterDropData))
	}
	return r
}

func EmptyMonsterDropData() interface{} {
	return &MonsterDropData{}
}
