package main

import (
	"encoding/hex"
	"fmt"
	"strings"
)

func main() {
	test2()
}
func test1() {
	a := "AAA"
	b := "BBB"
	c := "CCC"
	fmt.Println(strings.Join([]string{a, b, c}, "."))
}

func test2() {
	a := []byte{0x77, 0x77}
	s := hex.EncodeToString(a)
	fmt.Println(s)
}
