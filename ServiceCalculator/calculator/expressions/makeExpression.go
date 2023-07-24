package expressions

import (
	"log"
	"reflect"
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

// DeMorganRuleChangeInside this function can convert expression to OrAndNo, AndNoOr, AndOrNo
func DeMorganRuleChangeInside(e Group) Group {
	for i, j := range e.Data {
		if reflect.TypeOf(j) == reflect.TypeOf(NewGroup(false)) {
			e.Data[i] = DeMorganRuleChangeOutside(*j.ToGroup())
		}
	}
	return e
}

// GroupExp method to make groups depends at
func GroupExp(g Group, data LogicalTableData) Group {
	if len(data.Elements) == 0 {
		return g
	}
	group := NewGroup(false)
	for i, _ := range g.Data {
		if reflect.TypeOf(g.Data[i]) == reflect.TypeOf(group) {
			g.Data[i] = GroupExp(*g.Data[i].ToGroup(), data)
		}
	}
	i := data.Elements[0]
	n := NewGroup(false)
	log.Println((i + i - 1), len(g.Data))
	if len(g.Data) > (i + i - 1) {
		group.Data = g.Data[i+i-1 : len(g.Data)]
		n.Data = g.Data[0:(i + i - 1)]
		g.Data = []Elements{}
		g.Data = append(g.Data, n)
		log.Println(g.Data)
		g.Data = append(g.Data, group.Data...)
	}
	if len(g.Data) > (i + i - 1) {
		return GroupExp(g, data)
	}

	return g
}

// DeMorganRuleChangeOutside this function can convert expression to AndNo, OrNoAnd, OrNo, AndNoAnd
func DeMorganRuleChangeOutside(e Group) Group {
	e = *e.DeMorganChange().ToGroup()
	for i, j := range e.Data {
		e.Data[i] = j.DeMorganChange()
	}
	return e
}

func MakeExpression(testData LogicalTableData) Expression {
	if testData.LogicType == OrAndNo {
		e := GenerateGroups(testData.Type, testData.Data, testData.X)
		return NewExpression(GroupExp(e, testData), testData.Y)
	} else if testData.LogicType == AndNo || testData.LogicType == OrNo {
		return NewExpression(DeMorganRuleChangeOutside(GenerateGroups(testData.Type, testData.Data, testData.X)), testData.Y)
	} else if testData.LogicType == AndNoOr || testData.LogicType == AndOrNo {
		exp := DeMorganRuleChangeOutside(GenerateGroups(testData.Type, testData.Data, testData.X))
		return NewExpression(DeMorganRuleChangeInside(exp), testData.Y)
	} else if testData.LogicType == OrNoOr || testData.LogicType == AndNoAnd {
		exp := DeMorganRuleChangeInside(DeMorganRuleChangeOutside(GenerateGroups(testData.Type, testData.Data, testData.X)))
		return NewExpression(DeMorganRuleChangeOutside(exp), testData.Y)
	} else {
		return NewExpression(GenerateGroups(testData.Type, testData.Data, testData.X), testData.Y)
	}
}
