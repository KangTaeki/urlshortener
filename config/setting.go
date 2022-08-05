package config

import "github.com/banksalad/go-banksalad"

type Setting struct {
	ServiceName        string
	GRPCServerEndpoint string
	GRPCServerPort     string
	HTTPServerPort     string
	StatsdUDPPort      string

	Env       string
	SubEnvID  string
	Namespace string

	GracefulShutdownTimeoutMs int

	DBName               string
	ReadOnlyDBHost       string
	ReadWriteDBHost      string
	DBPort               int
	DBUser               string
	DBPassword           string
	DBSSLMode            string
	DBMaxIdleConns       int
	DBMaxOpenConns       int
	DBConnMaxLifetimeSec int
	DBStatsPeriodMs      int

	AuthGRPCEndpoint string
}

func NewSetting() Setting {
	return Setting{
		ServiceName:        "urlshortener",
		GRPCServerEndpoint: banksalad.MustGetEnv("GRPC_SERVER_ENDPOINT", "localhost:18081"),
		GRPCServerPort:     banksalad.MustGetEnv("GRPC_SERVER_PORT", "18081"),
		HTTPServerPort:     banksalad.MustGetEnv("HTTP_SERVER_PORT", "18082"),
		StatsdUDPPort:      banksalad.MustGetEnv("STATSD_UDP_PORT", "8125"),

		Env:       banksalad.MustGetEnv("ENV", "development"),
		SubEnvID:  banksalad.MustGetEnv("SUB_ENV_ID", "local"),
		Namespace: banksalad.MustGetEnv("NAMESPACE", "development-local"),

		GracefulShutdownTimeoutMs: banksalad.MustAtoi(banksalad.MustGetEnv("GRACEFUL_SHUTDOWN_TIMEOUT_MS", "10000")),

		DBName:               banksalad.MustGetEnv("DB_NAME", "urlshortener"),
		ReadOnlyDBHost:       banksalad.MustGetEnv("READ_ONLY_DB_HOST", "localhost"),
		ReadWriteDBHost:      banksalad.MustGetEnv("READ_WRITE_DB_HOST", "localhost"),
		DBPort:               banksalad.MustAtoi(banksalad.MustGetEnv("DB_PORT", "3306")),
		DBUser:               banksalad.MustGetEnv("DB_USER", "urlshortener_was"),
		DBPassword:           banksalad.MustGetEnv("DB_PASSWORD", ""),
		DBSSLMode:            banksalad.MustGetEnv("DB_SSL_MODE", "false"),
		DBMaxIdleConns:       banksalad.MustAtoi(banksalad.MustGetEnv("DB_MAX_IDLE_CONNS", "")),
		DBMaxOpenConns:       banksalad.MustAtoi(banksalad.MustGetEnv("DB_MAX_OPEN_CONNS", "")),
		DBConnMaxLifetimeSec: banksalad.MustAtoi(banksalad.MustGetEnv("DB_CONN_MAX_LIFETIME_SEC", "600")),
		DBStatsPeriodMs:      banksalad.MustAtoi(banksalad.MustGetEnv("DB_STATS_PERIOD_MS", "500")),

		AuthGRPCEndpoint: banksalad.MustGetEnv("AUTH_GRPC_ENDPOINT", "dns:///auth-headless:18081"),
	}
}
