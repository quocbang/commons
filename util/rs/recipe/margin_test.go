package recipe

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"gitlab.kenda.com.tw/kenda/commons/v2/proto/golang/dm/rs"
	"gitlab.kenda.com.tw/kenda/commons/v2/proto/golang/mes/v2/commons"
)

func TestProductErrorMargin(t *testing.T) {
	// Normal Use Cases
	{ // raw rubber
		min, max, err := ProductErrorMargin(rs.ProductType_NATURAL_RUBBER, &commons.Decimal{Value: 1, Exp: 0})
		assert.NoError(t, err)
		assert.Equal(t, &commons.Decimal{Value: 5, Exp: -1}, min)
		assert.Equal(t, &commons.Decimal{Value: 15, Exp: -1}, max)
	}
	{ // compounding ingredient (0~5kg)
		min, max, err := ProductErrorMargin(rs.ProductType_ANTIOXIDANT, &commons.Decimal{Value: 1, Exp: 0})
		assert.NoError(t, err)
		assert.Equal(t, &commons.Decimal{Value: 995, Exp: -3}, min)
		assert.Equal(t, &commons.Decimal{Value: 1005, Exp: -3}, max)
	}
	{ // compounding ingredient (6~15kg)
		min, max, err := ProductErrorMargin(rs.ProductType_ANTIOXIDANT, &commons.Decimal{Value: 6, Exp: 0})
		assert.NoError(t, err)
		assert.Equal(t, &commons.Decimal{Value: 595, Exp: -2}, min)
		assert.Equal(t, &commons.Decimal{Value: 605, Exp: -2}, max)
	}
	{ // compounding ingredient (> 15kg)
		min, max, err := ProductErrorMargin(rs.ProductType_ANTIOXIDANT, &commons.Decimal{Value: 15, Exp: 0})
		assert.NoError(t, err)
		assert.Equal(t, &commons.Decimal{Value: 1495, Exp: -2}, min)
		assert.Equal(t, &commons.Decimal{Value: 1505, Exp: -2}, max)
	}
	{ // accelerator ingredient (0~5kg)
		min, max, err := ProductErrorMargin(rs.ProductType_ACCELERATOR, &commons.Decimal{Value: 1, Exp: 0})
		assert.NoError(t, err)
		assert.Equal(t, &commons.Decimal{Value: 995, Exp: -3}, min)
		assert.Equal(t, &commons.Decimal{Value: 1005, Exp: -3}, max)
	}
	{ // accelerator ingredient (6~15kg)
		min, max, err := ProductErrorMargin(rs.ProductType_ACCELERATOR, &commons.Decimal{Value: 6, Exp: 0})
		assert.NoError(t, err)
		assert.Equal(t, &commons.Decimal{Value: 599, Exp: -2}, min)
		assert.Equal(t, &commons.Decimal{Value: 601, Exp: -2}, max)
	}
	{ // accelerator ingredient (> 15kg)
		min, max, err := ProductErrorMargin(rs.ProductType_ACCELERATOR, &commons.Decimal{Value: 15, Exp: 0})
		assert.NoError(t, err)
		assert.Equal(t, &commons.Decimal{Value: 1499, Exp: -2}, min)
		assert.Equal(t, &commons.Decimal{Value: 1501, Exp: -2}, max)
	}
	{ // reinforcing filler (0~5kg)
		min, max, err := ProductErrorMargin(rs.ProductType_REINFORCING_AGENT, &commons.Decimal{Value: 1, Exp: 0})
		assert.NoError(t, err)
		assert.Equal(t, &commons.Decimal{Value: 9, Exp: -1}, min)
		assert.Equal(t, &commons.Decimal{Value: 11, Exp: -1}, max)
	}
	{ // reinforcing filler (6~15kg)
		min, max, err := ProductErrorMargin(rs.ProductType_REINFORCING_AGENT, &commons.Decimal{Value: 6, Exp: 0})
		assert.NoError(t, err)
		assert.Equal(t, &commons.Decimal{Value: 57, Exp: -1}, min)
		assert.Equal(t, &commons.Decimal{Value: 63, Exp: -1}, max)
	}
	{ // reinforcing filler (> 15kg)
		min, max, err := ProductErrorMargin(rs.ProductType_REINFORCING_AGENT, &commons.Decimal{Value: 15, Exp: 0})
		assert.NoError(t, err)
		assert.Equal(t, &commons.Decimal{Value: 147, Exp: -1}, min)
		assert.Equal(t, &commons.Decimal{Value: 153, Exp: -1}, max)
	}
	{ // mixing process oil (0~5kg)
		min, max, err := ProductErrorMargin(rs.ProductType_RUBBER_PROCESS_OIL, &commons.Decimal{Value: 1, Exp: 0})
		assert.NoError(t, err)
		assert.Equal(t, &commons.Decimal{Value: 85, Exp: -2}, min)
		assert.Equal(t, &commons.Decimal{Value: 115, Exp: -2}, max)
	}
	{ // mixing process oil (6~15kg)
		min, max, err := ProductErrorMargin(rs.ProductType_RUBBER_PROCESS_OIL, &commons.Decimal{Value: 6, Exp: 0})
		assert.NoError(t, err)
		assert.Equal(t, &commons.Decimal{Value: 57, Exp: -1}, min)
		assert.Equal(t, &commons.Decimal{Value: 63, Exp: -1}, max)
	}
	{ // mixing process oil (> 15kg)
		min, max, err := ProductErrorMargin(rs.ProductType_RUBBER_PROCESS_OIL, &commons.Decimal{Value: 16, Exp: 0})
		assert.NoError(t, err)
		assert.Equal(t, &commons.Decimal{Value: 155, Exp: -1}, min)
		assert.Equal(t, &commons.Decimal{Value: 165, Exp: -1}, max)
	}
	// bad use cases
	{ // missing product type
		_, _, err := ProductErrorMargin(rs.ProductType_PRODUCT_TYPE_UNSPECIFIED, nil)
		assert.Error(t, err)
	}
	{ // missing quantity
		_, _, err := ProductErrorMargin(rs.ProductType_NATURAL_RUBBER, nil)
		assert.Error(t, err)
	}
	{ // bad quantity
		_, _, err := ProductErrorMargin(rs.ProductType_NATURAL_RUBBER, &commons.Decimal{Value: -1})
		assert.Error(t, err)
	}
	{ // unsupported product type
		_, _, err := ProductErrorMargin(rs.ProductType_PARTS, &commons.Decimal{Value: 5, Exp: 0})
		assert.Error(t, err)
	}
}
