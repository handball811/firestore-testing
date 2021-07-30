package format

import (
	"fmt"
	"reflect"
)

var _ Formatter = new(DocRef)

type Map struct {
	Prefix string
	Target map[string]interface{}

	list []struct {
		key   string
		value interface{}
	}
	index int
}

func (m *Map) Next() (string, error) {
	if m.index == 0 {
		m.list = make([]struct {
			key   string
			value interface{}
		}, 0, len(m.Target))
		for k, v := range m.Target {
			m.list = append(m.list, struct {
				key   string
				value interface{}
			}{
				key:   k,
				value: v,
			})
		}
	}
	i := m.index
	if len(m.list) == i {
		return "", ErrFin
	}
	resp := fmt.Sprintf(
		"%s%v: (%v -> %v)",
		m.Prefix,
		m.list[i].key,
		getType(m.list[i].value),
		m.list[i].value)
	m.index = i + 1
	return resp, nil
}

func getType(t interface{}) string {
	if t == nil {
		return "Nil"
	}
	return reflect.TypeOf(t).String()
}
