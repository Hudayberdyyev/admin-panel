package repository

import (
	"context"
	"fmt"

	"github.com/Hudayberdyyev/admin-panel-backend/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type AuthPostgres struct {
	db *pgxpool.Pool
}

func NewAuthPostgres(db *pgxpool.Pool) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user models.User) (int, error) {
	var id int
	query := fmt.Sprintf("insert into %s (username, password_hash) values ($1, $2) returning id", adminsTable)
	err := r.db.QueryRow(context.Background(), query, user.Username, user.Password).Scan(&id)
	return id, err
}

func (r *AuthPostgres) GetUserId(username, passwordHash string) (int, error) {
	var id int
	query := fmt.Sprintf("select id from %s where username=$1 and password_hash=$2", adminsTable)
	err := r.db.QueryRow(context.Background(), query, username, passwordHash).Scan(&id)
	return id, err
}
