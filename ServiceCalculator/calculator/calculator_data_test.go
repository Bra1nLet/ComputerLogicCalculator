package calculator

import exp "MicroservicesPet/ServiceCalculator/calculator/expressions"

// TestCNFStepOrAndNo is a Data to test first step of program
var TestCNFStepOrAndNo = exp.Expression{
	Data: exp.Group{
		Data: []exp.Elements{
			exp.Group{
				Data: []exp.Elements{
					exp.Element{Data: "X1", IsNegative: false},
					exp.Operation{Data: exp.Or},
					exp.Element{Data: "X2", IsNegative: false},
					exp.Operation{Data: exp.Or},
					exp.Element{Data: "X3", IsNegative: true},
				},
				IsNegative: false,
			},
			exp.Operation{Data: exp.And},
			exp.Group{
				Data: []exp.Elements{
					exp.Element{Data: "X1", IsNegative: false},
					exp.Operation{Data: exp.Or},
					exp.Element{Data: "X2", IsNegative: true},
					exp.Operation{Data: exp.Or},
					exp.Element{Data: "X3", IsNegative: false},
				},
				IsNegative: false,
			},
			exp.Operation{Data: exp.And},
			exp.Group{
				Data: []exp.Elements{
					exp.Element{Data: "X1", IsNegative: true},
					exp.Operation{Data: exp.Or},
					exp.Element{Data: "X2", IsNegative: false},
					exp.Operation{Data: exp.Or},
					exp.Element{Data: "X3", IsNegative: false},
				},
				IsNegative: false,
			},
			exp.Operation{Data: exp.And},
			exp.Group{
				Data: []exp.Elements{
					exp.Element{Data: "X1", IsNegative: true},
					exp.Operation{Data: exp.Or},
					exp.Element{Data: "X2", IsNegative: true},
					exp.Operation{Data: exp.Or},
					exp.Element{Data: "X3", IsNegative: false},
				},
				IsNegative: false,
			},
		},
		IsNegative: false,
	},
	Value: "Y1",
}

// TestDNFStepOrAndNo is a Data to test first step of program
var TestDNFStepOrAndNo = exp.Expression{
	Data: exp.Group{
		Data: []exp.Elements{
			exp.Group{
				Data: []exp.Elements{
					exp.Element{Data: "X1", IsNegative: true},
					exp.Operation{Data: exp.And},
					exp.Element{Data: "X2", IsNegative: true},
					exp.Operation{Data: exp.And},
					exp.Element{Data: "X3", IsNegative: true},
				},
				IsNegative: false,
			},
			exp.Operation{Data: exp.Or},
			exp.Group{
				Data: []exp.Elements{
					exp.Element{Data: "X1", IsNegative: true},
					exp.Operation{Data: exp.And},
					exp.Element{Data: "X2", IsNegative: false},
					exp.Operation{Data: exp.And},
					exp.Element{Data: "X3", IsNegative: false},
				},
				IsNegative: false,
			},
			exp.Operation{Data: exp.Or},
			exp.Group{
				Data: []exp.Elements{
					exp.Element{Data: "X1", IsNegative: false},
					exp.Operation{Data: exp.And},
					exp.Element{Data: "X2", IsNegative: true},
					exp.Operation{Data: exp.And},
					exp.Element{Data: "X3", IsNegative: false},
				},
				IsNegative: false,
			},
			exp.Operation{Data: exp.Or},
			exp.Group{
				Data: []exp.Elements{
					exp.Element{Data: "X1", IsNegative: false},
					exp.Operation{Data: exp.And},
					exp.Element{Data: "X2", IsNegative: false},
					exp.Operation{Data: exp.And},
					exp.Element{Data: "X3", IsNegative: false},
				},
				IsNegative: false,
			},
		},
		IsNegative: false,
	},
	Value: "Y1",
}

// TestCNFStepAndNo is a Data for testing  rule of De Morgan
var TestCNFStepAndNo = exp.Expression{
	Data: exp.Group{
		Data: []exp.Elements{
			exp.Group{
				Data: []exp.Elements{
					exp.Element{Data: "X1", IsNegative: false},
					exp.Operation{Data: exp.Or},
					exp.Element{Data: "X2", IsNegative: false},
					exp.Operation{Data: exp.Or},
					exp.Element{Data: "X3", IsNegative: true},
				},
				IsNegative: true,
			},
			exp.Operation{Data: exp.Or},
			exp.Group{
				Data: []exp.Elements{
					exp.Element{Data: "X1", IsNegative: false},
					exp.Operation{Data: exp.Or},
					exp.Element{Data: "X2", IsNegative: true},
					exp.Operation{Data: exp.Or},
					exp.Element{Data: "X3", IsNegative: false},
				},
				IsNegative: true,
			},
			exp.Operation{Data: exp.Or},
			exp.Group{
				Data: []exp.Elements{
					exp.Element{Data: "X1", IsNegative: true},
					exp.Operation{Data: exp.Or},
					exp.Element{Data: "X2", IsNegative: false},
					exp.Operation{Data: exp.Or},
					exp.Element{Data: "X3", IsNegative: false},
				},
				IsNegative: true,
			},
			exp.Operation{Data: exp.Or},
			exp.Group{
				Data: []exp.Elements{
					exp.Element{Data: "X1", IsNegative: true},
					exp.Operation{Data: exp.Or},
					exp.Element{Data: "X2", IsNegative: true},
					exp.Operation{Data: exp.Or},
					exp.Element{Data: "X3", IsNegative: false},
				},
				IsNegative: true,
			},
		},
		IsNegative: true,
	},
	Value: "Y1",
}

// TestDNFStepOrNo is a Data for testing  rule of De Morgan
var TestDNFStepOrNo = exp.Expression{
	Data: exp.Group{
		Data: []exp.Elements{
			exp.Group{
				Data: []exp.Elements{
					exp.Element{Data: "X1", IsNegative: true},
					exp.Operation{Data: exp.And},
					exp.Element{Data: "X2", IsNegative: true},
					exp.Operation{Data: exp.And},
					exp.Element{Data: "X3", IsNegative: true},
				},
				IsNegative: true,
			},
			exp.Operation{Data: exp.And},
			exp.Group{
				Data: []exp.Elements{
					exp.Element{Data: "X1", IsNegative: true},
					exp.Operation{Data: exp.And},
					exp.Element{Data: "X2", IsNegative: false},
					exp.Operation{Data: exp.And},
					exp.Element{Data: "X3", IsNegative: false},
				},
				IsNegative: true,
			},
			exp.Operation{Data: exp.And},
			exp.Group{
				Data: []exp.Elements{
					exp.Element{Data: "X1", IsNegative: false},
					exp.Operation{Data: exp.And},
					exp.Element{Data: "X2", IsNegative: true},
					exp.Operation{Data: exp.And},
					exp.Element{Data: "X3", IsNegative: false},
				},
				IsNegative: true,
			},
			exp.Operation{Data: exp.And},
			exp.Group{
				Data: []exp.Elements{
					exp.Element{Data: "X1", IsNegative: false},
					exp.Operation{Data: exp.And},
					exp.Element{Data: "X2", IsNegative: false},
					exp.Operation{Data: exp.And},
					exp.Element{Data: "X3", IsNegative: false},
				},
				IsNegative: true,
			},
		},
		IsNegative: true,
	},
	Value: "Y1",
}

// TestDNFStepAndOrNo is a Data for testing  rule of De Morgan
var TestDNFStepAndOrNo = exp.Expression{
	Data: exp.Group{
		Data: []exp.Elements{
			exp.Group{
				Data: []exp.Elements{
					exp.Element{Data: "X1", IsNegative: false},
					exp.Operation{Data: exp.Or},
					exp.Element{Data: "X2", IsNegative: false},
					exp.Operation{Data: exp.Or},
					exp.Element{Data: "X3", IsNegative: false},
				},
				IsNegative: false,
			},
			exp.Operation{Data: exp.And},
			exp.Group{
				Data: []exp.Elements{
					exp.Element{Data: "X1", IsNegative: false},
					exp.Operation{Data: exp.Or},
					exp.Element{Data: "X2", IsNegative: true},
					exp.Operation{Data: exp.Or},
					exp.Element{Data: "X3", IsNegative: true},
				},
				IsNegative: false,
			},
			exp.Operation{Data: exp.And},
			exp.Group{
				Data: []exp.Elements{
					exp.Element{Data: "X1", IsNegative: true},
					exp.Operation{Data: exp.Or},
					exp.Element{Data: "X2", IsNegative: false},
					exp.Operation{Data: exp.Or},
					exp.Element{Data: "X3", IsNegative: true},
				},
				IsNegative: false,
			},
			exp.Operation{Data: exp.And},
			exp.Group{
				Data: []exp.Elements{
					exp.Element{Data: "X1", IsNegative: true},
					exp.Operation{Data: exp.Or},
					exp.Element{Data: "X2", IsNegative: true},
					exp.Operation{Data: exp.Or},
					exp.Element{Data: "X3", IsNegative: true},
				},
				IsNegative: false,
			},
		},
		IsNegative: true,
	},
	Value: "Y1",
}

// TestCNFStepAndNoOr is a Data for testing  rule of De Morgan
var TestCNFStepAndNoOr = exp.Expression{
	Data: exp.Group{
		Data: []exp.Elements{
			exp.Group{
				Data: []exp.Elements{
					exp.Element{Data: "X1", IsNegative: true},
					exp.Operation{Data: exp.And},
					exp.Element{Data: "X2", IsNegative: true},
					exp.Operation{Data: exp.And},
					exp.Element{Data: "X3", IsNegative: false},
				},
				IsNegative: false,
			},
			exp.Operation{Data: exp.Or},
			exp.Group{
				Data: []exp.Elements{
					exp.Element{Data: "X1", IsNegative: true},
					exp.Operation{Data: exp.And},
					exp.Element{Data: "X2", IsNegative: false},
					exp.Operation{Data: exp.And},
					exp.Element{Data: "X3", IsNegative: true},
				},
				IsNegative: false,
			},
			exp.Operation{Data: exp.Or},
			exp.Group{
				Data: []exp.Elements{
					exp.Element{Data: "X1", IsNegative: false},
					exp.Operation{Data: exp.And},
					exp.Element{Data: "X2", IsNegative: true},
					exp.Operation{Data: exp.And},
					exp.Element{Data: "X3", IsNegative: true},
				},
				IsNegative: false,
			},
			exp.Operation{Data: exp.Or},
			exp.Group{
				Data: []exp.Elements{
					exp.Element{Data: "X1", IsNegative: false},
					exp.Operation{Data: exp.And},
					exp.Element{Data: "X2", IsNegative: false},
					exp.Operation{Data: exp.And},
					exp.Element{Data: "X3", IsNegative: true},
				},
				IsNegative: false,
			},
		},
		IsNegative: true,
	},
	Value: "Y1",
}

// TestDNFStepOrNoOr is a Data for testing  rule of De Morgan
var TestDNFStepOrNoOr = exp.Expression{
	Data: exp.Group{
		Data: []exp.Elements{
			exp.Group{
				Data: []exp.Elements{
					exp.Element{Data: "X1", IsNegative: false},
					exp.Operation{Data: exp.Or},
					exp.Element{Data: "X2", IsNegative: false},
					exp.Operation{Data: exp.Or},
					exp.Element{Data: "X3", IsNegative: false},
				},
				IsNegative: true,
			},
			exp.Operation{Data: exp.Or},
			exp.Group{
				Data: []exp.Elements{
					exp.Element{Data: "X1", IsNegative: false},
					exp.Operation{Data: exp.Or},
					exp.Element{Data: "X2", IsNegative: true},
					exp.Operation{Data: exp.Or},
					exp.Element{Data: "X3", IsNegative: true},
				},
				IsNegative: true,
			},
			exp.Operation{Data: exp.Or},
			exp.Group{
				Data: []exp.Elements{
					exp.Element{Data: "X1", IsNegative: true},
					exp.Operation{Data: exp.Or},
					exp.Element{Data: "X2", IsNegative: false},
					exp.Operation{Data: exp.Or},
					exp.Element{Data: "X3", IsNegative: true},
				},
				IsNegative: true,
			},
			exp.Operation{Data: exp.Or},
			exp.Group{
				Data: []exp.Elements{
					exp.Element{Data: "X1", IsNegative: true},
					exp.Operation{Data: exp.Or},
					exp.Element{Data: "X2", IsNegative: true},
					exp.Operation{Data: exp.Or},
					exp.Element{Data: "X3", IsNegative: true},
				},
				IsNegative: true,
			},
		},
		IsNegative: false,
	},
	Value: "Y1",
}

// TestCNFStepAndNoAnd is a Data for testing  rule of De Morgan
var TestCNFStepAndNoAnd = exp.Expression{
	Data: exp.Group{
		Data: []exp.Elements{
			exp.Group{
				Data: []exp.Elements{
					exp.Element{Data: "X1", IsNegative: true},
					exp.Operation{Data: exp.And},
					exp.Element{Data: "X2", IsNegative: true},
					exp.Operation{Data: exp.And},
					exp.Element{Data: "X3", IsNegative: false},
				},
				IsNegative: true,
			},
			exp.Operation{Data: exp.And},
			exp.Group{
				Data: []exp.Elements{
					exp.Element{Data: "X1", IsNegative: true},
					exp.Operation{Data: exp.And},
					exp.Element{Data: "X2", IsNegative: false},
					exp.Operation{Data: exp.And},
					exp.Element{Data: "X3", IsNegative: true},
				},
				IsNegative: true,
			},
			exp.Operation{Data: exp.And},
			exp.Group{
				Data: []exp.Elements{
					exp.Element{Data: "X1", IsNegative: false},
					exp.Operation{Data: exp.And},
					exp.Element{Data: "X2", IsNegative: true},
					exp.Operation{Data: exp.And},
					exp.Element{Data: "X3", IsNegative: true},
				},
				IsNegative: true,
			},
			exp.Operation{Data: exp.And},
			exp.Group{
				Data: []exp.Elements{
					exp.Element{Data: "X1", IsNegative: false},
					exp.Operation{Data: exp.And},
					exp.Element{Data: "X2", IsNegative: false},
					exp.Operation{Data: exp.And},
					exp.Element{Data: "X3", IsNegative: true},
				},
				IsNegative: true,
			},
		},
		IsNegative: false,
	},
	Value: "Y1",
}

// TestGroup is a Data for testing grouping expression with elements 2&, 2 or
var TestGroup = exp.Expression{
	Data: exp.Group{
		Data: []exp.Elements{
			exp.Group{
				Data: []exp.Elements{
					exp.Group{
						Data: []exp.Elements{
							exp.Group{
								Data: []exp.Elements{
									exp.Group{
										Data: []exp.Elements{
											exp.Element{Data: "X1", IsNegative: true},
											exp.Operation{Data: exp.And},
											exp.Element{Data: "X2", IsNegative: true},
										},
										IsNegative: false,
									},
									exp.Operation{Data: exp.And},
									exp.Element{Data: "X3", IsNegative: true},
								},
								IsNegative: false,
							},
							exp.Operation{Data: exp.Or},
							exp.Group{
								Data: []exp.Elements{
									exp.Group{
										Data: []exp.Elements{
											exp.Element{Data: "X1", IsNegative: true},
											exp.Operation{Data: exp.And},
											exp.Element{Data: "X2", IsNegative: false},
										},
										IsNegative: false,
									},
									exp.Operation{Data: exp.And},
									exp.Element{Data: "X3", IsNegative: false},
								},
								IsNegative: false,
							},
						},
						IsNegative: false,
					},
					exp.Operation{Data: exp.Or},
					exp.Group{
						Data: []exp.Elements{
							exp.Group{
								Data: []exp.Elements{
									exp.Element{Data: "X1", IsNegative: false},
									exp.Operation{Data: exp.And},
									exp.Element{Data: "X2", IsNegative: true},
								},
								IsNegative: false,
							},
							exp.Operation{Data: exp.And},
							exp.Element{Data: "X3", IsNegative: false},
						},
						IsNegative: false,
					},
				},
				IsNegative: false,
			},
			exp.Operation{Data: exp.Or},
			exp.Group{
				Data: []exp.Elements{
					exp.Group{
						Data: []exp.Elements{
							exp.Element{Data: "X1", IsNegative: false},
							exp.Operation{Data: exp.And},
							exp.Element{Data: "X2", IsNegative: false},
						},
						IsNegative: false,
					},
					exp.Operation{Data: exp.And},
					exp.Element{Data: "X3", IsNegative: false},
				},
				IsNegative: false,
			},
		},
		IsNegative: false,
	},
	Value: "Y1",
}
