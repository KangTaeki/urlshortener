package server

import (
	"context"

	"github.com/banksalad/urlshortener/server/handler"
	"github.com/KangTaeki/urlshortener/gen/go/apis/v1/urlshortener"
)

// verify UrlshortenerServer implements all interface methods
var _ urlshortener.UrlshortenerServer = (*UrlshortenerServer)(nil)

func (s *UrlshortenerServer) HealthCheck(ctx context.Context, req *urlshortener.HealthCheckRequest) (*urlshortener.HealthCheckResponse, error) {
	return handler.HealthCheck()(ctx, req)
}
