package main

import (
	"bouncedb/format"
	"bouncedb/http"
	"bouncedb/user"
	"fmt"
)

func main() {
	json := format.ToJsonObject("test: 1337, naruto: \"Super Intense Minecraft Player 😎 I'm a SIMP.\" ")

	fmt.Println(format.FromJsonObject(json))

	go http.Http()
	user.User()

	select {}
}
