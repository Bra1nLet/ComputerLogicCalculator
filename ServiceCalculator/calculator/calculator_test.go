package calculator

import (
	exp "MicroservicesPet/ServiceCalculator/calculator/expressions"
	"testing"
)

func TestIsGroupsSame(t *testing.T) {
	testDataReqOp1 := []exp.Group{{Data: []exp.Elements{exp.Operation{Data: 0}, exp.Operation{Data: 1}}, IsNegative: false}, {Data: []exp.Elements{exp.Operation{Data: 1}, exp.Operation{Data: 0}}, IsNegative: true}}
	testDataReqOp2 := []exp.Group{{Data: []exp.Elements{exp.Operation{Data: 0}, exp.Operation{Data: 1}}, IsNegative: false}, {Data: []exp.Elements{exp.Operation{Data: 1}, exp.Operation{Data: 1}}, IsNegative: true}}
	testDataRes := []bool{true, false}
	lenData := len(testDataRes) - 1
	for ; 0 <= lenData; lenData-- {
		result := testDataReqOp1[lenData].IsTheSame(testDataReqOp2[lenData])
		if result != testDataRes[lenData] {
			t.Error("IsOperationsSame don't working correct.")
			t.Error("Expected:", testDataRes[lenData])
			t.Error("Received:", result)
		}
	}
}

func TestIsOperationsSame(t *testing.T) {
	testDataReqOp1 := []exp.Elements{exp.Operation{Data: 1}, exp.Operation{Data: 0}, exp.Operation{Data: 1}, exp.Operation{Data: 0}, exp.Element{Data: "X1", IsNegative: true}, exp.Operation{Data: 0}}
	testDataReqOp2 := []exp.Elements{exp.Operation{Data: 1}, exp.Operation{Data: 0}, exp.Operation{Data: 0}, exp.Operation{Data: 1}, exp.Operation{Data: 0}, exp.Element{Data: "X1", IsNegative: true}}
	testDataRes := []bool{true, true, false, false, false, false}
	lenData := len(testDataRes) - 1
	for ; 0 <= lenData; lenData-- {
		result := testDataReqOp1[lenData].IsTheSame(testDataReqOp2[lenData])
		if result != testDataRes[lenData] {
			t.Error("IsOperationsSame don't working correct.")
			t.Error("Expected:", testDataRes[lenData])
			t.Error("Received:", result)
		}
	}
}

func TestIsElementsSame(t *testing.T) {
	testDataReqE1 := []exp.Element{{Data: "X1", IsNegative: true}, {Data: "X1", IsNegative: true}, {Data: "X1", IsNegative: true}, {Data: "X1", IsNegative: false}, {Data: "X2", IsNegative: false}}
	testDataReqE2 := []exp.Elements{exp.Element{Data: "X1", IsNegative: true}, exp.Element{Data: "X1", IsNegative: false}, exp.Group{Data: []exp.Elements{}, IsNegative: false}, exp.Operation{Data: 1}, exp.Element{Data: "X2", IsNegative: false}}
	testDataRes := []bool{true, false, false, false, true}
	lenData := len(testDataRes) - 1
	for ; 0 <= lenData; lenData-- {
		result := testDataReqE1[lenData].IsTheSame(testDataReqE2[lenData])
		if result != testDataRes[lenData] {
			t.Error("IsOperationsSame don't working correct.")
			t.Error("Expected:", testDataRes[lenData])
			t.Error("Received:", result)
		}
	}
}

func TestFillNumbers(t *testing.T) {
	testDataReqNum := []string{"0", "0", "0", "1", "1"}
	testDataReqC := []int{2, 3, 0, 1, 5}
	testDataRes := []string{"00", "000", "", "1", "11111"}
	lenData := len(testDataReqNum) - 1
	for ; 0 <= lenData; lenData-- {
		result := exp.FillNumbers(testDataReqNum[lenData], testDataReqC[lenData])
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
		result := exp.Pow(testDataNum[lenData], testDataPow[lenData])
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
		result := exp.IntToBin(testDataReqNum[lenData], testDataReqC[lenData])
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
	var elements = [2]int{3, 2}
	var xData = []string{"X1", "X2", "X3"}
	var yData = "Y1"
	var testData1 = []exp.LogicalTableData{
		{Y: yData, X: xData, Data: data, LogicType: exp.OrAndNo, Elements: elements, Type: exp.DNF},
		{Y: yData, X: xData, Data: data, LogicType: exp.OrAndNo, Elements: elements, Type: exp.CNF},
		{Y: yData, X: xData, Data: data, LogicType: exp.AndNo, Elements: elements, Type: exp.DNF},
		{Y: yData, X: xData, Data: data, LogicType: exp.OrNo, Elements: elements, Type: exp.CNF}}
	var testRes = []exp.Expression{TestDNFStepOrAndNo, TestCNFStepOrAndNo, TestDNFStepOrNo, TestCNFStepAndNo}
	for i := 0; i < len(testData1); i++ {
		result := exp.MakeExpression(testData1[i])
		if !testRes[i].CompareExpressions(result) {
			t.Error("Error !")
			t.Error("Expected:", testRes[i])
			t.Error("Received:", result)
		}
	}
}
