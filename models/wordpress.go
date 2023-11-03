package models

import "github.com/lib/pq"

type WordpressFootPrint struct {
	Model
	WebPageID     int            `json:"pageId"`
	WordpressHtml pq.StringArray `json:"wordpressHtml" gorm:"type:text[]"`
}

func (WordpressFootPrint) TableName() string {
	return "wordpress"
}
func CreateWordPressFootPrint(pageId int, html []string) *WordpressFootPrint {
	wordpress := &WordpressFootPrint{
		WebPageID:     pageId,
		WordpressHtml: html,
	}

	db.Create(wordpress)

	return wordpress
}
