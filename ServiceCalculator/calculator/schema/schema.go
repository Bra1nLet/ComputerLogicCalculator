package schema

type Schema struct {
	Nodes []Node `json:"nodes"`
	Edges []Edge `json:"edges"`
}

type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
}
