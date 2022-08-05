package server

import (
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"

	"github.com/KangTaeki/urlshortener/config"
	"github.com/KangTaeki/urlshortener/gen/go/apis/v1/urlshortener"
	"github.com/banksalad/go-banksalad"
)

type UrlshortenerServer struct {
	cfg config.Config
}

func NewUrlshortenerServer(cfg config.Config) (*UrlshortenerServer, error) {
	return &UrlshortenerServer{cfg: cfg}, nil
}

func (s *UrlshortenerServer) Config() config.Config {
	return s.cfg
}

func (s *UrlshortenerServer) RegisterServer(srv *grpc.Server) {
	urlshortener.RegisterUrlshortenerServer(srv, s)
}

func NewGRPCServer(cfg config.Config) (*grpc.Server, error) {
	logrus.ErrorKey = "grpc.error"

	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(
			banksalad.UnaryServerInterceptor(cfg.StatsdClient().CloneWithPrefixExtension(".grpc"), log),
		),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionAge: 30 * time.Second,
		}),
	)

	urlshortenerServer, err := NewUrlshortenerServer(cfg)
	if err != nil {
		return nil, err
	}

	urlshortener.RegisterUrlshortenerServer(grpcServer, urlshortenerServer)
	reflection.Register(grpcServer)

	return grpcServer, nil
}
