package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"github.com/banksalad/urlshortener/client"
	"github.com/banksalad/urlshortener/config"
	"github.com/banksalad/urlshortener/server"
	"github.com/banksalad/urlshortener/server/db"
	"github.com/banksalad/go-banksalad"
)

func main() {
	ctx := context.Background()

	logrus.SetFormatter(&logrus.JSONFormatter{})
	log := logrus.StandardLogger()

	setting := config.NewSetting()

	statsdCli := banksalad.NewStatsdClient(banksalad.StatsdOptions{
		StatsdAddr: fmt.Sprintf("localhost:%s", setting.StatsdUDPPort),
		Prefix:     fmt.Sprintf("%s.%s", setting.Namespace, setting.ServiceName),
		Logger:     log,
	})
	defer func() {
		if err := statsdCli.Close(); err != nil {
			log.WithError(err).Error("failed to close statsd client")
		}
	}()

	logrus.AddHook(banksalad.NewLogLevelStatHook(statsdCli))
	logrus.AddHook(banksalad.NewLogStacktraceHook())

	closeTracer, err := banksalad.InitializeTracer(banksalad.TracerOptions{
		ServiceName: setting.ServiceName,
		Namespace:   setting.Namespace,
		Env:         setting.Env,
	})
	if err != nil {
		log.WithError(err).Fatal("failed to initialize jaeger client")
	}
	defer closeTracer()

	readWriteDB := db.MustGetDB(db.NewSetting(
		setting.ReadWriteDBHost,
		setting.DBPort,
		setting.DBUser,
		setting.DBPassword,
		setting.DBName,
		setting.DBMaxIdleConns,
		setting.DBMaxOpenConns,
		setting.DBConnMaxLifetimeSec,
	))
	defer func() {
		if err := readWriteDB.Close(); err != nil {
			log.WithError(err).Error("failed to close read write db")
		}
	}()
	go func() {
		for range time.Tick(time.Duration(setting.DBStatsPeriodMs) * time.Millisecond) {
			db.StatDB(statsdCli.CloneWithPrefixExtension(".mysql.rw"), readWriteDB)
		}
	}()

	readOnlyDB := db.MustGetDB(db.NewSetting(
		setting.ReadOnlyDBHost,
		setting.DBPort,
		setting.DBUser,
		setting.DBPassword,
		setting.DBName,
		setting.DBMaxIdleConns,
		setting.DBMaxOpenConns,
		setting.DBConnMaxLifetimeSec,
	))
	defer func() {
		if err := readOnlyDB.Close(); err != nil {
			log.WithError(err).Error("failed to close read only db")
		}
	}()
	go func() {
		for range time.Tick(time.Duration(setting.DBStatsPeriodMs) * time.Millisecond) {
			db.StatDB(statsdCli.CloneWithPrefixExtension(".mysql.ro"), readOnlyDB)
		}
	}()

	authCli := client.GetAuthClient(setting.AuthGRPCEndpoint, setting.ServiceName)

	cfg := config.NewConfig(
		setting,
		statsdCli,
		readWriteDB,
		readOnlyDB,
		authCli,
	)

	grpcServer, err := server.NewGRPCServer(cfg)
	if err != nil {
		log.Fatal(err)
	}

	httpServer, err := server.NewHTTPServer(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		lis, err := net.Listen("tcp", ":"+cfg.Setting().GRPCServerPort)
		if err != nil {
			log.WithError(err).Fatal("net.Listen")
		}

		log.WithField("port", cfg.Setting().GRPCServerPort).Info("starting urlshortener gRPC server")
		if err := grpcServer.Serve(lis); err != nil && err != grpc.ErrServerStopped {
			log.Fatal(err)
		}
	}()

	go func() {
		log.WithField("port", cfg.Setting().HTTPServerPort).Info("starting urlshortener HTTP server")
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(quit)

	<-quit

	time.Sleep(time.Duration(cfg.Setting().GracefulShutdownTimeoutMs) * time.Millisecond)

	log.Info("stopping urlshortener HTTP server")
	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	log.Info("stopping urlshortener gRPC server")
	grpcServer.GracefulStop()
}
