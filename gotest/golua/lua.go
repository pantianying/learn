package main

import (
	lua "github.com/yuin/gopher-lua"
)

func main() {
	//test1()
	test2()
}
func test1() {
	L := lua.NewState()
	defer L.Close()
	if err := L.DoString(`print("hello")`); err != nil {
		panic(err)
	}
}
func test2() {
	L := lua.NewState()
	defer L.Close()
	if err := L.DoFile("hello.lua"); err != nil {
		panic(err)
	}
}
