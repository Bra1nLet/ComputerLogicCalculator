package schema

type Node struct {
	Id       string   `json:"id"`
	Position Position `json:"position"`
	Element  Element  `json:"data"`
	Type     string   `json:"type"`
}

func (n Node) Same(node Node) bool {
	if node.Id != n.Id || node.Position != n.Position || n.Type != node.Type {
		return false
	}

	return true
}
