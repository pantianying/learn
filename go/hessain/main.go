package main

import (
	hessian "github.com/apache/dubbo-go-hessian2"
	"github.com/pantianying/learn/go/debug"
)

func main() {
	target := "🤣xxx"
	e := hessian.NewEncoder()
	err := e.Encode(target)
	if err != nil {
		debug.Debug(err)
	}
	encodedBuf := e.Buffer()
	debug.Debug("after encode:", encodedBuf, string(encodedBuf))
	devodedBuf, _ := decode(encodedBuf)
	debug.Debug("after decode:", devodedBuf)

}
func decode(b []byte) (interface{}, error) {
	d := hessian.NewDecoder(b)
	r, e := d.Decode()
	if e != nil {
		return nil, e
	}
	return r, nil
}
