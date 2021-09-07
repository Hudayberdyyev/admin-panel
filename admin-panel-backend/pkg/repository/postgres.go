package repository

import (
	"context"
	"net/url"
	"strconv"

	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	adminsTable       = "admins"
	categoryTable     = "categories"
	categoryTextTable = "categories_text"
	authorsTable      = "authors"
	authorsTextTable  = "authors_text"
	newsTable         = "news"
	newsTextTable     = "news_text"
	tagsTable         = "tags"
	tagsTextTable     = "tags_text"
	newsTagsTable     = "news_tags"
	newsContentTable  = "news_content"
	extraTable        = "extra_messages"
	maxOpenConns      = 60
	connMaxLifetime   = 120
	maxIdleConns      = 30
	connMaxIdleTime   = 20
)

type Config struct {
	Host     string
	Port     uint16
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*pgxpool.Pool, error) {
	// db, err := pgx.Connect(pgx.ConnConfig{
	// 	Host:     cfg.Host,
	// 	Port:     cfg.Port,
	// 	Database: cfg.DBName,
	// 	User:     cfg.Username,
	// 	Password: cfg.Password,
	// })
	// postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable
	dbURL := "postgres://" + cfg.Username + ":" + url.QueryEscape(cfg.Password) + "@" + cfg.Host + ":" + strconv.Itoa(int(cfg.Port)) + "/" + cfg.DBName + "?sslmode=" + cfg.SSLMode
	dbPool, err := pgxpool.Connect(context.Background(), dbURL)

	if err != nil {
		return nil, err
	}

	// db.SetMaxOpenConns(maxOpenConns)
	// db.SetConnMaxLifetime(connMaxLifetime * time.Second)
	// db.SetMaxIdleConns(maxIdleConns)
	// db.SetConnMaxIdleTime(connMaxIdleTime * time.Second)

	err = dbPool.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	return dbPool, nil
}
