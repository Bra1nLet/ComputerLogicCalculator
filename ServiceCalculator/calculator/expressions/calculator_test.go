package expressions

import (
	"reflect"
	"testing"
)

func TestFillNumbers(t *testing.T) {
	testDataReqNum := []string{"0", "0", "0", "1", "1"}
	testDataReqC := []int{2, 3, 0, 1, 5}
	testDataRes := []string{"00", "000", "", "1", "11111"}
	lenData := len(testDataReqNum) - 1
	for ; 0 <= lenData; lenData-- {
		result := FillNumbers(testDataReqNum[lenData], testDataReqC[lenData])
		if result != testDataRes[lenData] {
			t.Error("FillNumbers don't working correct.")
			t.Error("Expected:", testDataRes[lenData])
			t.Error("Received:", result)
		}
	}
}

func TestPow(t *testing.T) {
	testDataNum := []int{2, 3, 7, 4, 15, 0}
	testDataPow := []int{3, 4, 6, 2, 3, 1}
	testDataRes := []int{8, 81, 117649, 16, 3375, 0}
	lenData := len(testDataRes) - 1
	for ; 0 <= lenData; lenData-- {
		result := Pow(testDataNum[lenData], testDataPow[lenData])
		if result != testDataRes[lenData] {
			t.Error("FillNumbers don't working correct.")
			t.Error("Expected:", testDataRes[lenData])
			t.Error("Received:", result)
		}
	}
}

func TestIntToBin(t *testing.T) {
	testDataReqNum := []int{0, 1, 22, 14, 8, 7, 10}
	testDataReqC := []int{3, 3, 5, 5, 4, 3, 4}
	testDataRes := []string{"000", "001", "10110", "01110", "1000", "111", "1010"}
	lenData := len(testDataReqNum) - 1
	for ; 0 <= lenData; lenData-- {
		result := IntToBin(testDataReqNum[lenData], testDataReqC[lenData])
		if result != testDataRes[lenData] {
			t.Error("IntToBin don't working correct.")
			t.Log("Data:", testDataReqNum[lenData], testDataReqC[lenData])
			t.Log("Expected:", testDataRes[lenData])
			t.Log("Received:", result)
		}
	}
}

func TestMakeExpression(t *testing.T) {
	var data = []int{1, 0, 0, 1, 0, 1, 0, 1}
	var elements = LimitElements{
		Or:    0,
		OrNo:  0,
		And:   0,
		AndNo: 0,
	}
	var xData = []string{"X1", "X2", "X3"}
	var yData = "Y1"
	var testData1 = []LogicalTableData{
		{Y: yData, X: xData, Data: data, LogicType: OrAndNo, Limits: elements, Type: DNF},
		{Y: yData, X: xData, Data: data, LogicType: OrAndNo, Limits: elements, Type: CNF},
		{Y: yData, X: xData, Data: data, LogicType: AndNo, Limits: elements, Type: DNF},
		{Y: yData, X: xData, Data: data, LogicType: OrNo, Limits: elements, Type: CNF},
		{Y: yData, X: xData, Data: data, LogicType: AndNoOr, Limits: elements, Type: DNF},
		{Y: yData, X: xData, Data: data, LogicType: AndOrNo, Limits: elements, Type: CNF},
		{Y: yData, X: xData, Data: data, LogicType: OrNoOr, Limits: elements, Type: DNF},
		{Y: yData, X: xData, Data: data, LogicType: AndNoAnd, Limits: elements, Type: CNF},
	}
	var testRes = []Expression{TestDNFStepOrAndNo, TestCNFStepOrAndNo, TestDNFStepOrNo, TestCNFStepAndNo, TestDNFStepAndOrNo, TestCNFStepAndNoOr, TestDNFStepOrNoOr, TestCNFStepAndNoAnd}
	for i := 0; i < len(testData1); i++ {
		result := MakeExpression(testData1[i])
		if !testRes[i].CompareExpressions(result) {
			t.Error("Error in test #", i+1)
			t.Error("Expected:", testRes[i])
			t.Error("Received:", result)
		}
	}
}

func TestGroup_DetermineOperation(t *testing.T) {
	var data = Group{
		Data:       []Elements{Element{Data: "X1", IsNegative: false}, Operation{Data: And}, Element{Data: "X2", IsNegative: false}, Operation{Data: And}, Element{Data: "X3", IsNegative: false}},
		IsNegative: false,
	}
	solution := NewOperation(And)
	res := data.DetermineOperation()
	if res != solution {
		t.Error("Failed ")
		t.Error("Expected ", solution)
		t.Error("Received ", res)
	}
}

func TestRuleDeMorgan(t *testing.T) {
	var data = []int{1, 0, 0, 1, 0, 1, 0, 1}
	var elements = LimitElements{
		Or:    0,
		OrNo:  0,
		And:   0,
		AndNo: 0,
	}
	var xData = []string{"X1", "X2", "X3"}
	var yData = "Y1"
	testData := LogicalTableData{Y: yData, X: xData, Data: data, LogicType: OrAndNo, Limits: elements, Type: DNF}
	c := MakeExpression(testData).Data
	c2 := DeMorganRuleChangeInside(DeMorganRuleChange(DeMorganRuleChangeInside(DeMorganRuleChange(*c.ToGroup()))))
	if !reflect.DeepEqual(c, c2) {
		t.Error("Not works")
	}
}

func TestGrouping(t *testing.T) {
	var data = []int{1, 0, 0, 1, 0, 1, 0, 1}
	var elements = LimitElements{
		Or:    2,
		OrNo:  2,
		And:   2,
		AndNo: 2,
	}
	var xData = []string{"X1", "X2", "X3"}
	var yData = "Y1"
	var testData1 = []LogicalTableData{
		{Y: yData, X: xData, Data: data, LogicType: OrAndNo, Limits: elements, Type: DNF},
	}
	var testRes = []Expression{TestGroup}
	for i := 0; i < len(testData1); i++ {
		result := MakeExpression(testData1[i])
		if !testRes[i].CompareExpressions(result) {
			t.Error("Error in test #", i+1)
			t.Error("Expected:", testRes[i])
			t.Error("Received:", result)
		}
	}
}
