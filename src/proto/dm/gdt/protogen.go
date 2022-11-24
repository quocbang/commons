package gdt

//go:generate protoc --go_out=paths=source_relative:../../../../proto/golang/dm/gdt query_gdtbar.proto
//go:generate protoc --go_out=paths=source_relative:../../../../proto/golang/dm/gdt query_gdtpat.proto

//go:generate protoc -I=. -I=../../.. -I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.11.3/third_party/googleapis --go_out=paths=source_relative,plugins=grpc:../../../../proto/golang/dm/gdt service.proto
//go:generate protoc -I=. -I=../../.. -I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.11.3/third_party/googleapis --grpc-gateway_out=paths=source_relative,logtostderr=true:../../../../proto/golang/dm/gdt service.proto

//go:generate go vet ../../../../...
