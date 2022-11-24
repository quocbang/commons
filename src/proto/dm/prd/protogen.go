package prd

//go:generate protoc -I=. -I=../../.. --go_out=paths=source_relative:../../../../proto/golang/dm/prd label.proto
//go:generate protoc --go_out=paths=source_relative:../../../../proto/golang/dm/prd query_prd_man.proto
//go:generate protoc --go_out=paths=source_relative:../../../../proto/golang/dm/prd mtrl_ctrl_record.proto
//go:generate protoc --go_out=paths=source_relative:../../../../proto/golang/dm/prd mtrl_extended_record.proto
//go:generate protoc --go_out=paths=source_relative:../../../../proto/golang/dm/prd query_batch_quantity.proto
//go:generate protoc --go_out=paths=source_relative:../../../../proto/golang/dm/prd add_pcr_first_move_in_material_record.proto
//go:generate protoc --go_out=paths=source_relative:../../../../proto/golang/dm/prd add_pcr_second_move_in_material_record.proto
//go:generate protoc  -I=. -I=../../.. --go_out=paths=source_relative:../../../../proto/golang/dm/prd label.proto

//go:generate protoc -I=. -I=../../.. -I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.11.3/third_party/googleapis --go_out=paths=source_relative,plugins=grpc:../../../../proto/golang/dm/prd service.proto
//go:generate protoc -I=. -I=../../.. -I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.11.3/third_party/googleapis --grpc-gateway_out=paths=source_relative,logtostderr=true:../../../../proto/golang/dm/prd service.proto

//go:generate go vet ../../../../...
