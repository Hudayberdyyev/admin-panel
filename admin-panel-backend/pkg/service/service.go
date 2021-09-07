package service

import (
	"github.com/Hudayberdyyev/admin-panel-backend/models"
	"github.com/Hudayberdyyev/admin-panel-backend/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(username, password string) (string, error)
	Parse(accessToken string) (int, error)
}

type News interface {
	GetAllCategories(hl string) ([]models.Category, error)
	GetAllAuthors(hl string) ([]models.Author, error)
	GetAllNewsByCategoryAndAuthorId(pagination models.Pagination, selectorQuery string, selectorParams []interface{}, hl string) ([]models.News, error)
	GetTagsByNewsTextId(newsTextId int, hl string) ([]models.Tag, error)
	GetContentByNewsTextId(newsTextId int) ([]models.NewsContent, error)
	GetNewsCountForAllAuthors() ([]models.AuthorsInfo, error)
	GetNewsCountForAllCategories(hl string) ([]models.CategoryInfo, error)
}

type ExtraMessages interface {
	Create(input models.ExtraMessages) (int, error)
	Delete(id int) error
	GetAll() ([]models.ExtraMessages, error)
}

type Service struct {
	Authorization
	News
	ExtraMessages
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		News:          NewNewsService(repos.News),
		ExtraMessages: NewExtraService(repos.ExtraMessages),
	}
}
