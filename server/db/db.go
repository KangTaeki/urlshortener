package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"github.com/smira/go-statsd"
)

type Setting struct {
	Host            string
	Port            int
	User            string
	Password        string
	Name            string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
}

const driverName = "mysql"

func MustGetDB(s Setting) *sql.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true",
		s.User,
		s.Password,
		s.Host,
		s.Port,
		s.Name,
	)

	db, err := sql.Open(driverName, dsn)
	if err != nil {
		logrus.WithError(err).Fatal("sql.Open")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		logrus.WithError(err).Fatal("db.PingContext")
	}

	db.SetMaxIdleConns(s.MaxIdleConns)
	db.SetMaxOpenConns(s.MaxOpenConns)
	db.SetConnMaxLifetime(s.ConnMaxLifetime)

	return db
}

func NewSetting(
	host string,
	port int,
	user string,
	password string,
	name string,
	maxIdleConns int,
	maxOpenConns int,
	connMaxLifetimeSec int,
) Setting {
	return Setting{
		Host:            host,
		Port:            port,
		User:            user,
		Password:        password,
		Name:            name,
		MaxIdleConns:    maxIdleConns,
		MaxOpenConns:    maxOpenConns,
		ConnMaxLifetime: time.Duration(connMaxLifetimeSec) * time.Second,
	}
}

// StatDB stats current database statistics through a given statsd client.
// See also https://golang.org/pkg/database/sql/#DB.Stats
func StatDB(statsdCli *statsd.Client, db *sql.DB) {
	cur := db.Stats()
	statsdCli.Gauge(".pool.max", int64(cur.MaxOpenConnections))
	statsdCli.Gauge(".pool.open", int64(cur.OpenConnections))
	statsdCli.Gauge(".pool.use", int64(cur.InUse))
	statsdCli.Gauge(".pool.idle", int64(cur.Idle))
	statsdCli.Gauge(".closed.idle", cur.MaxIdleClosed)
	statsdCli.Gauge(".closed.lifetime", cur.MaxLifetimeClosed)
	statsdCli.Gauge(".wait_count", cur.WaitCount)
	statsdCli.Gauge(".wait_duration_ms", cur.WaitDuration.Milliseconds())
}
