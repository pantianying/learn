package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "AAA"
	b := "BBB"
	c := "CCC"
	fmt.Println(strings.Join([]string{a, b, c}, "."))
}
