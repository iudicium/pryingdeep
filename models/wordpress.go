package models

import "github.com/lib/pq"

type WordpressFootPrint struct {
	Model
	WebPageId     int            `json:"pageId"`
	WebPage       WebPage        `json:"page" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	WordpressHtml pq.StringArray `json:"wordpressHtml" gorm:"type:text[]"`
}

func CreateWordPressFootPrint(pageId int, html []string) *WordpressFootPrint {
	wordpress := &WordpressFootPrint{
		WebPageId:     pageId,
		WordpressHtml: html,
	}

	db.Create(wordpress)

	return wordpress
}
