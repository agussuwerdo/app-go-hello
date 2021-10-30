package main

import (
	"fmt"

	go_hello "github.com/agussuwerdo/go-hello/v2"
)

func main() {
	// untuk get module
	// go get github.com/agussuwerdo/go-hello
	fmt.Println(go_hello.SayHello("agus"))
}
