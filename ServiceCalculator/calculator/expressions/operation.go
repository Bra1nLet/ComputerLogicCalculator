package expressions

type Operation struct {
	Data int `json:"operation"`
}

func NewOperation(data int) Operation {
	return Operation{Data: data}
}

func (o Operation) DeMorganChange() Elements {
	switch o.Data {
	case And:
		o.Data = Or
	case Or:
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

func (o Operation) ToGroup() *Group {
	return nil
}

func (o Operation) ToOperation() *Operation {
	return &o
}
