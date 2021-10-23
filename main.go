package main

import (
	"bouncedb/format"
	"bouncedb/http"
	"bouncedb/user"
	"fmt"
)

func main() {
	fmt.Println(format.ToJsonObject("test: 1337, naruto: \"Super Intense Minecraft Player 😎 I'm a SIMP.\" "))

	go http.Http()
	user.User()

	select {}
}
