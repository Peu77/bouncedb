package controller

import (
	"fmt"
	"net/http"
)

func Insert(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		database := request.URL.Path[len("/insert/"):]
		fmt.Fprintf(writer, "insert database: "+database+"\n")
	}
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
