package models

import "github.com/lib/pq"

type Email struct {
	Model
	WebPageId int            `json:"pageId"`
	WebPage   WebPage        `json:"page" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Emails    pq.StringArray `json:"emails" gorm:"type:text[]"`
}

func CreateEmails(pageId int, emails []string) *Email {
	email := &Email{
		WebPageId: pageId,
		Emails:    emails,
	}

	db.Create(email)

	return email
}
