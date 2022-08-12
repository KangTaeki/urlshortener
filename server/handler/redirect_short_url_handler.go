package handler

import (
	"context"
	"github.com/KangTaeki/urlshortener/gen/go/apis/v1/urlshortener"
	//"github.com/banksalad/idl/gen/go/apis/v1/urlshortener"
)

type RedirectShortUrlHandlerFunc func(ctx context.Context, req *urlshortener.RedirectShortUrlRequest) (*urlshortener.RedirectShortUrlResponse, error)

func RedirectShortUrl() RedirectShortUrlHandlerFunc {
	return func(ctx context.Context, req *urlshortener.RedirectShortUrlRequest) (*urlshortener.RedirectShortUrlResponse, error) {
		return &urlshortener.RedirectShortUrlResponse{}, nil
	}
}
