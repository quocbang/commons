package mes

//go:generate protoc --go_out=paths=source_relative:../../../../proto/golang/dm/mes work_order_struct.proto
//go:generate protoc --go_out=paths=source_relative:../../../../proto/golang/dm/mes create_work_order_head.proto
//go:generate protoc --go_out=paths=source_relative:../../../../proto/golang/dm/mes create_work_order.proto
//go:generate protoc --go_out=paths=source_relative:../../../../proto/golang/dm/mes modify_work_order_status.proto
//go:generate protoc --go_out=paths=source_relative:../../../../proto/golang/dm/mes add_m_sheet_bom_record.proto
//go:generate protoc --go_out=paths=source_relative:../../../../proto/golang/dm/mes add_material.proto
//go:generate protoc --go_out=paths=source_relative:../../../../proto/golang/dm/mes create_prepare_order.proto
//go:generate protoc --go_out=paths=source_relative:../../../../proto/golang/dm/mes query_prepare_order_list.proto
//go:generate protoc --go_out=paths=source_relative:../../../../proto/golang/dm/mes modify_mtrllst_record.proto
//go:generate protoc -I=. -I=../../.. --go_out=paths=source_relative:../../../../proto/golang/dm/mes query_material_list.proto
//go:generate protoc --go_out=paths=source_relative:../../../../proto/golang/dm/mes delete_materials.proto
//go:generate protoc --go_out=paths=source_relative:../../../../proto/golang/dm/mes modify_prepare_order_info.proto
//go:generate protoc --go_out=paths=source_relative:../../../../proto/golang/dm/mes query_code.proto
//go:generate protoc --go_out=paths=source_relative:../../../../proto/golang/dm/mes query_sheet_info.proto
//go:generate protoc --go_out=paths=source_relative:../../../../proto/golang/dm/mes create_sheet_info.proto
//go:generate protoc --go_out=paths=source_relative:../../../../proto/golang/dm/mes modify_sheet_info.proto
//go:generate protoc --go_out=paths=source_relative:../../../../proto/golang/dm/mes query_replace_mtrl.proto
//go:generate protoc --go_out=paths=source_relative:../../../../proto/golang/dm/mes create_container_info.proto
//go:generate protoc --go_out=paths=source_relative:../../../../proto/golang/dm/mes query_container_info.proto
//go:generate protoc --go_out=paths=source_relative:../../../../proto/golang/dm/mes delete_container.proto
//go:generate protoc --go_out=paths=source_relative:../../../../proto/golang/dm/mes modify_container_info.proto
//go:generate protoc --go_out=paths=source_relative:../../../../proto/golang/dm/mes modify_storage_id.proto
//go:generate protoc --go_out=paths=source_relative:../../../../proto/golang/dm/mes delete_mconmtrlst.proto
//go:generate protoc --go_out=paths=source_relative:../../../../proto/golang/dm/mes create_mconmtrlst.proto
//go:generate protoc --go_out=paths=source_relative:../../../../proto/golang/dm/mes modify_mconmtrlst.proto
//go:generate protoc --go_out=paths=source_relative:../../../../proto/golang/dm/mes query_mconmtrlst.proto
//go:generate protoc --go_out=paths=source_relative:../../../../proto/golang/dm/mes delete_meqpt.proto
//go:generate protoc --go_out=paths=source_relative:../../../../proto/golang/dm/mes create_meqpt.proto
//go:generate protoc --go_out=paths=source_relative:../../../../proto/golang/dm/mes modify_meqpt.proto
//go:generate protoc --go_out=paths=source_relative:../../../../proto/golang/dm/mes query_meqpt.proto
//go:generate protoc --go_out=paths=source_relative:../../../../proto/golang/dm/mes delete_meqptcon.proto
//go:generate protoc --go_out=paths=source_relative:../../../../proto/golang/dm/mes create_meqptcon.proto
//go:generate protoc --go_out=paths=source_relative:../../../../proto/golang/dm/mes modify_meqptcon.proto
//go:generate protoc --go_out=paths=source_relative:../../../../proto/golang/dm/mes query_meqptcon.proto
//go:generate protoc --go_out=paths=source_relative:../../../../proto/golang/dm/mes query_eqpt_list.proto
//go:generate protoc --go_out=paths=source_relative:../../../../proto/golang/dm/mes query_eqpt_material_list.proto
//go:generate protoc --go_out=paths=source_relative:../../../../proto/golang/dm/mes query_eqpt_detail.proto
//go:generate protoc --go_out=paths=source_relative:../../../../proto/golang/dm/mes carrier_operating.proto
//go:generate protoc --go_out=paths=source_relative:../../../../proto/golang/dm/mes carrier_material_operating.proto
//go:generate protoc -I=. -I=../../.. --go_out=paths=source_relative:../../../../proto/golang/dm/mes mtrl_tx_list.proto
//go:generate protoc --go_out=paths=source_relative:../../../../proto/golang/dm/mes materials_sheet_operating.proto
//go:generate protoc --go_out=paths=source_relative:../../../../proto/golang/dm/mes mtx_list.proto
//go:generate protoc --go_out=paths=source_relative:../../../../proto/golang/dm/mes mtx_ship_list.proto
//go:generate protoc --go_out=paths=source_relative:../../../../proto/golang/dm/mes add_meas_value.proto
//go:generate protoc --go_out=paths=source_relative:../../../../proto/golang/dm/mes get_mtrl_list_by_dept.proto
//go:generate protoc --go_out=paths=source_relative:../../../../proto/golang/dm/mes add_log_record.proto
//go:generate protoc --go_out=paths=source_relative:../../../../proto/golang/dm/mes get_container_by_container_type.proto

//go:generate protoc -I=. -I=../../.. -I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.11.3/third_party/googleapis --go_out=paths=source_relative,plugins=grpc:../../../../proto/golang/dm/mes service.proto

//go:generate go vet ../../../../...
