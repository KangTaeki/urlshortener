package server

import (
	"context"
	"net/http"

	"github.com/KangTaeki/urlshortener/config"
	"github.com/KangTaeki/urlshortener/gen/go/apis/v1/urlshortener"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/banksalad/go-banksalad/grpcgateway/v2"
)

func NewHTTPServer(ctx context.Context, cfg config.Config) (*http.Server, error) {
	mux := runtime.NewServeMux(
		runtime.WithMarshalerOption(
			runtime.MIMEWildcard,
			&runtime.JSONPb{
				MarshalOptions: protojson.MarshalOptions{
					UseProtoNames:   true,
					EmitUnpopulated: true,
				},
				UnmarshalOptions: protojson.UnmarshalOptions{
					DiscardUnknown: true,
				},
			},
		),
		runtime.WithForwardResponseOption(convertHTTPStatusCodeByProto),
		runtime.WithErrorHandler(grpcgateway.ExternalHTTPErrorHandler),
		runtime.WithIncomingHeaderMatcher(grpcgateway.IncomingHTTPHeaderMatcher()),
		runtime.WithOutgoingHeaderMatcher(grpcgateway.OutgoingHTTPHeaderMatcher()),
	)
	options := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	if err := urlshortener.RegisterUrlshortenerHandlerFromEndpoint(
		ctx,
		mux,
		cfg.Setting().GRPCServerEndpoint,
		options,
	); err != nil {
		return nil, err
	}

	server := &http.Server{
		Addr:    ":" + cfg.Setting().HTTPServerPort,
		Handler: mux,
	}

	return server, nil
}

func convertHTTPStatusCodeByProto(_ context.Context, w http.ResponseWriter, p proto.Message) error {
	w.Header().Set("Location", p.RedirectUrl)
	w.WriteHeader(http.StatusFound)

	return nil
}
