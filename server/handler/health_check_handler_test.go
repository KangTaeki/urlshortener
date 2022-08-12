package handler

import (
	"context"
	"testing"

	"github.com/KangTaeki/urlshortener/gen/go/apis/v1/urlshortener"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	ctx := context.Background()
	req := &urlshortener.HealthCheckRequest{}

	resp, err := HealthCheck()(ctx, req)

	assert.NoError(t, err)
	assert.Equal(t, &urlshortener.HealthCheckResponse{}, resp)
}
