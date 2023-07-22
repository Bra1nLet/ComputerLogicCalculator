package expressions

import (
	"strconv"
)

func Pow(num int, pow int) int {
	result := num
	for ; pow > 1; pow-- {
		result *= num
	}
	return result
}

func FillNumbers(num string, count int) string {
	result := ""
	for i := 0; i < count; i++ {
		result += num
	}
	return result
}

func IntToBin(num int, count int) string {
	n := strconv.FormatInt(int64(num), 2)
	return FillNumbers("0", count-len(n)) + n
}

func isNegative(i int, oType int) bool {
	if i == 0 && oType == DNF || oType == CNF && i == 1 {
		return true
	} else {
		return false
	}
}

func reverse(i int) int {
	if i == 0 {
		return 1
	} else if i == 1 {
		return 0
	}
	return -1
}

func GenerateGroups(dataType int, data []int, x []string) Group {
	xCount := len(x)
	expression := NewGroup(false)
	for i := 0; i < len(data); i++ {
		if data[i] == dataType {
			express := NewGroup(false)
			for n, j := range IntToBin(i, xCount) {
				d, _ := strconv.Atoi(string(j))
				AddGroupElement(&express, NewElement(x[n], isNegative(d, dataType)))
				if n != xCount-1 {
					AddGroupElement(&express, NewOperation(reverse(dataType)))
				} else {
					if len(expression.Data) != 0 {
						AddGroupElement(&expression, NewOperation(dataType))
					}
					AddGroupElement(&expression, express)
				}
			}
		}
	}
	return expression
}

// deMorganRuleChangeInside this function can convert expression to OrAndNo, AndNoOr, AndOrNo
func deMorganRuleChangeInside(e Group) Group {
	e = *e.DeMorganChange().ToGroup()
	for i, j := range e.Data {
		e.Data[i] = j.DeMorganChange()
	}
	return e
}

// deMorganRuleChangeInside this function can convert expression to AndNo, OrNoAnd, OrNo, AndNoAnd
func deMorganRuleChangeOutside(e Group) Group {
	e = *e.DeMorganChange().ToGroup()
	for i, j := range e.Data {
		e.Data[i] = j.DeMorganChange()
	}
	return e
}

func MakeExpression(testData LogicalTableData) Expression {
	if testData.LogicType == OrAndNo {
		return NewExpression(GenerateGroups(testData.Type, testData.Data, testData.X), testData.Y)
	} else if testData.LogicType == AndNo {
		return NewExpression(deMorganRuleChangeOutside(GenerateGroups(testData.Type, testData.Data, testData.X)), testData.Y)
	} else if testData.LogicType == OrNo {
		return NewExpression(deMorganRuleChangeOutside(GenerateGroups(testData.Type, testData.Data, testData.X)), testData.Y)
	} else {
		return NewExpression(GenerateGroups(testData.Type, testData.Data, testData.X), testData.Y)
	}
}
