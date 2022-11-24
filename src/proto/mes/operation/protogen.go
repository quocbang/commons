package mes

// NOTE: $GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.11.3 should
//       match the version in go.mod and Dockerfile. Once everyone is on non-Windows
//       systems we can use go list to get the path

// operation
//go:generate protoc -I=. -I=../../.. --go_out=paths=source_relative:../../../../proto/golang/mes list_factory_ids.proto
//go:generate protoc -I=. -I=../../.. --go_out=paths=source_relative:../../../../proto/golang/mes list_product_types.proto
//go:generate protoc -I=. -I=../../.. --go_out=paths=source_relative:../../../../proto/golang/mes list_equipment_ids.proto
//go:generate protoc -I=. -I=../../.. --go_out=paths=source_relative:../../../../proto/golang/mes list_product_id.proto
//go:generate protoc -I=. -I=../../.. --go_out=paths=source_relative:../../../../proto/golang/mes list_product_ids.proto
//go:generate protoc -I=. -I=../../.. --go_out=paths=source_relative:../../../../proto/golang/mes list_departments.proto
//go:generate protoc -I=. -I=../../.. --go_out=paths=source_relative:../../../../proto/golang/mes list_codes.proto
//go:generate protoc -I=. -I=../../.. --go_out=paths=source_relative:../../../../proto/golang/mes container_operating.proto
//go:generate protoc -I=. -I=../../.. --go_out=paths=source_relative:../../../../proto/golang/mes container_material_operating.proto
//go:generate protoc -I=. -I=../../.. --go_out=paths=source_relative:../../../../proto/golang/mes equipment_container_operating.proto
//go:generate protoc -I=. -I=../../.. --go_out=paths=source_relative:../../../../proto/golang/mes equipment_operating.proto
//go:generate protoc -I=. -I=../../.. --go_out=paths=source_relative:../../../../proto/golang/mes set_initial_process.proto
//go:generate protoc -I=. -I=../../.. --go_out=paths=source_relative:../../../../proto/golang/mes rapid_check_operating.proto
//go:generate protoc -I=. -I=../../.. --go_out=paths=source_relative:../../../../proto/golang/mes carrier_operating.proto
//go:generate protoc -I=. -I=../../.. --go_out=paths=source_relative:../../../../proto/golang/mes transaction_list_operating.proto
//go:generate protoc -I=. -I=../../.. --go_out=paths=source_relative:../../../../proto/golang/mes list_measurement_params.proto
//go:generate protoc -I=. -I=../../.. --go_out=paths=source_relative:../../../../proto/golang/mes transaction_operating.proto
//go:generate protoc -I=. -I=../../.. --go_out=paths=source_relative:../../../../proto/golang/mes transaction_shipping_operating.proto
//go:generate protoc -I=. -I=../../.. --go_out=paths=source_relative:../../../../proto/golang/mes get_bom_materials.proto
//go:generate protoc -I=. -I=../../.. --go_out=paths=source_relative:../../../../proto/golang/mes list_process_equipments.proto
//go:generate protoc -I=. -I=../../.. --go_out=paths=source_relative:../../../../proto/golang/mes list_recipe_all_process.proto
//go:generate protoc -I=. -I=../../.. --go_out=paths=source_relative:../../../../proto/golang/mes list_materials_categories_by_owned_department.proto
//go:generate protoc -I=. -I=../../.. --go_out=paths=source_relative:../../../../proto/golang/mes transaction_operating_pda.proto
