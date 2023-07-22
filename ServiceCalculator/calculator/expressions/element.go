package expressions

type Element struct {
	Data       string `json:"element"`
	IsNegative bool   `json:"negative"`
}

func NewElement(data string, isNegative bool) Element {
	return Element{Data: data, IsNegative: isNegative}
}

func (e Element) DeMorganChange() Elements {
	e.IsNegative = !e.IsExpNegative()
	return e
}

func (e Element) GetElements() []Elements {
	return nil
}

func (e Element) IsExpNegative() bool {
	return e.IsNegative
}

func (e Element) GetData() any {
	return e.Data
}

func (e Element) IsTheSame(element Elements) bool {
	return CompareElements(e, element)
}

func (e Element) ToGroup() *Group {
	return nil
}
