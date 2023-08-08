package expressions

import (
	"reflect"
	"strconv"
)

const (
	and   = 0
	or    = 1
	andNo = 2
	orNo  = 3
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
			e.Data[i] = DeMorganRuleChange(*j.ToGroup())
		}
	}
	return e
}

// ChoseLimitElements Can choose automatically number for LimitElements
func ChoseLimitElements(g int) int {
	switch g {
	case 3:
		return 2
	case 5:
		return 3
	case 7:
		return 4
	}
	return 8
}

func findOperation(o Operation, g Group) int {
	if g.IsNegative {
		return o.Data + 2
	}
	return o.Data
}

func notZero(n, data int) int {
	if n == 0 {
		return data
	}
	return n
}

func findLimit(g Group, limits LimitElements) int {
	limit := ChoseLimitElements(len(g.Data))
	d := g.DetermineOperation()
	if d != nil {
		switch findOperation(*d.ToOperation(), g) {
		case and:
			limit = notZero(limits.And, limit)
		case andNo:
			limit = notZero(limits.AndNo, limit)
		case or:
			limit = notZero(limits.Or, limit)
		case orNo:
			limit = notZero(limits.OrNo, limit)
		}
	}
	return limit
}

// MakeGroup method to make groups depends at LimitElements
func MakeGroup(g Group, limits LimitElements) Group {
	limit := findLimit(g, limits)

	for i, e := range g.Data {
		if e.ToGroup() != nil {
			g.Data[i] = MakeGroup(*e.ToGroup(), limits)
		}
	}

	group(&g, limit+limit-1)
	return g
}

func group(g *Group, limit int) {
	n := NewGroup(false)
	gr := NewGroup(false)
	if len(g.Data) > (limit) {
		gr.Data = g.Data[limit:len(g.Data)]
		n.Data = g.Data[0:(limit)]
		g.Data = []Elements{}
		g.Data = append(g.Data, n)
		g.Data = append(g.Data, gr.Data...)
		if len(g.Data) > (limit) {
			group(g, limit)
		}
	}
}

// DeMorganRuleChange this function can convert expression to AndNo, OrNoAnd, OrNo, AndNoAnd format
func DeMorganRuleChange(e Group) Group {
	e = *e.DeMorganChange().ToGroup()
	for i, j := range e.Data {
		e.Data[i] = j.DeMorganChange()
	}
	return e
}

func MakeExpression(testData LogicalTableData) Expression {
	data := NewGroup(false)
	switch testData.LogicType {
	case OrAndNo:
		data = GenerateGroups(testData.Type, testData.Data, testData.X)
	case AndNo, OrNo:
		data = DeMorganRuleChange(GenerateGroups(testData.Type, testData.Data, testData.X))
	case AndNoOr, AndOrNo:
		exp := DeMorganRuleChange(GenerateGroups(testData.Type, testData.Data, testData.X))
		data = DeMorganRuleChangeInside(exp)
	case OrNoOr, AndNoAnd:
		exp := DeMorganRuleChangeInside(DeMorganRuleChange(GenerateGroups(testData.Type, testData.Data, testData.X)))
		data = DeMorganRuleChange(exp)
	}
	return NewExpression(MakeGroup(data, testData.Limits), testData.Y)
}
