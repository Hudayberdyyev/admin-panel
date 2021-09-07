package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/Hudayberdyyev/admin-panel-backend/models"
	"github.com/Hudayberdyyev/admin-panel-backend/server"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	FrontDateLayout = "02.01.2006"
)

type NewsPostgres struct {
	db *pgxpool.Pool
}

func NewNewsPostgres(db *pgxpool.Pool) *NewsPostgres {
	return &NewsPostgres{db: db}
}

func (r *NewsPostgres) GetAllCategories(hl string) ([]models.Category, error) {
	var categories []models.Category

	query := fmt.Sprintf("select c.id, t.title from %s c join %s t on c.id = t.category_id where t.hl=$1", categoryTable, categoryTextTable)

	rows, err := r.db.Query(context.Background(), query, hl)

	defer rows.Close()

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var cat models.Category
		if err = rows.Scan(&cat.Id, &cat.Text); err != nil {
			return nil, err
		}
		categories = append(categories, cat)
	}

	return categories, nil
}

func (r *NewsPostgres) GetAllAuthors(hl string) ([]models.Author, error) {
	var authors []models.Author
	query := fmt.Sprintf("select a.id, at.title from %s a join %s at on a.id = at.author_id where at.hl=$1", authorsTable, authorsTextTable)
	rows, err := r.db.Query(context.Background(), query, hl)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var author models.Author
		if err = rows.Scan(&author.Id, &author.Name); err != nil {
			return nil, err
		}
		authors = append(authors, author)
	}

	return authors, nil
}

func (r *NewsPostgres) GetAllNewsByCategoryAndAuthorId(pagination models.Pagination, selectorQuery string, selectorParams []interface{}, hl string) ([]models.News, error) {
	offset := (pagination.Page - 1) * pagination.Limit
	var newsList []models.News

	query := fmt.Sprintf(`select n.id, n.status, n.view_count, n.publish_date, nt.id, nt.title, nt.hl, ct.category_id, ct.title from %s n join %s nt on n.id = nt.news_id 
						join %s ct on ct.category_id = n.category_id 
						where ct.hl='%s' %s order by %s offset %d limit %d`, newsTable, newsTextTable, categoryTextTable,
		hl, selectorQuery, pagination.Sort, offset, pagination.Limit)

	rows, err := r.db.Query(context.Background(), query, selectorParams...)

	defer rows.Close()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var news models.News
		var pDate time.Time
		if err = rows.Scan(&news.Id, &news.Status, &news.ViewCount, &pDate, &news.NewsTextId, &news.Title, &news.Hl, &news.Categories.Id, &news.Categories.Text); err != nil {
			return nil, err
		}
		news.PublishDate = pDate.Format(FrontDateLayout)

		newsList = append(newsList, news)
	}

	return newsList, nil
}

func (r *NewsPostgres) GetTagsByNewsTextId(newsTextId int, hl string) ([]models.Tag, error) {
	var tags []models.Tag
	query := fmt.Sprintf(`select ttext.id, ttext.name from %s ntext join %s ntags on ntext.news_id=ntags.news_id join %s ttext on ttext.tag_id=ntags.tag_id where ntext.id=$1 and ttext.hl=$2`, newsTextTable, newsTagsTable, tagsTextTable)
	rows, err := r.db.Query(context.Background(), query, newsTextId, hl)

	defer rows.Close()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var tag models.Tag
		if err = rows.Scan(&tag.Id, &tag.Name); err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}

	return tags, nil
}

func (r *NewsPostgres) GetAuthorIdByNewsTextId(newsTextId int) (int, error) {
	var id int
	query := fmt.Sprintf("select n.author_id from %s n join %s nt on n.id = nt.news_id where nt.id=$1", newsTable, newsTextTable)
	row := r.db.QueryRow(context.Background(), query, newsTextId)
	err := row.Scan(&id)
	return id, err
}

func (r *NewsPostgres) GetContentByNewsTextId(newsTextId int) ([]models.NewsContent, error) {
	var contentList []models.NewsContent
	query := fmt.Sprintf("select nc.id, nc.value, nc.tag, nc.attr from %s nc where nc.news_text_id=$1 order by nc.id", newsContentTable)
	rows, err := r.db.Query(context.Background(), query, newsTextId)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var content models.NewsContent
		if err = rows.Scan(&content.ID, &content.Value, &content.Tag, &content.Attr); err != nil {
			return nil, err
		}
		contentList = append(contentList, content)
	}

	return contentList, nil
}

func (r *NewsPostgres) GetNewsCountForAllAuthors() ([]models.AuthorsInfo, error) {
	var authors []models.AuthorsInfo
	query := fmt.Sprintf("select a.id, a.name, a.logo , coalesce(count(n.id), 0) as news_count, coalesce(sum(n.open_count), 0) as open_count from %s as a join %s as n on n.author_id=a.id group by a.id order by a.id asc", authorsTable, newsTable)
	rows, err := r.db.Query(context.Background(), query)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var author models.AuthorsInfo
		if err := rows.Scan(&author.Id, &author.Name, &author.Image, &author.NewsCount, &author.OpenCount); err != nil {
			return nil, err
		}
		author.Image = server.AppConf.Protocol + "://" + server.AppConf.IP + ":" + server.AppConf.Port + "/logo/" + author.Image
		authors = append(authors, author)
	}

	return authors, nil
}

func (r *NewsPostgres) GetNewsCountForAllCategories(hl string) ([]models.CategoryInfo, error) {
	var categories []models.CategoryInfo
	query := fmt.Sprintf("select ct.id, ct.title , coalesce(count(n.id), 0) as news_count, coalesce(sum(n.open_count), 0) as open_count from %s as ct join %s as n on n.category_id=ct.category_id where ct.hl=$1 group by ct.id", categoryTextTable, newsTable)
	rows, err := r.db.Query(context.Background(), query, hl)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var category models.CategoryInfo
		if err := rows.Scan(&category.Id, &category.Name, &category.NewsCount, &category.OpenCount); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}
