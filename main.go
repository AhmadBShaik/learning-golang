package main

import (
	"fmt"

	"learning-golang/controlflow"
	"learning-golang/helloworld"
)

// custom module is created using ```go mod init ahmadbshaik/golang-practice``` to modularize the code

func main() {
	fmt.Println(helloworld.HelloWorld())

	controlflow.ConditionalStatements1()
	controlflow.ConditionalStatements2()

	controlflow.Loop1()
	controlflow.Loop2()
	controlflow.Loop3()
}
