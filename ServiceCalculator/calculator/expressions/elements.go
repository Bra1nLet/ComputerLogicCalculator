package expressions

type Elements interface {
	IsExpNegative() bool
	GetData() any
	GetElements() []Elements
	DeMorganChange() Elements
	ToGroup() *Group
	ToOperation() *Operation
}
