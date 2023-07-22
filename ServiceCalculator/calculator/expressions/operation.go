package expressions

type Operation struct {
	Data int `json:"operation"`
}

func NewOperation(data int) Operation {
	return Operation{Data: data}
}

func (o Operation) DeMorganChange() Elements {
	if o.Data == And {
		o.Data = Or
	} else if o.Data == Or {
		o.Data = And
	}
	return o
}

func (o Operation) GetElements() []Elements {
	return nil
}

func (o Operation) IsExpNegative() bool {
	return false
}

func (o Operation) GetData() any {
	return o.Data
}

func (o Operation) IsTheSame(operation Elements) bool {
	return CompareElements(o, operation)
}

func (o Operation) ToGroup() *Group {
	return nil
}
