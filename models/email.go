package models

import "github.com/lib/pq"

type Email struct {
	Model
	WebPageID int            `json:"pageId"`
	Emails    pq.StringArray `json:"emails" gorm:"type:text[]"`
}

func CreateEmails(pageId int, emails []string) *Email {
	email := &Email{
		WebPageID: pageId,
		Emails:    emails,
	}

	db.Create(email)

	return email
}
