package controller

import (
	"fmt"
	"net/http"
)

func Insert(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		fmt.Fprintf(writer, "insert "+request.Method+"\n")
	}
}

func Update(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "PUT" {
		fmt.Fprintf(writer, "update "+request.Method+"\n")
	}
}

func Delete(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "DELETE" {
		fmt.Fprintf(writer, "delete "+request.Method+"\n")
	}
}

func Get(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		fmt.Fprintf(writer, "get "+request.Method+"\n")
	}
}
