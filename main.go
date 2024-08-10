package main

import (
	"fmt"

	"golang-practice/helloworld"
)

// custom module is created using ```go mod init ahmadbshaik/golang-practice``` to modularize the code

func main() {
	fmt.Println(helloworld.HelloWorld())
}
