package schema

type InputElement struct {
	Label string `json:"label,omitempty"`
	Type  string `json:"type,omitempty"`
}

func (ie InputElement) GetLogicElement() *LogicElement {
	return nil
}

func (ie InputElement) GetName() string {
	return ie.Label
}

func (ie InputElement) GetType() string {
	return ie.Type
}
