package main

import (
	"fmt"

	"learning-golang/helloworld"
)

// custom module is created using ```go mod init ahmadbshaik/golang-practice``` to modularize the code

func main() {
	fmt.Println(helloworld.HelloWorld())
}
