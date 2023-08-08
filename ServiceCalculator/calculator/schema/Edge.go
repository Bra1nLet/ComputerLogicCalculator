package schema

type Edge struct {
	Id           string `json:"id"`
	Source       string `json:"source"`
	Target       string `json:"target"`
	TargetHandle string `json:"targetHandle"`
	SourceHandle string `json:"sourceHandle"`
}
