package main

import (
	exp "MicroservicesPet/ServiceCalculator/calculator/expressions"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func handleArticles(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	log.Println("Request from user")
	js, err := json.Marshal(exp.MakeExpression(testData)) // exp.MakeExpression(testData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

var data = []int{1, 0, 0, 1, 0, 1, 0, 1}
var elements = []int{9}
var lt = exp.OrAndNo
var xData = []string{"X1", "X2", "X3"}
var yData = "Y1"
var testData = exp.LogicalTableData{Y: yData, X: xData, Data: data, LogicType: lt, Elements: elements, Type: exp.DNF}

func main() {
	route := mux.NewRouter()
	route.HandleFunc("/data", handleArticles)
	route.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data, _ := json.Marshal(exp.MakeExpression(testData))
		w.Write(data)
	}).Methods(http.MethodGet)

	s := http.Server{
		Addr:    ":8080",
		Handler: route,
	}
	s.ListenAndServe()
}

var TestGroup = exp.Expression{
	Data: exp.Group{
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
		},
		IsNegative: false,
	},
	Value: "Y1",
}
