package models

import "time"

type Tag struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Category struct {
	Id   int    `json:"id"`
	Text string `json:"text"`
}

type User struct {
	Id       int    `json:"-"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type News struct {
	Id          int    `json:"id"`
	NewsTextId  int    `json:"news_text_id"`
	Status      int    `json:"status"`
	ViewCount   int    `json:"view_count"`
	Title       string `json:"title"`
	Hl          string `json:"hl"`
	PublishDate string `json:"publish_date"`
	Tags        []Tag
	Categories  Category
}

type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Pagination struct {
	Page  int    `json:"page"`
	Limit int    `json:"limit"`
	Sort  string `json:"sort"`
}

type Attributes struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type NewsContent struct {
	ID    int          `json:"id"`
	Value string       `json:"value"`
	Tag   string       `json:"tag"`
	Attr  []Attributes `json:"attributes"`
}

type AuthorsInfo struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Image     string `json:"logo"`
	NewsCount int    `json:"count"`
	OpenCount int    `json:"open"`
}

type CategoryInfo struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	NewsCount int    `json:"count"`
	OpenCount int    `json:"open"`
}

type ExtraMessages struct {
	Id          int       `json:"id"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"required"`
	CreatedAt   time.Time `json:"time"`
}
