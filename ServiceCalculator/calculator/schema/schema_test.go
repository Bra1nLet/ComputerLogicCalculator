package schema

import (
	exp "MicroservicesPet/ServiceCalculator/calculator/expressions"
	"reflect"
	"testing"
)

func TestExpToSchema(t *testing.T) {
	var data = []int{1, 0, 0, 1, 0, 1, 0, 1}
	var elements = exp.LimitElements{
		Or:    0,
		OrNo:  0,
		And:   0,
		AndNo: 0,
	}
	var xData = []string{"X1", "X2", "X3"}
	var yData = "Y1"
	if !SchemaIsSame(ExpToSchema(exp.MakeExpression(exp.LogicalTableData{Y: yData, X: xData, Data: data, LogicType: exp.OrAndNo, Limits: elements, Type: exp.DNF})), testTest) {
		t.Error("Dont Working")
	}
}

func SchemaIsSame(schema1, schema2 Schema) bool {
	return reflect.DeepEqual(schema1, schema2)
}

var testTest = Schema{
	Nodes: []Node{
		{Id: "1", Position: Position{X: 0, Y: 0}, Element: LogicElement{NodeName: "&", Ports: 3, NoElements: []int{0, 1, 2}}, Type: "LogicNode"},
		{Id: "2", Position: Position{X: 0, Y: 200}, Element: LogicElement{NodeName: "&", Ports: 3, NoElements: []int{0}}, Type: "LogicNode"},
		{Id: "3", Position: Position{X: 0, Y: 400}, Element: LogicElement{NodeName: "&", Ports: 3, NoElements: []int{1}}, Type: "LogicNode"},
		{Id: "4", Position: Position{X: 0, Y: 600}, Element: LogicElement{NodeName: "&", Ports: 3}, Type: "LogicNode"},
		{Id: "5", Position: Position{X: 200, Y: 100}, Element: LogicElement{NodeName: "1", Ports: 2}, Type: "LogicNode"},
		{Id: "6", Position: Position{X: 200, Y: 500}, Element: LogicElement{NodeName: "1", Ports: 2}, Type: "LogicNode"},
		{Id: "7", Position: Position{X: 400, Y: 300}, Element: LogicElement{NodeName: "1", Ports: 2}, Type: "LogicNode"},
		{Id: "8", Position: Position{X: -100, Y: 40}, Element: InputElement{Label: "X1", Type: "source"}, Type: "VariableNode"},
		{Id: "9", Position: Position{X: -100, Y: 60}, Element: InputElement{Label: "X2", Type: "source"}, Type: "VariableNode"},
		{Id: "10", Position: Position{X: -100, Y: 80}, Element: InputElement{Label: "X3", Type: "source"}, Type: "VariableNode"},
		{Id: "11", Position: Position{X: -100, Y: 240}, Element: InputElement{Label: "X1", Type: "source"}, Type: "VariableNode"},
		{Id: "12", Position: Position{X: -100, Y: 260}, Element: InputElement{Label: "X2", Type: "source"}, Type: "VariableNode"},
		{Id: "13", Position: Position{X: -100, Y: 280}, Element: InputElement{Label: "X3", Type: "source"}, Type: "VariableNode"},
		{Id: "14", Position: Position{X: -100, Y: 440}, Element: InputElement{Label: "X1", Type: "source"}, Type: "VariableNode"},
		{Id: "15", Position: Position{X: -100, Y: 460}, Element: InputElement{Label: "X2", Type: "source"}, Type: "VariableNode"},
		{Id: "16", Position: Position{X: -100, Y: 480}, Element: InputElement{Label: "X3", Type: "source"}, Type: "VariableNode"},
		{Id: "17", Position: Position{X: -100, Y: 640}, Element: InputElement{Label: "X1", Type: "source"}, Type: "VariableNode"},
		{Id: "18", Position: Position{X: -100, Y: 660}, Element: InputElement{Label: "X2", Type: "source"}, Type: "VariableNode"},
		{Id: "19", Position: Position{X: -100, Y: 680}, Element: InputElement{Label: "X3", Type: "source"}, Type: "VariableNode"},
		{Id: "20", Position: Position{X: 600, Y: 350}, Element: InputElement{Label: "Y1", Type: "target"}, Type: "VariableNode"},
	},
	Edges: []Edge{
		{Id: "1", Source: "8", Target: "1", TargetHandle: "0", SourceHandle: "0"},
		{Id: "2", Source: "9", Target: "1", TargetHandle: "1", SourceHandle: "0"},
		{Id: "3", Source: "10", Target: "1", TargetHandle: "2", SourceHandle: "0"},
		{Id: "4", Source: "11", Target: "2", TargetHandle: "0", SourceHandle: "0"},
		{Id: "5", Source: "12", Target: "2", TargetHandle: "1", SourceHandle: "0"},
		{Id: "6", Source: "13", Target: "2", TargetHandle: "2", SourceHandle: "0"},
		{Id: "7", Source: "14", Target: "3", TargetHandle: "0", SourceHandle: "0"},
		{Id: "8", Source: "15", Target: "3", TargetHandle: "1", SourceHandle: "0"},
		{Id: "9", Source: "16", Target: "3", TargetHandle: "2", SourceHandle: "0"},
		{Id: "10", Source: "17", Target: "4", TargetHandle: "0", SourceHandle: "0"},
		{Id: "11", Source: "18", Target: "4", TargetHandle: "1", SourceHandle: "0"},
		{Id: "12", Source: "19", Target: "4", TargetHandle: "2", SourceHandle: "0"},

		{Id: "13", Source: "1", Target: "5", TargetHandle: "0", SourceHandle: "0"},
		{Id: "14", Source: "2", Target: "5", TargetHandle: "1", SourceHandle: "0"},
		{Id: "15", Source: "3", Target: "6", TargetHandle: "0", SourceHandle: "0"},
		{Id: "16", Source: "4", Target: "6", TargetHandle: "1", SourceHandle: "0"},

		{Id: "17", Source: "5", Target: "7", TargetHandle: "0", SourceHandle: "0"},
		{Id: "18", Source: "6", Target: "7", TargetHandle: "1", SourceHandle: "0"},

		{Id: "19", Source: "7", Target: "20", TargetHandle: "0", SourceHandle: "0"},
	},
}
