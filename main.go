package main

import (
	"bouncedb/format"
	"bouncedb/http"
	"bouncedb/user"
	"fmt"
)

func main() {
	fmt.Println(format.ToJsonObject("test: 12 \n naruto: \"Du denkst wirklich, dass du lustig bist?\""))

	http.Http()
	user.User()
}
