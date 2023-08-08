package expressions

import (
	"errors"
	"reflect"
)

// Group struct is used to group element
type Group struct {
	Data       []Elements `json:"group"`
	IsNegative bool       `json:"negative"`
}

func (g Group) GroupType() (any, error) {
	for _, i := range g.Data {
		if reflect.TypeOf(i.GetData()) == reflect.TypeOf(NewOperation(And)) {
			return i.GetData(), nil
		}
	}
	return nil, errors.New("invalid group")
}

func NewGroup(isNegative bool) Group {
	return Group{Data: []Elements{}, IsNegative: isNegative}
}

func (g Group) DetermineOperation() Elements {
	for i, e := range g.Data {
		if e.ToOperation() != nil {
			return g.Data[i]
		}
	}
	return nil
}

func (g Group) DeMorganChange() Elements {
	g.IsNegative = !g.IsExpNegative()
	return g
}

func AddGroupElement(g *Group, element Elements) {
	g.Data = append(g.Data, element)
}

func (g Group) IsExpNegative() bool {
	return g.IsNegative
}

func (g Group) GetElements() []Elements {
	return g.Data
}

func (g Group) GetData() any {
	return g.Data
}

func (g Group) ToGroup() *Group {
	return &g
}

func (g Group) ToOperation() *Operation {
	return nil
}
