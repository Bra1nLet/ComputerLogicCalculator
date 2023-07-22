package expressions

// Expression use to make logical expressions, like
// Y1 = -(X1 + X2) * (X1 + (-X2))
type Expression struct {
	Data  Elements `json:"elements"`
	Value string   `json:"value"`
}

func NewExpression(data Elements, value string) Expression {
	return Expression{Data: data, Value: value}
}

func (e Expression) CompareExpressions(expression Expression) bool {
	if e.Value != expression.Value {
		return false
	}
	return CompareElements(e.Data, expression.Data)
}
