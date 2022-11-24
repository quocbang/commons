package kendaql

//go:generate protoc -I=. -I=../../.. --go_out=paths=source_relative:../../../../proto/golang/dm/kendaql data_type.proto
//go:generate protoc -I=. -I=../../.. --go_out=paths=source_relative:../../../../proto/golang/dm/kendaql db.proto
//go:generate protoc -I=. -I=../../.. --go_out=paths=source_relative:../../../../proto/golang/dm/kendaql field.proto
//go:generate protoc -I=. -I=../../.. --go_out=paths=source_relative:../../../../proto/golang/dm/kendaql table.proto
//go:generate protoc -I=. -I=../../.. --go_out=paths=source_relative:../../../../proto/golang/dm/kendaql mutation.proto
//go:generate protoc -I=. -I=../../.. --go_out=paths=source_relative:../../../../proto/golang/dm/kendaql query.proto

//go:generate protoc -I=. -I=../../.. -I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.11.3/third_party/googleapis --go_out=paths=source_relative,plugins=grpc:../../../../proto/golang/dm/kendaql service.proto
//go:generate protoc -I=. -I=../../.. -I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.11.3/third_party/googleapis --grpc-gateway_out=paths=source_relative,logtostderr=true:../../../../proto/golang/dm/kendaql service.proto

//go:generate go vet ../../../../...
