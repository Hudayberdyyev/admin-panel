package repository

import (
	"context"
	"fmt"

	"github.com/Hudayberdyyev/admin-panel-backend/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type ExtraPostrgres struct {
	db *pgxpool.Pool
}

func NewExtraPostgres(db *pgxpool.Pool) *ExtraPostrgres {
	return &ExtraPostrgres{
		db: db,
	}
}

func (r *ExtraPostrgres) Create(input models.ExtraMessages) (int, error) {
	var id int
	query := fmt.Sprintf("insert into %s (title, description) values ($1, $2) returning id", extraTable)
	err := r.db.QueryRow(context.Background(), query, input.Title, input.Description).Scan(&id)
	return id, err
}

func (r *ExtraPostrgres) Delete(id int) error {
	query := fmt.Sprintf("delete from %s where id=$1", extraTable)
	_, err := r.db.Exec(context.Background(), query, id)
	return err
}

func (r *ExtraPostrgres) GetAll() ([]models.ExtraMessages, error) {
	var extraMsg []models.ExtraMessages
	query := fmt.Sprintf("select id, title, description, created_at from %s ", extraTable)
	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var extra models.ExtraMessages
		if err = rows.Scan(&extra.Id, &extra.Title, &extra.Description, &extra.CreatedAt); err != nil {
			continue
		}
		extraMsg = append(extraMsg, extra)
	}

	return extraMsg, nil
}
