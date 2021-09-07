package service

import (
	"context"
	"strconv"

	"github.com/Hudayberdyyev/admin-panel-backend/models"
	"github.com/Hudayberdyyev/admin-panel-backend/pkg/repository"
	"github.com/Hudayberdyyev/admin-panel-backend/server"
	"github.com/Hudayberdyyev/admin-panel-backend/storage"
	"github.com/minio/minio-go/v7"
)

type NewsService struct {
	repos repository.News
}

func NewNewsService(repos repository.News) *NewsService {
	return &NewsService{repos: repos}
}

func (s *NewsService) GetAllCategories(hl string) ([]models.Category, error) {
	return s.repos.GetAllCategories(hl)
}

func (s *NewsService) GetAllAuthors(hl string) ([]models.Author, error) {
	return s.repos.GetAllAuthors(hl)
}

func (s *NewsService) GetAllNewsByCategoryAndAuthorId(pagination models.Pagination, selectorQuery string, selectorParams []interface{}, hl string) ([]models.News, error) {
	return s.repos.GetAllNewsByCategoryAndAuthorId(pagination, selectorQuery, selectorParams, hl)
}

func (s *NewsService) GetTagsByNewsTextId(newsTextId int, hl string) ([]models.Tag, error) {
	return s.repos.GetTagsByNewsTextId(newsTextId, hl)
}

func (s *NewsService) GetContentByNewsTextId(newsTextId int) ([]models.NewsContent, error) {
	contents, err := s.repos.GetContentByNewsTextId(newsTextId)
	if err != nil {
		return nil, err
	}
	for _, content := range contents {
		for iterator, v := range content.Attr {
			if v.Key == "src" {
				object, err := storage.ImageStorage.Client.GetObject(context.Background(), storage.ContentBucket, strconv.Itoa(content.ID), minio.GetObjectOptions{})
				_, err = object.Stat()
				if err != nil {
					authorId, authErr := s.repos.GetAuthorIdByNewsTextId(newsTextId)

					if authErr != nil {
						content.Attr[iterator].Value = ""
						continue
					}

					if uploadErr := storage.ImageStorage.UploadImage(context.Background(), storage.ContentBucket, v.Value, strconv.Itoa(content.ID), authorId); uploadErr != nil {
						content.Attr[iterator].Value = ""
						continue
					}
				}
				content.Attr[iterator].Value = server.AppConf.Protocol + "://" + server.AppConf.IP + ":" + server.AppConf.Port + "/image/" + storage.ContentBucketPattern + strconv.Itoa(content.ID)
			}
		}
	}
	return contents, nil
}

func (s *NewsService) GetNewsCountForAllAuthors() ([]models.AuthorsInfo, error) {
	return s.repos.GetNewsCountForAllAuthors()
}

func (s *NewsService) GetNewsCountForAllCategories(hl string) ([]models.CategoryInfo, error) {
	return s.repos.GetNewsCountForAllCategories(hl)
}
