package repository

import (
	"github.com/Hudayberdyyev/admin-panel-backend/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUserId(username, passwordHash string) (int, error)
}

type News interface {
	GetAllCategories(hl string) ([]models.Category, error)
	GetAllAuthors(hl string) ([]models.Author, error)
	GetAllNewsByCategoryAndAuthorId(pagination models.Pagination, selectorQuery string, selectorParams []interface{}, hl string) ([]models.News, error)
	GetTagsByNewsTextId(newsTextId int, hl string) ([]models.Tag, error)
	GetContentByNewsTextId(newsTextId int) ([]models.NewsContent, error)
	GetAuthorIdByNewsTextId(newsTextId int) (int, error)
	GetNewsCountForAllAuthors() ([]models.AuthorsInfo, error)
	GetNewsCountForAllCategories(hl string) ([]models.CategoryInfo, error)
}

type ExtraMessages interface {
	Create(input models.ExtraMessages) (int, error)
	Delete(id int) error
	GetAll() ([]models.ExtraMessages, error)
}

type Repository struct {
	Authorization
	News
	ExtraMessages
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		News:          NewNewsPostgres(db),
		ExtraMessages: NewExtraPostgres(db),
	}
}
