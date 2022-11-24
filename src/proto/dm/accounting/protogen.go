package accounting

//go:generate protoc --go_out=paths=source_relative:../../../../proto/golang/dm/accounting credit_or_debit.proto
//go:generate protoc --go_out=paths=source_relative:../../../../proto/golang/dm/accounting entry_type.proto
//go:generate protoc -I=. -I=../../.. --go_out=paths=source_relative:../../../../proto/golang/dm/accounting accounting_entry.proto

//go:generate protoc -I=. -I=../../.. -I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.11.3/third_party/googleapis --go_out=paths=source_relative,plugins=grpc:../../../../proto/golang/dm/accounting service.proto
//go:generate protoc -I=. -I=../../.. -I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.11.3/third_party/googleapis --grpc-gateway_out=paths=source_relative,logtostderr=true:../../../../proto/golang/dm/accounting service.proto

//go:generate go vet ../../../../...
