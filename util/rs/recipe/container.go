package recipe

import (
	"gitlab.kenda.com.tw/kenda/commons/v2/proto/golang/dm/rs"
	"gitlab.kenda.com.tw/kenda/commons/v2/proto/golang/dm/rs/container"
)

var mixingContainerTypes = []container.Type{
	container.Type_CC,
	container.Type_CLAY,
	container.Type_N220,
	container.Type_N330,
	container.Type_N339,
	container.Type_N550,
	container.Type_N660,
	container.Type_RECYCLE,
	container.Type_N234,
	container.Type_N326,
	container.Type_P150,
	container.Type_V500,
	container.Type_H1,
}

// ProductContainerTypes returns material types of specified product
var ProductContainerTypes = map[rs.ProductType][]container.Type{
	rs.ProductType_RUBBER: mixingContainerTypes,
}
