package expressions

// LogicalTableData is struct to marshal json data
// Sample of income data
//
//	{
//	  "y": ["Y1"],
//	  "x": ["X1", "X2", "X3"],
//	  "data": [1, 0, 0, 1, 0, 1, 0, 1],
//	  "limits": 0,
//	  "type": 1,
//	}
type LogicalTableData struct {
	Y         string        `json:"y"`
	X         []string      `json:"x"`
	Data      []int         `json:"data"`
	Limits    LimitElements `json:"limits"`
	LogicType int           `json:"logicType"`
	Type      int           `json:"type"`
}

// LimitElements structure, which can control size of elements.
// If 0, elements will choose automatically
type LimitElements struct {
	Or    int `json:"or"`
	OrNo  int `json:"orNo"`
	And   int `json:"and"`
	AndNo int `json:"andNo"`
}
