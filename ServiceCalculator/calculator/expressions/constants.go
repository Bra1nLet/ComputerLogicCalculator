package expressions

// And, Or - Logical operations
const (
	And = 0
	Or  = 1
)

// Variables which used to declare type of expression
// CNF - Conjunctive Normal Form
// DNF - Disjunctive Normal Form
const (
	CNF = 0
	DNF = 1
)

// Variables which used to choose elements type
// AndNo AndNoOr OrNoOr, these are DNF forms.
// OrNo AndOrNo AndNoAnd, these are CNF forms.
const (
	// OrAndNo This form can be used for CNF and DNF, it's a basic form without transformations.
	OrAndNo = 0

	AndNo   = 1
	AndNoOr = 2
	OrNoOr  = 3

	OrNo     = 4
	AndOrNo  = 5
	AndNoAnd = 6
)
