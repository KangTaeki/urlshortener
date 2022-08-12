package server

import (
	"context"
	"github.com/KangTaeki/urlshortener/gen/go/apis/v1/urlshortener"
	"github.com/KangTaeki/urlshortener/server/handler"
	//	"github.com/banksalad/idl/gen/go/apis/v1/urlshortener"
)

// verify UrlshortenerServer implements all interface methods
var _ urlshortener.UrlshortenerServer = (*UrlshortenerServer)(nil)

func (s *UrlshortenerServer) HealthCheck(ctx context.Context, req *urlshortener.HealthCheckRequest) (*urlshortener.HealthCheckResponse, error) {
	return handler.HealthCheck()(ctx, req)
}

func (s *UrlshortenerServer) MakeShortUrl(ctx context.Context, req *urlshortener.MakeShortUrlRequest) (*urlshortener.MakeShortUrlResponse, error) {
	return handler.MakeShortUrl()(ctx, req)
}

func (s *UrlshortenerServer) GetLongUrl(ctx context.Context, req *urlshortener.GetLongUrlRequest) (*urlshortener.GetLongUrlResponse, error) {
	return handler.GetLongUrl()(ctx, req)
}

func (s *UrlshortenerServer) RedirectShortUrl(ctx context.Context, req *urlshortener.RedirectShortUrlRequest) (*urlshortener.RedirectShortUrlResponse, error) {
	return handler.RedirectShortUrl()(ctx, req)
}
