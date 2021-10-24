package main

import (
	"bouncedb/format"
	"bouncedb/http"
	"bouncedb/user"
	"fmt"
)

type Test struct {
	name    string
	test    int
	boolean bool
}

func main() {
	// json := format.ToJsonObject("test: 1337, naruto: \"Super Intense Minecraft Player 😎 I'm a SIMP.\" ")
	v := Test{test: 1, name: "Peter", boolean: false}

	fmt.Println(format.FromStructToJson(v))

	go http.Http()
	user.User()

	select {}
}
