package main

import (
	exp "MicroservicesPet/ServiceCalculator/calculator/expressions"
	"MicroservicesPet/ServiceCalculator/calculator/schema"
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
	_, err = w.Write(js)
	if err != nil {
		log.Fatalf("Error " + r.Method)
	}
}

func handleSchema(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	log.Println("Request from user")
	js, err := json.Marshal(schema.TestSchema) // exp.MakeSchema(testDataSchema)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(js)
	if err != nil {
		log.Fatalf("Error " + r.Method)
	}
}

func handleInput(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	log.Println("Request from user with input")
	data := exp.LogicalTableData{}
	err := json.NewDecoder(r.Body).Decode(&data)
	js, err := json.Marshal(exp.MakeExpression(data)) // exp.MakeSchema(testDataSchema)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(js)
	if err != nil {
		log.Fatalf("Error")
	}
}

var data = []int{1, 0, 0, 1, 0, 1, 0, 1}
var elements = exp.LimitElements{
	Or:    0,
	OrNo:  0,
	And:   0,
	AndNo: 0,
}
var lt = exp.OrAndNo
var xData = []string{"X1", "X2", "X3"}
var yData = "Y1"
var testData = exp.LogicalTableData{Y: yData, X: xData, Data: data, LogicType: lt, Limits: elements, Type: exp.DNF}

func main() {
	route := mux.NewRouter()
	route.HandleFunc("/data", handleArticles)
	route.HandleFunc("/schema", handleSchema)
	route.HandleFunc("/input", handleInput)
	route.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data, _ := json.Marshal(exp.MakeExpression(testData))
		_, err := w.Write(data)
		if err != nil {
			log.Fatalf("Error")
		}
	}).Methods(http.MethodGet)

	s := http.Server{
		Addr:    ":8080",
		Handler: route,
	}
	err := s.ListenAndServe()
	if err != nil {
		return
	}
}
