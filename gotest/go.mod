module github.com/pantianying/keeplearning/gotest

go 1.12

require github.com/gorilla/mux v1.7.3

require (
	github.com/apache/dubbo-go-hessian2 v1.2.5-0.20190731020727-1697039810c8
	github.com/didip/tollbooth v4.0.2+incompatible
	github.com/hashicorp/golang-lru v0.5.3 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/throttled/throttled v2.2.4+incompatible // indirect
	github.com/yuin/gopher-lua v0.0.0-20190514113301-1cd887cd7036
	golang.org/x/time v0.0.0-20190921001708-c4c64cad1fd0 // indirect
)

replace github.com/apache/dubbo-go-hessian2 => github.com/pantianying/dubbo-go-hessian2 v1.0.2-0.20190826014112-14dd19e36832
