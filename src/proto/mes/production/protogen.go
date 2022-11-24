package mes

// NOTE: $GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.11.3 should
//       match the version in go.mod and Dockerfile. Once everyone is on non-Windows
//       systems we can use go list to get the path

// production
//go:generate protoc -I=. -I=../../.. --go_out=paths=source_relative:../../../../proto/golang/mes check_material_lot_id.proto
//go:generate protoc -I=. -I=../../.. --go_out=paths=source_relative:../../../../proto/golang/mes check_product_tools.proto
//go:generate protoc -I=. -I=../../.. --go_out=paths=source_relative:../../../../proto/golang/mes check_scan_material.proto
//go:generate protoc -I=. -I=../../.. --go_out=paths=source_relative:../../../../proto/golang/mes common.proto
//go:generate protoc -I=. -I=../../.. --go_out=paths=source_relative:../../../../proto/golang/mes equipment_material_operating.proto
//go:generate protoc -I=. -I=../../.. --go_out=paths=source_relative:../../../../proto/golang/mes material_operating.proto
//go:generate protoc -I=. -I=../../.. --go_out=paths=source_relative:../../../../proto/golang/mes mp_service.proto
//go:generate protoc -I=. -I=../../.. --go_out=paths=source_relative:../../../../proto/golang/mes work.proto
//go:generate protoc -I=. -I=../../.. -I=$GOPATH/pkg/mod/github.com/gogo/protobuf@v1.3.1 --go_out=paths=source_relative:../../../../proto/golang/mes work_order.proto
//go:generate protoc -I=. -I=../../.. --go_out=paths=source_relative:../../../../proto/golang/mes worker.proto
//go:generate protoc -I=. -I=../../.. --go_out=paths=source_relative:../../../../proto/golang/mes mp_service.proto
//go:generate protoc -I=. -I=../../.. --go_out=paths=source_relative:../../../../proto/golang/mes material_inventory.proto
