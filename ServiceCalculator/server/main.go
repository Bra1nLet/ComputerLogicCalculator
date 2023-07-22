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
	js, err := json.Marshal(exp.MakeExpression(testData))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

var data = []int{1, 0, 0, 1, 0, 1, 0, 1}
var elements = [2]int{3, 2}
var lt = exp.OrAndNo
var xData = []string{"X1", "X2", "X3"}
var yData = "Y1"
var testData = exp.LogicalTableData{Y: yData, X: xData, Data: data, LogicType: lt, Elements: elements, Type: exp.CNF}

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
