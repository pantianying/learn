package main

import (
	"fmt"
	"github.com/pantianying/keeplearning/gotest/debug"
	"reflect"
)

type A struct {
	Num int
	Com string
}

//https://studygolang.com/articles/12348?fr=sidebar
func main() {
	//fmtRefeclt()
	//test2()
	//user := User{1, "Allen.Wu", 25}
	//
	//test3(user)
	result := GetFunctionName(MyFunction)
	fmt.Println(result)
}

func test1() {
	var a = A{
		Num: 1,
		Com: "xxx",
	}
	v := reflect.ValueOf(a)
	fmt.Println("kind()", v.Kind())
	fmt.Println("Type()", v.Type())
	fmt.Println(v.Field(1))
}
func test2() {
	var num float64 = 1.2345

	pointer := reflect.ValueOf(&num)
	value := reflect.ValueOf(num)

	// 可以理解为“强制转换”，但是需要注意的时候，转换的时候，如果转换的类型不完全符合，则直接panic
	// Golang 对类型要求非常严格，类型一定要完全符合
	// 如下两个，一个是*float64，一个是float64，如果弄混，则会panic
	convertPointer := pointer.Interface().(*float64)
	convertValue := value.Interface().(float64)

	fmt.Println(convertPointer)
	fmt.Println(convertValue)
}

func test3(input interface{}) {
	getType := reflect.TypeOf(input)
	fmt.Println("get Type is :", getType.Name())

	getValue := reflect.ValueOf(input)
	fmt.Println("get all Fields is:", getValue)

	// 获取方法字段
	// 1. 先获取interface的reflect.Type，然后通过NumField进行遍历
	// 2. 再通过reflect.Type的Field获取其Field
	// 3. 最后通过Field的Interface()得到对应的value
	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		value := getValue.Field(i).Interface()
		debug.Debugf("%s: %v = %v\n", field.Name, field.Type, value)
	}

	// 获取方法
	// 1. 先获取interface的reflect.Type，然后通过.NumMethod进行遍历
	for i := 0; i < getType.NumMethod(); i++ {
		m := getType.Method(i)
		debug.Debugf("%s: %v\n", m.Name, m.Type)
	}
}

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) ReflectCallFunc() {
	debug.Debug("Allen.Wu ReflectCallFunc")
}
