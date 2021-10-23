package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Data struct {
	Name    string
	Version int
	Time    string
}

func toJson(object interface{}) string {
	jsonResponse, _ := json.MarshalIndent(object, "", "\t")
	return string(jsonResponse)
}

func Insert(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		database := request.URL.Path[len("/insert/"):]
		//fmt.Fprintf(writer, "insert database: "+database+"\n")
		fmt.Println(database)
		data := Data{"Peu77", 77, "heute lol"}
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusCreated)
		JData, err := json.Marshal(data)
		if err != nil {

		}
		writer.Write(JData)
	}
}

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	data := Data{"Peu77", 77, "heute lol"}

	fmt.Fprintf(w, toJson(data))
	return
}

func Update(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "PUT" {
		database := request.URL.Path[len("/update/"):]
		fmt.Fprintf(writer, "update "+request.Method+" "+database+"\n")
	}
}

func Delete(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "DELETE" {
		database := request.URL.Path[len("/delete/"):]
		fmt.Fprintf(writer, "delete "+request.Method+" "+database+" \n")
	}
}

func Get(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		database := request.URL.Path[len("/get/"):]
		fmt.Fprintf(writer, "get "+request.Method+" "+database+"\n")
	}
}
