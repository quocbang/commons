package plm

//go:generate protoc --go_out=paths=source_relative:../../../../proto/golang/dm/plm plmmas.proto
//go:generate protoc --go_out=paths=source_relative:../../../../proto/golang/dm/plm plmatt.proto
//go:generate protoc --go_out=paths=source_relative:../../../../proto/golang/dm/plm list_plmmas_tree.proto
//go:generate protoc --go_out=paths=source_relative:../../../../proto/golang/dm/plm update_plmatt_status.proto
//go:generate protoc --go_out=paths=source_relative:../../../../proto/golang/dm/plm get_plmatt_max_seq.proto
//go:generate protoc -I=. -I=../../.. --go_out=paths=source_relative:../../../../proto/golang/dm/plm list_plmmas.proto
//go:generate protoc -I=. -I=../../.. --go_out=paths=source_relative:../../../../proto/golang/dm/plm list_plmatt.proto

//go:generate protoc -I=. -I=../../.. -I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.11.3/third_party/googleapis --go_out=paths=source_relative,plugins=grpc:../../../../proto/golang/dm/plm service.proto
//go:generate protoc -I=. -I=../../.. -I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.11.3/third_party/googleapis --grpc-gateway_out=paths=source_relative,logtostderr=true:../../../../proto/golang/dm/plm service.proto

//go:generate go vet ../../../../...
