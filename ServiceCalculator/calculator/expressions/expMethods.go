package expressions

import (
	"reflect"
)

func IsTypesMatch(element1 any, element2 any) bool {
	return reflect.TypeOf(element1) == reflect.TypeOf(element2)
}

func CompareGroups(element1 Elements, element2 Elements) bool {
	if len(element1.GetElements()) != len(element2.GetElements()) {
		return false
	}
	for i := range element1.GetElements() {
		if !CompareElements(element1.GetElements()[i], element2.GetElements()[i]) {
			return false
		}
	}
	return true
}

func CompareElements(element1 Elements, element2 Elements) bool {
	if !IsTypesMatch(element1, element2) {
		return false
	}
	if element1.IsExpNegative() != element2.IsExpNegative() {
		return false
	}
	if IsTypesMatch(element1, Group{}) {
		return CompareGroups(element1, element2)
	} else {
		if element1.GetData() != element2.GetData() {
			return false
		}
	}
	return true
}
