package rs

import (
	"errors"
	"fmt"

	"gitlab.kenda.com.tw/kenda/commons/v2/proto/golang/dm/rs"
)

var factoriesToSubCompanyAcronym = map[rs.FactoryID]string{
	rs.FactoryID_KY:  "U",
	rs.FactoryID_KU:  "U",
	rs.FactoryID_KS:  "S",
	rs.FactoryID_KC:  "C",
	rs.FactoryID_KT:  "T",
	rs.FactoryID_KJ:  "T",
	rs.FactoryID_KV:  "V",
	rs.FactoryID_KV2: "V",
	rs.FactoryID_KI:  "I",
}

//FactoriesToSubCompanyAcronym returns sub-company acronym of specified factories, return 0 if cross company.
func FactoriesToSubCompanyAcronym(factories []rs.FactoryID) (string, error) {
	if factories == nil || len(factories) == 0 {
		return "", errors.New("empty factories")
	}

	// TODO: 等之後資料改成公版後再改成map[subcompany]string
	companies := map[string]string{}
	var ca string
	var ok bool
	for _, v := range factories {
		ca, ok = factoriesToSubCompanyAcronym[v]
		if !ok {
			return "", fmt.Errorf("bad factory ID=%s", v)
		}
		companies[ca] = ca
	}

	if len(companies) > 1 {
		return "0", nil // cross company: shared spec/recipe
	}
	return ca, nil
}

var factoriesToSubCompanyValue = map[rs.FactoryID]string{
	rs.FactoryID_KY:  "1",
	rs.FactoryID_KU:  "1",
	rs.FactoryID_KS:  "2",
	rs.FactoryID_KC:  "3",
	rs.FactoryID_KT:  "7",
	rs.FactoryID_KJ:  "7",
	rs.FactoryID_KV:  "4",
	rs.FactoryID_KV2: "4",
	rs.FactoryID_KI:  "I",
}

//FactoriesToSubCompanyValue returns sub-company value of specified factories, return 0 if cross company.
func FactoriesToSubCompanyValue(factories []rs.FactoryID) (string, error) {
	if factories == nil || len(factories) == 0 {
		return "", errors.New("empty factories")
	}

	companies := map[string]string{}
	var value string
	var ok bool
	for _, v := range factories {
		value, ok = factoriesToSubCompanyValue[v]
		if !ok {
			return "", fmt.Errorf("bad factory ID=%s", v)
		}
		companies[value] = value
	}

	if len(companies) > 1 {
		return "0", nil // cross company: shared spec/recipe
	}
	return value, nil
}
