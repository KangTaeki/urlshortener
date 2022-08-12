package handler

import (
	"context"
	"github.com/KangTaeki/urlshortener/gen/go/apis/v1/urlshortener"
	//"github.com/banksalad/idl/gen/go/apis/v1/urlshortener"
)

type GetLongUrlHandlerFunc func(ctx context.Context, req *urlshortener.GetLongUrlRequest) (*urlshortener.GetLongUrlResponse, error)

func GetLongUrl() GetLongUrlHandlerFunc {
	return func(ctx context.Context, req *urlshortener.GetLongUrlRequest) (*urlshortener.GetLongUrlResponse, error) {
		return &urlshortener.GetLongUrlResponse{}, nil
	}
}
