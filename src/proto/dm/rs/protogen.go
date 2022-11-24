package rs

//go:generate protoc -I=. -I=../../.. -I=$GOPATH/pkg/mod/github.com/gogo/protobuf@v1.3.1 --gogo_out=paths=source_relative:../../../../proto/golang/dm/rs recipe.proto
//go:generate protoc -I=. -I=../../.. -I=$GOPATH/pkg/mod/github.com/gogo/protobuf@v1.3.1 --gogo_out=paths=source_relative:../../../../proto/golang/dm/rs tire_profile.proto
//go:generate protoc -I=. -I=../../.. -I=$GOPATH/pkg/mod/github.com/gogo/protobuf@v1.3.1 --go_out=paths=source_relative:../../../../proto/golang/dm/rs material.proto
//go:generate protoc -I=. -I=../../.. -I=$GOPATH/pkg/mod/github.com/gogo/protobuf@v1.3.1 --go_out=paths=source_relative:../../../../proto/golang/dm/rs process.proto
//go:generate protoc -I=. -I=../../.. -I=$GOPATH/pkg/mod/github.com/gogo/protobuf@v1.3.1 --go_out=paths=source_relative:../../../../proto/golang/dm/rs property.proto
//go:generate protoc -I=. -I=../../.. -I=$GOPATH/pkg/mod/github.com/gogo/protobuf@v1.3.1 --go_out=paths=source_relative:../../../../proto/golang/dm/rs recipe_info.proto
//go:generate protoc -I=. -I=../../.. -I=$GOPATH/pkg/mod/github.com/gogo/protobuf@v1.3.1 --go_out=paths=source_relative:../../../../proto/golang/dm/rs recipe_dump.proto
//go:generate protoc -I=. -I=../../.. -I=$GOPATH/pkg/mod/github.com/gogo/protobuf@v1.3.1 --go_out=paths=source_relative:../../../../proto/golang/dm/rs spec_info.proto
//go:generate protoc -I=. -I=../../.. -I=$GOPATH/pkg/mod/github.com/gogo/protobuf@v1.3.1 --go_out=paths=source_relative:../../../../proto/golang/dm/rs trimspec_info.proto
//go:generate protoc -I=. -I=../../.. -I=$GOPATH/pkg/mod/github.com/gogo/protobuf@v1.3.1 --go_out=paths=source_relative:../../../../proto/golang/dm/rs enums.proto
//go:generate protoc -I=. -I=../../.. -I=$GOPATH/pkg/mod/github.com/gogo/protobuf@v1.3.1 --go_out=paths=source_relative:../../../../proto/golang/dm/rs container/enum_type.proto
//go:generate protoc -I=. -I=../../.. --go_out=paths=source_relative:../../../../proto/golang/dm/rs freshness.proto

//go:generate protoc -I=. -I=../../.. -I=$GOPATH/pkg/mod/github.com/gogo/protobuf@v1.3.1 -I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.11.3/third_party/googleapis --go_out=paths=source_relative,plugins=grpc:../../../../proto/golang/dm/rs service.proto
//go:generate protoc -I=. -I=../../.. -I=$GOPATH/pkg/mod/github.com/gogo/protobuf@v1.3.1 -I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.11.3/third_party/googleapis --grpc-gateway_out=paths=source_relative,logtostderr=true:../../../../proto/golang/dm/rs service.proto

//go:generate go vet ../../../../...
