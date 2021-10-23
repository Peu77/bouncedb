package http

import (
	"bouncedb/http/controller"
	"fmt"
	"net/http"
)

func Http() {
	http.HandleFunc("/insert/", controller.Insert)
	http.HandleFunc("/update", controller.Update)
	http.HandleFunc("/delete", controller.Delete)
	http.HandleFunc("/get", controller.Get)
	http.HandleFunc("/test", controller.HandleRequest)

	err := http.ListenAndServe(":4001", nil)
	if err != nil {
		return
	}
	fmt.Println("init http server")
}
