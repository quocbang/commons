package mes

// NOTE: $GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.11.3 should
//       match the version in go.mod and Dockerfile. Once everyone is on non-Windows
//       systems we can use go list to get the path

//go:generate protoc -I=. -I=../.. --go_out=paths=source_relative:../../../proto/golang/mes code/enums.proto
//go:generate protoc -I=. -I=../.. --go_out=paths=source_relative:../../../proto/golang/mes code/error/enums.proto

//go:generate protoc -I=. -I=../.. -I=$GOPATH/pkg/mod/github.com/gogo/protobuf@v1.3.1 -I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.11.3/third_party/googleapis --go_out=paths=source_relative,plugins=grpc:../../../proto/golang/mes services.proto
//go:generate protoc -I=. -I=../.. -I=$GOPATH/pkg/mod/github.com/gogo/protobuf@v1.3.1 -I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.11.3/third_party/googleapis --grpc-gateway_out=paths=source_relative,logtostderr=true:../../../proto/golang/mes services.proto
//go:generate protoc -I=. -I=../.. -I=$GOPATH/pkg/mod/github.com/gogo/protobuf@v1.3.1 -I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.11.3/third_party/googleapis --swagger_out=logtostderr=true:../../../proto/golang/mes/doc services.proto

//go:generate go vet ../../../...
