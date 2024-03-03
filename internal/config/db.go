package middleware

import (
	"fmt"
	db "socialForumBackend/internal/database"
	"time"

	"github.com/jmoiron/sqlx"
)

type DatabaseConfig struct {
	Name                  string
	Host                  string
	User                  string
	Password              string
	Port                  int
	MaxPoolSize           int
	ReadTimeout           time.Duration
	WriteTimeout          time.Duration
	ConnectionMaxLifeTime time.Duration
}

var Database DatabaseConfig

func initDatabaseConfig() {
	Database = DatabaseConfig{
		Name:                  mustGetString("DB_NAME"),
		Host:                  mustGetString("DB_HOST"),
		User:                  mustGetString("DB_USER"),
		Password:              mustGetString("DB_PASSWORD"),
		Port:                  mustGetInt("DB_PORT"),
		MaxPoolSize:           mustGetInt("DB_POOL_SIZE"),
		ReadTimeout:           mustGetDurationMs("DB_READ_TIMEOUT_MS"),
		WriteTimeout:          mustGetDurationMs("DB_WRITE_TIMEOUT_MS"),
		ConnectionMaxLifeTime: mustGetDurationMinute("DB_CONNECTION_MAX_LIFETIME_MINUTE"),
	}
}

func (dc DatabaseConfig) ConnectionURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		dc.User,
		dc.Password,
		dc.Host,
		dc.Port,
		dc.Name,
	)
}

func InitDB() *sqlx.DB {
	dBConfig, err := db.Init(&db.Config{
		Driver:          "postgres",
		URL:             Database.ConnectionURL(),
		MaxIdleConns:    Database.MaxPoolSize,
		MaxOpenConns:    Database.MaxPoolSize,
		ConnMaxLifeTime: Database.ConnectionMaxLifeTime,
	})
	if err != nil {
		panic("failed to initialise DB : " + Database.ConnectionURL() + " : " + err.Error())
	}
	return dBConfig
}

func CloseDB() {
	err := db.Close()
	if err != nil {
		return
	}
}
