# rs系統說明

因為MongoDB是用bson格式的資料做CRUD，所以有使用到bson tag須import github.com/gogo/protobuf

## generate 例子

```go
//go:generate protoc -I=. -I=$GOPATH/pkg/mod/github.com/gogo/protobuf@v1.3.1 --gogo_out=../../../proto/golang/rs recipe.proto
```

## 沒使用到bson tag，卻需要 import gogo/protobuf的情況

以Process.proto為例，因為有使用到recipe的structure，所以解析時也是需要先import recipe.proto用到的資源才能正常解析
