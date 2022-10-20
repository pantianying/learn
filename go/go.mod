module github.com/pantianying/learn/go

go 1.12

require github.com/gorilla/mux v1.7.3

require (
	github.com/TuyaInc/pulsar-client-go v0.0.0-20201113072945-941bdb3e51a3
	github.com/apache/dubbo-go-hessian2 v1.2.5-0.20190731020727-1697039810c8
	github.com/apache/pulsar/pulsar-function-go v0.0.0-20191217081830-dfc8c1cad331
	github.com/coreos/bbolt v0.0.0-00010101000000-000000000000 // indirect
	github.com/coreos/etcd v3.3.10+incompatible
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/pkg v0.0.0-20220810130054-c7d1c02cb6cf // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/didip/tollbooth v4.0.2+incompatible
	github.com/golang/protobuf v1.5.2
	github.com/gomodule/redigo v1.8.9 // indirect
	github.com/google/btree v1.1.2 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0 // indirect
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.16.0 // indirect
	github.com/javtube/javtube-sdk-go v1.0.22-0.20221018092240-b5fdca39718a
	github.com/jonboulle/clockwork v0.3.0 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/prometheus/client_golang v1.13.0 // indirect
	github.com/samuel/go-zookeeper v0.0.0-00010101000000-000000000000
	github.com/soheilhy/cmux v0.1.5 // indirect
	github.com/throttled/throttled v2.2.5+incompatible
	github.com/tmc/grpc-websocket-proxy v0.0.0-20220101234140-673ab2c3ae75 // indirect
	github.com/xiang90/probing v0.0.0-20190116061207-43a291ad63a2 // indirect
	github.com/yuin/gopher-lua v0.0.0-20190514113301-1cd887cd7036
	go.etcd.io/bbolt v1.3.5 // indirect
	go.uber.org/zap v1.23.0 // indirect
	golang.org/x/time v0.1.0 // indirect
	google.golang.org/grpc v1.33.1
)

replace (
	github.com/apache/dubbo-go-hessian2 => github.com/pantianying/dubbo-go-hessian2 v1.0.2-0.20190826014112-14dd19e36832

	github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.5

	//github.com/samuel/go-zookeeper v0.0.0-20180130194729-c4fab1ac1bec

	github.com/samuel/go-zookeeper => github.com/samuel/go-zookeeper v0.0.0-20180130194729-c4fab1ac1bec
	go.etcd.io/bbolt v1.3.5 => github.com/coreos/bbolt v1.3.5
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
)
