package server

import (
	"context"

	"github.com/KangTaeki/urlshortener/gen/go/apis/v1/urlshortener"
	"github.com/KangTaeki/urlshortener/server/handler"
)

// verify UrlshortenerServer implements all interface methods
var _ urlshortener.UrlshortenerServer = (*UrlshortenerServer)(nil)

func (s *UrlshortenerServer) HealthCheck(ctx context.Context, req *urlshortener.HealthCheckRequest) (*urlshortener.HealthCheckResponse, error) {
	return handler.HealthCheck()(ctx, req)
}
