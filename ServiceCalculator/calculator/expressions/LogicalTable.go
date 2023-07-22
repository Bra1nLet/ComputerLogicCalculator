package expressions

// LogicalTableData is struct to marshal json data
// Sample of income data
//
//	{
//	  "y": ["Y1"],
//	  "x": ["X1", "X2", "X3"],
//	  "data": [1, 0, 0, 1, 0, 1, 0, 1],
//	  "elements": "AND/OR",
//	  "type": "DNF"
//	}
type LogicalTableData struct {
	Y         string   `json:"y"`
	X         []string `json:"x"`
	Data      []int    `json:"data"`
	Elements  [2]int   `json:"elements"`
	LogicType int      `json:"logicType"`
	Type      int      `json:"type"`
}
