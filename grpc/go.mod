module github.com/myuser/myrepo

go 1.17

replace proto => ./proto

require (
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.8.0
	github.com/sirupsen/logrus v1.8.1
	golang.org/x/net v0.0.0-20220127200216-cd36cc0744dd
	google.golang.org/genproto v0.0.0-20220302033224-9aa15565e42a
	google.golang.org/grpc v1.44.0
	google.golang.org/protobuf v1.27.1
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/sys v0.0.0-20211216021012-1d35b9e2eb4e // indirect
	golang.org/x/text v0.3.7 // indirect
)
