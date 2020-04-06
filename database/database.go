package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"time"
	"xorm.io/xorm"
	xlog "xorm.io/xorm/log"
)

var (
	engine *xorm.Engine
)

func mustGetenv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		log.Printf("Warning: %s environment variable not set.\n", key)
	}
	return v
}

func configureConnectionPool(engine *xorm.Engine) {
	engine.SetMaxIdleConns(5)
	engine.SetMaxOpenConns(7)
	engine.SetConnMaxLifetime(1800)
}

func configureLogger(engine *xorm.Engine) {
	logger := engine.Logger()
	logger.ShowSQL(true)
	logger.SetLevel(xlog.LOG_DEBUG)
}

func configureTimezone(engine *xorm.Engine) {
	loc, err := time.LoadLocation("UTC")
	if err != nil {
		panic(err)
	}
	engine.SetTZLocation(loc)
	engine.SetTZDatabase(loc)
}

func Ping() error {
	return engine.Ping()
}

func InitDatabase() {
	dbHost := os.Getenv("DB_TCP_HOST")
	if dbHost == "" {
		err := initSocketConnectionPool()
		if err != nil {
			log.Fatalf("Socket connection is unavailable")
		}
	} else {
		err := initTcpConnectionPool()
		if err != nil {
			log.Fatalf("Tcp connection is unavailable")
		}
	}
}

func initSocketConnectionPool() error {
	var (
		err                    error
		dbUser                 = mustGetenv("DB_USER")
		dbPwd                  = mustGetenv("DB_PASS")
		instanceConnectionName = mustGetenv("INSTANCE_CONNECTION_NAME")
		dbName                 = mustGetenv("DB_NAME")
	)

	uri := fmt.Sprintf("%s:%s@unix(/cloudsql/%s)/%s", dbUser, dbPwd, instanceConnectionName, dbName)
	engine, err = xorm.NewEngine("mysql", uri)
	if err != nil {
		return fmt.Errorf("xorm.NewEngine: %v", err)
	}

	// configure settings
	configureConnectionPool(engine)
	configureTimezone(engine)
	return nil
}

func initTcpConnectionPool() error {
	var (
		err       error
		dbUser    = mustGetenv("DB_USER")
		dbPwd     = mustGetenv("DB_PASS")
		dbTcpHost = mustGetenv("DB_TCP_HOST")
		dbName    = mustGetenv("DB_NAME")
	)

	uri := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPwd, dbTcpHost, dbName)
	engine, err = xorm.NewEngine("mysql", uri)
	if err != nil {
		return fmt.Errorf("xorm.NewEngine: %v", err)
	}

	// configure settings
	configureConnectionPool(engine)
	configureLogger(engine)
	configureTimezone(engine)
	return nil
}
