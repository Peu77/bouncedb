package http

import (
	"fmt"
	"net/http"
)

func Http() {
	http.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "hello\n")
	})

	err := http.ListenAndServe(":4001", nil)
	if err != nil {
		return
	}
}
