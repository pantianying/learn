module github.com/pantianying/learn/go

go 1.12

require github.com/gorilla/mux v1.7.3

require (
	github.com/apache/dubbo-go-hessian2 v1.2.5-0.20190731020727-1697039810c8
	github.com/apache/pulsar/pulsar-function-go v0.0.0-20191217081830-dfc8c1cad331
	github.com/coreos/etcd v3.3.10+incompatible
	github.com/didip/tollbooth v4.0.2+incompatible
	github.com/golang/protobuf v1.3.2
	github.com/yuin/gopher-lua v0.0.0-20190514113301-1cd887cd7036
	google.golang.org/grpc v1.21.0
)

replace github.com/apache/dubbo-go-hessian2 => github.com/pantianying/dubbo-go-hessian2 v1.0.2-0.20190826014112-14dd19e36832

//github.com/samuel/go-zookeeper v0.0.0-20180130194729-c4fab1ac1bec

replace github.com/samuel/go-zookeeper => github.com/samuel/go-zookeeper v0.0.0-20180130194729-c4fab1ac1bec
