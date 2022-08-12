package handler

import (
	"context"
	"github.com/KangTaeki/urlshortener/gen/go/apis/v1/urlshortener"
	//"github.com/banksalad/idl/gen/go/apis/v1/urlshortener"
)

type MakeShortUrlHandlerFunc func(ctx context.Context, req *urlshortener.MakeShortUrlRequest) (*urlshortener.MakeShortUrlResponse, error)

func MakeShortUrl() MakeShortUrlHandlerFunc {
	return func(ctx context.Context, req *urlshortener.MakeShortUrlRequest) (*urlshortener.MakeShortUrlResponse, error) {
		return &urlshortener.MakeShortUrlResponse{}, nil
	}
}
