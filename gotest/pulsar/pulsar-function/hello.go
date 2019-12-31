package main

import (
	"fmt"

	"github.com/apache/pulsar/pulsar-function-go/pf"
)

func hello() {
	fmt.Println("hello pulsar function")
}

func main() {
	pf.Start(hello)
}
