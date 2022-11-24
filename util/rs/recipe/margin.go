package recipe

import (
	"errors"
	"fmt"

	"gitlab.kenda.com.tw/kenda/commons/v2/proto/golang/dm/rs"
	"gitlab.kenda.com.tw/kenda/commons/v2/proto/golang/mes/v2/commons"

	"gitlab.kenda.com.tw/kenda/commons/v2/util/decimal"
	rsUtil "gitlab.kenda.com.tw/kenda/commons/v2/util/rs"
)

// ProductErrorMargin returns the lowest and highest acceptable values that are different from actual fed.
// For materials of mixing process,it's depends on fed weight, please make sure input quantity is in kilogram (kg).
func ProductErrorMargin(productType rs.ProductType, quantity *commons.Decimal) (min *commons.Decimal, max *commons.Decimal, err error) {
	if productType == rs.ProductType_PRODUCT_TYPE_UNSPECIFIED {
		return nil, nil, errors.New("missing product type")
	}
	if quantity == nil {
		return nil, nil, errors.New("missing quantity")
	}
	qty, ok := decimal.ToFloat(quantity, quantity.GetExp())
	if !ok {
		return nil, nil, fmt.Errorf("bad quantity=%v: error when converting decimal number", quantity)
	}
	if qty < 0 {
		return nil, nil, fmt.Errorf("bad quantity=%v: less than zero", quantity)
	}

	var margin float64
	switch {
	case rsUtil.IsRawRubber(productType) || productType == rs.ProductType_RUBBER:
		margin = MarginRawRubber()
	case rsUtil.IsCompoundingIngredient(productType):
		margin, err = MarginCompoundingIngredient(qty)
	case rsUtil.IsAcceleratorIngredient(productType):
		margin, err = MarginAcceleratorIngredient(qty)
	case rsUtil.IsReinforcingFiller(productType):
		margin, err = MarginReinforcingFiller(qty)
	case rsUtil.IsMixingProcessOil(productType):
		margin, err = MarginMixingProcessOil(qty)
	default:
		return nil, nil, fmt.Errorf("fail to get error margin: unsupported product type=%s", productType)
	}

	if err != nil {
		return nil, nil, fmt.Errorf("fail to get error margin: product type=%s, error=%s", productType, err)
	}
	min, err = decimal.FromFloat(qty - margin)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to convert min value to *Decimail type: %v", err)
	}
	max, err = decimal.FromFloat(qty + margin)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to convert max value to *Decimail type: %v", err)
	}
	return min, max, nil
}

// MarginRawRubber returns margin of raw-rubber
func MarginRawRubber() (margin float64) {
	return 0.5
}

// MarginCompoundingIngredient returns error margin of compounding ingredient
func MarginCompoundingIngredient(kg float64) (margin float64, err error) {
	if kg < 0 {
		return margin, fmt.Errorf("bad value=%v", kg)
	} else if kg >= 0 && kg <= 5 {
		return 0.005, nil
	} else if kg >= 6 && kg <= 15 {
		return 0.05, nil
	} else if kg > 15 {
		return 0.5, nil
	} else {
		return margin, fmt.Errorf("value=%v out of range", kg)
	}
}

// MarginAcceleratorIngredient returns error margin of accelerator ingredient
func MarginAcceleratorIngredient(kg float64) (margin float64, err error) {
	if kg < 0 {
		return margin, fmt.Errorf("bad value=%v", kg)
	} else if kg >= 0 && kg <= 5 {
		return 0.005, nil
	} else if kg >= 6 && kg <= 15 {
		return 0.01, nil
	} else if kg > 15 {
		return 0.02, nil
	} else {
		return margin, fmt.Errorf("value=%v out of range", kg)
	}
}

// MarginMixingProcessOil returns error margin of mixing process oil
func MarginMixingProcessOil(kg float64) (margin float64, err error) {
	if kg < 0 {
		return margin, fmt.Errorf("bad value=%v", kg)
	} else if kg >= 0 && kg <= 5 {
		return 0.15, nil
	} else if kg >= 6 && kg <= 15 {
		return 0.3, nil
	} else if kg > 15 {
		return 0.5, nil
	} else {
		return margin, fmt.Errorf("value=%v out of range", kg)
	}
}

// MarginReinforcingFiller returns error margin of reinforcing filler (60~67,69: carbon, withe carbon ...)
func MarginReinforcingFiller(kg float64) (margin float64, err error) {
	if kg < 0 {
		return margin, fmt.Errorf("bad value=%v", kg)
	}
	if kg >= 0 && kg <= 5 {
		return 0.1, nil
	}
	if kg >= 6 && kg <= 15 {
		return 0.3, nil
	}
	if kg > 15 {
		return 0.5, nil
	}
	return margin, fmt.Errorf("value=%v out of range", kg)
}
