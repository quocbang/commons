package context

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestWithRequestID(t *testing.T) {
	assert := assert.New(t)

	requestID := "my_request_id"
	ctx := WithRequestID(context.Background(), requestID)
	assert.Equal(requestID, RequestID(ctx))
}

func TestWithLogger(t *testing.T) {
	assert := assert.New(t)

	logger := zap.L()
	ctx := WithLogger(context.Background(), logger)
	assert.Same(logger, Logger(ctx))

	assert.NotNil(Logger(context.Background()))
}

func TestWithUserID(t *testing.T) {
	assert := assert.New(t)

	{ // success
		userID := "my_user_id"
		ctx := WithUserID(context.Background(), userID)
		assert.Equal(userID, UserID(ctx))
	}
	{ // context without userID
		assert.Empty(UserID(context.Background()))
	}
}
