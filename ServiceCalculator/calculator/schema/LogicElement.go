package schema

type LogicElement struct {
	NodeName   string `json:"nodeName,omitempty"`
	Ports      int    `json:"ports,omitempty"`
	NoElements []int  `json:"noElements,omitempty"`
	Type       string `json:"type,omitempty"`
}

func (le LogicElement) GetLogicElement() *LogicElement {
	return &le
}

func (le LogicElement) GetName() string {
	return le.NodeName
}

func (le LogicElement) GetType() string {
	return le.Type
}
