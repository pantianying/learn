package main

import (
	"encoding/json"
	"fmt"
)

var v string = `{"aaa":"value1"}`

type TestType1 struct {
	AaA string
}

func main() {
	test2()
}
func test1() {
	var test TestType1
	json.Unmarshal([]byte(v), &test)
	fmt.Printf("resp:%+v", test)
}

var v2 string = `{"a":1,"c":{"cc":"xxx"}}`

func test2() {
	var test interface{}
	json.Unmarshal([]byte(v2), &test)
	fmt.Printf("resp:%+v", test)
}

var v3 string = `{
  "pid1,pid2":{
    " standardCode_1":"dpCode_1",
    " standardCode_2":"dpCode_2"
  },
  "pid3":{
    " standardCode_1":"dpCode1_1",
    " standardCode_2":"dpCode_2"
  }
}`

func test3() {
	var m map[string]interface{}
	json.Unmarshal([]byte(v3), &m)
}
