package handler

import (
	"context"
	"github.com/KangTaeki/urlshortener/gen/go/apis/v1/urlshortener"
)

type HealthCheckHandlerFunc func(ctx context.Context, req *urlshortener.HealthCheckRequest) (*urlshortener.HealthCheckResponse, error)

func HealthCheck() HealthCheckHandlerFunc {
	return func(ctx context.Context, req *urlshortener.HealthCheckRequest) (*urlshortener.HealthCheckResponse, error) {
		return &urlshortener.HealthCheckResponse{}, nil
	}
}
