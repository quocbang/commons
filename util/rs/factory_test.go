package rs

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"gitlab.kenda.com.tw/kenda/commons/v2/proto/golang/dm/rs"
)

func TestFactoriesToSubCompanyAcronym(t *testing.T) {
	// bad cases
	{ // nil
		_, err := FactoriesToSubCompanyAcronym(nil)
		assert.Equal(t, errors.New("empty factories"), err)
	}
	{ // empty slice
		_, err := FactoriesToSubCompanyAcronym([]rs.FactoryID{})
		assert.Equal(t, errors.New("empty factories"), err)
	}
	{ // factoryID unspecified
		_, err := FactoriesToSubCompanyAcronym([]rs.FactoryID{rs.FactoryID_FACTORY_UNSPECIFIED})
		assert.Equal(t, errors.New("bad factory ID=FACTORY_UNSPECIFIED"), err)
	}

	// good cases
	{ // KD → U
		company, err := FactoriesToSubCompanyAcronym([]rs.FactoryID{rs.FactoryID_KY, rs.FactoryID_KU})
		assert.NoError(t, err)
		assert.Equal(t, "U", company)
	}
	{ // cross company
		company, err := FactoriesToSubCompanyAcronym([]rs.FactoryID{rs.FactoryID_KY, rs.FactoryID_KU, rs.FactoryID_KS})
		assert.NoError(t, err)
		assert.Equal(t, "0", company)
	}
	{
		company, err := FactoriesToSubCompanyAcronym([]rs.FactoryID{rs.FactoryID_KC})
		assert.NoError(t, err)
		assert.Equal(t, "C", company)
	}
	{
		company, err := FactoriesToSubCompanyAcronym([]rs.FactoryID{rs.FactoryID_KY})
		assert.NoError(t, err)
		assert.Equal(t, "U", company)
	}
	{
		company, err := FactoriesToSubCompanyAcronym([]rs.FactoryID{rs.FactoryID_KS})
		assert.NoError(t, err)
		assert.Equal(t, "S", company)
	}
	{
		company, err := FactoriesToSubCompanyAcronym([]rs.FactoryID{rs.FactoryID_KV2})
		assert.NoError(t, err)
		assert.Equal(t, "V", company)
	}
	{
		company, err := FactoriesToSubCompanyAcronym([]rs.FactoryID{rs.FactoryID_KT})
		assert.NoError(t, err)
		assert.Equal(t, "T", company)
	}
	{
		company, err := FactoriesToSubCompanyAcronym([]rs.FactoryID{rs.FactoryID_KI})
		assert.NoError(t, err)
		assert.Equal(t, "I", company)
	}
	// TODO: add all cases
}

func TestFactoriesToSubCompanyValue(t *testing.T) {
	// bad cases
	{ // nil
		_, err := FactoriesToSubCompanyValue(nil)
		assert.Equal(t, errors.New("empty factories"), err)
	}
	{ // empty slice
		_, err := FactoriesToSubCompanyValue([]rs.FactoryID{})
		assert.Equal(t, errors.New("empty factories"), err)
	}
	{ // factoryID unspecified
		_, err := FactoriesToSubCompanyValue([]rs.FactoryID{rs.FactoryID_FACTORY_UNSPECIFIED})
		assert.Equal(t, errors.New("bad factory ID=FACTORY_UNSPECIFIED"), err)
	}

	// good cases
	{ // KD → U
		company, err := FactoriesToSubCompanyValue([]rs.FactoryID{rs.FactoryID_KY, rs.FactoryID_KU})
		assert.NoError(t, err)
		assert.Equal(t, "1", company)
	}
	{ // cross company
		company, err := FactoriesToSubCompanyValue([]rs.FactoryID{rs.FactoryID_KY, rs.FactoryID_KU, rs.FactoryID_KS})
		assert.NoError(t, err)
		assert.Equal(t, "0", company)
	}
	{
		company, err := FactoriesToSubCompanyValue([]rs.FactoryID{rs.FactoryID_KC})
		assert.NoError(t, err)
		assert.Equal(t, "3", company)
	}
	{
		company, err := FactoriesToSubCompanyValue([]rs.FactoryID{rs.FactoryID_KY})
		assert.NoError(t, err)
		assert.Equal(t, "1", company)
	}
	{
		company, err := FactoriesToSubCompanyValue([]rs.FactoryID{rs.FactoryID_KS})
		assert.NoError(t, err)
		assert.Equal(t, "2", company)
	}
	{
		company, err := FactoriesToSubCompanyValue([]rs.FactoryID{rs.FactoryID_KV2})
		assert.NoError(t, err)
		assert.Equal(t, "4", company)
	}
	{
		company, err := FactoriesToSubCompanyValue([]rs.FactoryID{rs.FactoryID_KT})
		assert.NoError(t, err)
		assert.Equal(t, "7", company)
	}
	{
		company, err := FactoriesToSubCompanyValue([]rs.FactoryID{rs.FactoryID_KI})
		assert.NoError(t, err)
		assert.Equal(t, "I", company)
	}
	// TODO: add all cases
}
