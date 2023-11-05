package models

import "github.com/lib/pq"

// Email stores processed emails that are found within a web page.
type Email struct {
	Model
	//WebPageID is the serves as a foreign key to web_pages
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
