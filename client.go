package urlshortener

import (
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/banksalad/go-banksalad"
	"github.com/KangTaeki/urlshortener/gen/go/apis/v1/urlshortener"
)

//go:generate mockgen -package urlshortener -destination ./mock_client.go -mock_names UrlshortenerClient=MockUrlshortenerClient github.com/KangTaeki/urlshortener/gen/go/apis/v1/urlshortener UrlshortenerClient
const serviceConfig = `{"loadBalancingPolicy":"round_robin"}`

var (
	once sync.Once
	cli  urlshortener.UrlshortenerClient

	// verify MockUrlshortenerClient implements all UrlshortenerClient interface methods
	_ urlshortener.UrlshortenerClient = (*MockUrlshortenerClient)(nil)
)

func GetUrlshortenerClient(serviceHost, caller string) urlshortener.UrlshortenerClient {
	once.Do(func() {
		conn, _ := grpc.Dial(
			serviceHost,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithDefaultServiceConfig(serviceConfig),
			grpc.WithUnaryInterceptor(banksalad.UnaryClientInterceptor(caller)),
		)

		cli = urlshortener.NewUrlshortenerClient(conn)
	})

	return cli
}
