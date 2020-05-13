package main

import (
	"github.com/pantianying/keeplearning/gotest/debug"
	"reflect"
)

func fmtRefeclt() {
	var (
		a = 1
		b = "w2"
	)
	debug.Debug("int------ kind :", reflect.TypeOf(a).Kind(), "type:", reflect.TypeOf(a).Name(), reflect.TypeOf(a).String())
	debug.Debug("string------ kind :", reflect.TypeOf(b).Kind(), "type:", reflect.TypeOf(b).Name(), reflect.TypeOf(b).String())
}
