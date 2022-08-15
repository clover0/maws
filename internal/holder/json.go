package holder

import "encoding/json"

type jsonMapper struct {
	Name   string         `json:"name"`
	Result map[string]any `json:"result"`
}

type jsonHolder struct {
	buff []jsonMapper
}

func NewJsonHolder() Holder {
	return &jsonHolder{}
}

func (c *jsonHolder) buildMap(s string) map[string]any {
	var m map[string]any
	err := json.Unmarshal([]byte(s), &m)
	if err != nil {
		panic(err)
	}
	return m
}
func (c *jsonHolder) Add(key, elem string) {
	m := jsonMapper{Name: key, Result: c.buildMap(elem)}
	c.buff = append(c.buff, m)
}

func (c *jsonHolder) OutAll() string {
	result, err := json.MarshalIndent(c.buff, "", "    ")
	if err != nil {
		panic(err)
	}
	return string(result)
}
