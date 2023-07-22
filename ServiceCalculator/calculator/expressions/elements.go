package expressions

type Elements interface {
	IsExpNegative() bool
	IsTheSame(group Elements) bool
	GetData() any
	GetElements() []Elements
	DeMorganChange() Elements
	ToGroup() *Group
}
