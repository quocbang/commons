package rs

import "gitlab.kenda.com.tw/kenda/commons/v2/proto/golang/dm/rs"

// IsRawRubber returns true if product type is raw-rubber (10~29: nature rubber, IIR ...)
func IsRawRubber(productType rs.ProductType) bool {
	if productType >= rs.ProductType_NATURAL_RUBBER &&
		productType < rs.ProductType_ANTIOXIDANT {
		return true
	}
	return false
}

// IsCompoundingIngredient returns true if product type is compounding ingredient (30~49)
func IsCompoundingIngredient(productType rs.ProductType) bool {
	if productType >= rs.ProductType_ANTIOXIDANT && productType < rs.ProductType_VULCANIZING_AGENT {
		return true
	}
	return false
}

// IsAcceleratorIngredient returns true if product type is compounding ingredient (50~59)
func IsAcceleratorIngredient(productType rs.ProductType) bool {
	if productType >= rs.ProductType_VULCANIZING_AGENT && productType < rs.ProductType_REINFORCING_AGENT {
		return true
	}
	return false
}

// IsMixingProcessOil returns true if product type is mixing process oil (68)
func IsMixingProcessOil(productType rs.ProductType) bool {
	if productType == rs.ProductType_RUBBER_PROCESS_OIL {
		return true
	}
	return false
}

// IsReinforcingFiller returns true if product type is reinforcing filler (60~67,69: carbon, withe carbon ...)
func IsReinforcingFiller(productType rs.ProductType) bool {
	if productType >= rs.ProductType_REINFORCING_AGENT &&
		productType < rs.ProductType_MATERIAL_PLY &&
		productType != rs.ProductType_RUBBER_PROCESS_OIL {
		return true
	}
	return false
}

var ingredientMaterialTypes = []rs.ProductType{
	rs.ProductType_ANTIOXIDANT,
	rs.ProductType_ACTIVATOR,
	rs.ProductType_TACKIFIER,
	rs.ProductType_ADHESIVE,
	rs.ProductType_PROCESSING_AID,
	rs.ProductType_PIGMENT,
	rs.ProductType_PERFUME,
	rs.ProductType_VULCANIZING_AGENT,
	rs.ProductType_ACCELERATOR,
	rs.ProductType_RETARDER,
	rs.ProductType_FOAMING_AGENT,
}

// ProductMaterialTypes returns material types of specified product
var ProductMaterialTypes = map[rs.ProductType][]rs.ProductType{
	rs.ProductType_RUBBER: []rs.ProductType{
		rs.ProductType_NATURAL_RUBBER,
		rs.ProductType_SYNTHETIC_RUBBER,
		rs.ProductType_PU_FOAM,
		rs.ProductType_BUTYL_RUBBER,
		rs.ProductType_RECLAIMED_RUBBER,
		rs.ProductType_ANTIOXIDANT,
		rs.ProductType_ACTIVATOR,
		rs.ProductType_TACKIFIER,
		rs.ProductType_ADHESIVE,
		rs.ProductType_PROCESSING_AID,
		rs.ProductType_PIGMENT,
		rs.ProductType_PERFUME,
		rs.ProductType_VULCANIZING_AGENT,
		rs.ProductType_ACCELERATOR,
		rs.ProductType_RETARDER,
		rs.ProductType_FOAMING_AGENT,
		rs.ProductType_REINFORCING_AGENT,
		rs.ProductType_FILLER,
		rs.ProductType_RUBBER_PROCESS_OIL,
		rs.ProductType_KEVLAR_BRAND_ENGINEERED_ELASTOMER,
		rs.ProductType_RUBBER,
		rs.ProductType_COMPOUNDING_INGREDIENT,
		rs.ProductType_ACCELERATOR_INGREDIENT,
		rs.ProductType_RETURN_COMPOUND,
	},
	rs.ProductType_COMPOUNDING_INGREDIENT: ingredientMaterialTypes,
	rs.ProductType_ACCELERATOR_INGREDIENT: ingredientMaterialTypes,
}
