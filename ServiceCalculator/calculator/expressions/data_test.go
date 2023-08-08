package expressions

// TestCNFStepOrAndNo is a Data to test first step of program
var TestCNFStepOrAndNo = Expression{
	Data: Group{
		Data: []Elements{
			Group{
				Data: []Elements{
					Element{Data: "X1", IsNegative: false},
					Operation{Data: Or},
					Element{Data: "X2", IsNegative: false},
					Operation{Data: Or},
					Element{Data: "X3", IsNegative: true},
				},
				IsNegative: false,
			},
			Operation{Data: And},
			Group{
				Data: []Elements{
					Element{Data: "X1", IsNegative: false},
					Operation{Data: Or},
					Element{Data: "X2", IsNegative: true},
					Operation{Data: Or},
					Element{Data: "X3", IsNegative: false},
				},
				IsNegative: false,
			},
			Operation{Data: And},
			Group{
				Data: []Elements{
					Element{Data: "X1", IsNegative: true},
					Operation{Data: Or},
					Element{Data: "X2", IsNegative: false},
					Operation{Data: Or},
					Element{Data: "X3", IsNegative: false},
				},
				IsNegative: false,
			},
			Operation{Data: And},
			Group{
				Data: []Elements{
					Element{Data: "X1", IsNegative: true},
					Operation{Data: Or},
					Element{Data: "X2", IsNegative: true},
					Operation{Data: Or},
					Element{Data: "X3", IsNegative: false},
				},
				IsNegative: false,
			},
		},
		IsNegative: false,
	},
	Value: "Y1",
}

// TestDNFStepOrAndNo is a Data to test first step of program
var TestDNFStepOrAndNo = Expression{
	Data: Group{
		Data: []Elements{
			Group{
				Data: []Elements{
					Element{Data: "X1", IsNegative: true},
					Operation{Data: And},
					Element{Data: "X2", IsNegative: true},
					Operation{Data: And},
					Element{Data: "X3", IsNegative: true},
				},
				IsNegative: false,
			},
			Operation{Data: Or},
			Group{
				Data: []Elements{
					Element{Data: "X1", IsNegative: true},
					Operation{Data: And},
					Element{Data: "X2", IsNegative: false},
					Operation{Data: And},
					Element{Data: "X3", IsNegative: false},
				},
				IsNegative: false,
			},
			Operation{Data: Or},
			Group{
				Data: []Elements{
					Element{Data: "X1", IsNegative: false},
					Operation{Data: And},
					Element{Data: "X2", IsNegative: true},
					Operation{Data: And},
					Element{Data: "X3", IsNegative: false},
				},
				IsNegative: false,
			},
			Operation{Data: Or},
			Group{
				Data: []Elements{
					Element{Data: "X1", IsNegative: false},
					Operation{Data: And},
					Element{Data: "X2", IsNegative: false},
					Operation{Data: And},
					Element{Data: "X3", IsNegative: false},
				},
				IsNegative: false,
			},
		},
		IsNegative: false,
	},
	Value: "Y1",
}

// TestCNFStepAndNo is a Data for testing  rule of De Morgan
var TestCNFStepAndNo = Expression{
	Data: Group{
		Data: []Elements{
			Group{
				Data: []Elements{
					Element{Data: "X1", IsNegative: false},
					Operation{Data: Or},
					Element{Data: "X2", IsNegative: false},
					Operation{Data: Or},
					Element{Data: "X3", IsNegative: true},
				},
				IsNegative: true,
			},
			Operation{Data: Or},
			Group{
				Data: []Elements{
					Element{Data: "X1", IsNegative: false},
					Operation{Data: Or},
					Element{Data: "X2", IsNegative: true},
					Operation{Data: Or},
					Element{Data: "X3", IsNegative: false},
				},
				IsNegative: true,
			},
			Operation{Data: Or},
			Group{
				Data: []Elements{
					Element{Data: "X1", IsNegative: true},
					Operation{Data: Or},
					Element{Data: "X2", IsNegative: false},
					Operation{Data: Or},
					Element{Data: "X3", IsNegative: false},
				},
				IsNegative: true,
			},
			Operation{Data: Or},
			Group{
				Data: []Elements{
					Element{Data: "X1", IsNegative: true},
					Operation{Data: Or},
					Element{Data: "X2", IsNegative: true},
					Operation{Data: Or},
					Element{Data: "X3", IsNegative: false},
				},
				IsNegative: true,
			},
		},
		IsNegative: true,
	},
	Value: "Y1",
}

// TestDNFStepOrNo is a Data for testing  rule of De Morgan
var TestDNFStepOrNo = Expression{
	Data: Group{
		Data: []Elements{
			Group{
				Data: []Elements{
					Element{Data: "X1", IsNegative: true},
					Operation{Data: And},
					Element{Data: "X2", IsNegative: true},
					Operation{Data: And},
					Element{Data: "X3", IsNegative: true},
				},
				IsNegative: true,
			},
			Operation{Data: And},
			Group{
				Data: []Elements{
					Element{Data: "X1", IsNegative: true},
					Operation{Data: And},
					Element{Data: "X2", IsNegative: false},
					Operation{Data: And},
					Element{Data: "X3", IsNegative: false},
				},
				IsNegative: true,
			},
			Operation{Data: And},
			Group{
				Data: []Elements{
					Element{Data: "X1", IsNegative: false},
					Operation{Data: And},
					Element{Data: "X2", IsNegative: true},
					Operation{Data: And},
					Element{Data: "X3", IsNegative: false},
				},
				IsNegative: true,
			},
			Operation{Data: And},
			Group{
				Data: []Elements{
					Element{Data: "X1", IsNegative: false},
					Operation{Data: And},
					Element{Data: "X2", IsNegative: false},
					Operation{Data: And},
					Element{Data: "X3", IsNegative: false},
				},
				IsNegative: true,
			},
		},
		IsNegative: true,
	},
	Value: "Y1",
}

// TestDNFStepAndOrNo is a Data for testing  rule of De Morgan
var TestDNFStepAndOrNo = Expression{
	Data: Group{
		Data: []Elements{
			Group{
				Data: []Elements{
					Element{Data: "X1", IsNegative: false},
					Operation{Data: Or},
					Element{Data: "X2", IsNegative: false},
					Operation{Data: Or},
					Element{Data: "X3", IsNegative: false},
				},
				IsNegative: false,
			},
			Operation{Data: And},
			Group{
				Data: []Elements{
					Element{Data: "X1", IsNegative: false},
					Operation{Data: Or},
					Element{Data: "X2", IsNegative: true},
					Operation{Data: Or},
					Element{Data: "X3", IsNegative: true},
				},
				IsNegative: false,
			},
			Operation{Data: And},
			Group{
				Data: []Elements{
					Element{Data: "X1", IsNegative: true},
					Operation{Data: Or},
					Element{Data: "X2", IsNegative: false},
					Operation{Data: Or},
					Element{Data: "X3", IsNegative: true},
				},
				IsNegative: false,
			},
			Operation{Data: And},
			Group{
				Data: []Elements{
					Element{Data: "X1", IsNegative: true},
					Operation{Data: Or},
					Element{Data: "X2", IsNegative: true},
					Operation{Data: Or},
					Element{Data: "X3", IsNegative: true},
				},
				IsNegative: false,
			},
		},
		IsNegative: true,
	},
	Value: "Y1",
}

// TestCNFStepAndNoOr is a Data for testing  rule of De Morgan
var TestCNFStepAndNoOr = Expression{
	Data: Group{
		Data: []Elements{
			Group{
				Data: []Elements{
					Element{Data: "X1", IsNegative: true},
					Operation{Data: And},
					Element{Data: "X2", IsNegative: true},
					Operation{Data: And},
					Element{Data: "X3", IsNegative: false},
				},
				IsNegative: false,
			},
			Operation{Data: Or},
			Group{
				Data: []Elements{
					Element{Data: "X1", IsNegative: true},
					Operation{Data: And},
					Element{Data: "X2", IsNegative: false},
					Operation{Data: And},
					Element{Data: "X3", IsNegative: true},
				},
				IsNegative: false,
			},
			Operation{Data: Or},
			Group{
				Data: []Elements{
					Element{Data: "X1", IsNegative: false},
					Operation{Data: And},
					Element{Data: "X2", IsNegative: true},
					Operation{Data: And},
					Element{Data: "X3", IsNegative: true},
				},
				IsNegative: false,
			},
			Operation{Data: Or},
			Group{
				Data: []Elements{
					Element{Data: "X1", IsNegative: false},
					Operation{Data: And},
					Element{Data: "X2", IsNegative: false},
					Operation{Data: And},
					Element{Data: "X3", IsNegative: true},
				},
				IsNegative: false,
			},
		},
		IsNegative: true,
	},
	Value: "Y1",
}

// TestDNFStepOrNoOr is a Data for testing  rule of De Morgan
var TestDNFStepOrNoOr = Expression{
	Data: Group{
		Data: []Elements{
			Group{
				Data: []Elements{
					Element{Data: "X1", IsNegative: false},
					Operation{Data: Or},
					Element{Data: "X2", IsNegative: false},
					Operation{Data: Or},
					Element{Data: "X3", IsNegative: false},
				},
				IsNegative: true,
			},
			Operation{Data: Or},
			Group{
				Data: []Elements{
					Element{Data: "X1", IsNegative: false},
					Operation{Data: Or},
					Element{Data: "X2", IsNegative: true},
					Operation{Data: Or},
					Element{Data: "X3", IsNegative: true},
				},
				IsNegative: true,
			},
			Operation{Data: Or},
			Group{
				Data: []Elements{
					Element{Data: "X1", IsNegative: true},
					Operation{Data: Or},
					Element{Data: "X2", IsNegative: false},
					Operation{Data: Or},
					Element{Data: "X3", IsNegative: true},
				},
				IsNegative: true,
			},
			Operation{Data: Or},
			Group{
				Data: []Elements{
					Element{Data: "X1", IsNegative: true},
					Operation{Data: Or},
					Element{Data: "X2", IsNegative: true},
					Operation{Data: Or},
					Element{Data: "X3", IsNegative: true},
				},
				IsNegative: true,
			},
		},
		IsNegative: false,
	},
	Value: "Y1",
}

// TestCNFStepAndNoAnd is a Data for testing  rule of De Morgan
var TestCNFStepAndNoAnd = Expression{
	Data: Group{
		Data: []Elements{
			Group{
				Data: []Elements{
					Element{Data: "X1", IsNegative: true},
					Operation{Data: And},
					Element{Data: "X2", IsNegative: true},
					Operation{Data: And},
					Element{Data: "X3", IsNegative: false},
				},
				IsNegative: true,
			},
			Operation{Data: And},
			Group{
				Data: []Elements{
					Element{Data: "X1", IsNegative: true},
					Operation{Data: And},
					Element{Data: "X2", IsNegative: false},
					Operation{Data: And},
					Element{Data: "X3", IsNegative: true},
				},
				IsNegative: true,
			},
			Operation{Data: And},
			Group{
				Data: []Elements{
					Element{Data: "X1", IsNegative: false},
					Operation{Data: And},
					Element{Data: "X2", IsNegative: true},
					Operation{Data: And},
					Element{Data: "X3", IsNegative: true},
				},
				IsNegative: true,
			},
			Operation{Data: And},
			Group{
				Data: []Elements{
					Element{Data: "X1", IsNegative: false},
					Operation{Data: And},
					Element{Data: "X2", IsNegative: false},
					Operation{Data: And},
					Element{Data: "X3", IsNegative: true},
				},
				IsNegative: true,
			},
		},
		IsNegative: false,
	},
	Value: "Y1",
}

// TestGroup is a Data for testing grouping expression with elements 2&, 2 or
var TestGroup = Expression{
	Data: Group{
		Data: []Elements{
			Group{
				Data: []Elements{
					Group{
						Data: []Elements{
							Group{
								Data: []Elements{
									Group{
										Data: []Elements{
											Element{Data: "X1", IsNegative: true},
											Operation{Data: And},
											Element{Data: "X2", IsNegative: true},
										},
										IsNegative: false,
									},
									Operation{Data: And},
									Element{Data: "X3", IsNegative: true},
								},
								IsNegative: false,
							},
							Operation{Data: Or},
							Group{
								Data: []Elements{
									Group{
										Data: []Elements{
											Element{Data: "X1", IsNegative: true},
											Operation{Data: And},
											Element{Data: "X2", IsNegative: false},
										},
										IsNegative: false,
									},
									Operation{Data: And},
									Element{Data: "X3", IsNegative: false},
								},
								IsNegative: false,
							},
						},
						IsNegative: false,
					},
					Operation{Data: Or},
					Group{
						Data: []Elements{
							Group{
								Data: []Elements{
									Element{Data: "X1", IsNegative: false},
									Operation{Data: And},
									Element{Data: "X2", IsNegative: true},
								},
								IsNegative: false,
							},
							Operation{Data: And},
							Element{Data: "X3", IsNegative: false},
						},
						IsNegative: false,
					},
				},
				IsNegative: false,
			},
			Operation{Data: Or},
			Group{
				Data: []Elements{
					Group{
						Data: []Elements{
							Element{Data: "X1", IsNegative: false},
							Operation{Data: And},
							Element{Data: "X2", IsNegative: false},
						},
						IsNegative: false,
					},
					Operation{Data: And},
					Element{Data: "X3", IsNegative: false},
				},
				IsNegative: false,
			},
		},
		IsNegative: false,
	},
	Value: "Y1",
}
