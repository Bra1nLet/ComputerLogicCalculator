package expressions

// Group struct is used to group element
type Group struct {
	Data       []Elements `json:"group"`
	IsNegative bool       `json:"negative"`
}

func NewGroup(isNegative bool) Group {
	return Group{Data: []Elements{}, IsNegative: isNegative}
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

func (g Group) IsTheSame(group Elements) bool {
	return CompareElements(g, group)
}

func (g Group) ToGroup() *Group {
	return &g
}
