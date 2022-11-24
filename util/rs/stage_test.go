package rs

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"gitlab.kenda.com.tw/kenda/commons/v2/proto/golang/dm/rs"
)

func TestToErpStageName(t *testing.T) {
	{
		_, err := ToErpStageName(rs.StageStatus_STAGE_STATUS_UNSPECIFIED)
		assert.Equal(t, "bad stage=STAGE_STATUS_UNSPECIFIED", err.Error())
	}
	{
		_, err := ToErpStageName(rs.StageStatus_DEPRECATED)
		assert.Equal(t, "bad stage=DEPRECATED", err.Error())
	}
	// good cases
	{
		stage, err := ToErpStageName(rs.StageStatus_TEMPORARY)
		assert.NoError(t, err)
		assert.Equal(t, "R", stage)
	}
	{
		stage, err := ToErpStageName(rs.StageStatus_PRODUCTION)
		assert.NoError(t, err)
		assert.Equal(t, "N", stage)
	}
	{
		stage, err := ToErpStageName(rs.StageStatus_EXPERIMENTAL)
		assert.NoError(t, err)
		assert.Equal(t, "P", stage)
	}
	{
		stage, err := ToErpStageName(rs.StageStatus_TRIAL_RUN)
		assert.NoError(t, err)
		assert.Equal(t, "T", stage)
	}
	// TODO: add all enums tests
}

func TestToErpStageValue(t *testing.T) {
	{
		_, err := ToErpStageValue(rs.StageStatus_DEPRECATED)
		assert.Equal(t, "bad stage=DEPRECATED", err.Error())
	}
	{
		_, err := ToErpStageValue(rs.StageStatus_STAGE_STATUS_UNSPECIFIED)
		assert.Equal(t, "bad stage=STAGE_STATUS_UNSPECIFIED", err.Error())
	}
	// good cases
	{
		stage, err := ToErpStageValue(rs.StageStatus_PRODUCTION)
		assert.NoError(t, err)
		assert.Equal(t, 1, stage)
	}
	{
		stage, err := ToErpStageValue(rs.StageStatus_TEMPORARY)
		assert.NoError(t, err)
		assert.Equal(t, 0, stage)
	}
	{
		stage, err := ToErpStageValue(rs.StageStatus_EXPERIMENTAL)
		assert.NoError(t, err)
		assert.Equal(t, 3, stage)
	}
	{
		stage, err := ToErpStageValue(rs.StageStatus_TRIAL_RUN)
		assert.NoError(t, err)
		assert.Equal(t, 6, stage)
	}
	// TODO: add all enums tests
}
