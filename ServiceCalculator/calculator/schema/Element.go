package schema

type Element interface {
	GetLogicElement() *LogicElement
	GetName() string
	GetType() string
}
