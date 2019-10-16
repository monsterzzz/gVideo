package model

import (
	"github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model `json:"-"`
	VideoID    int    `json:"-"`
	UserID     int    `json:"-"`
	Content    string `json:"content"`
	Time       string `json:"time"`
}

func (c *Comment) AfterFind() (err error) {
	c.Time = c.CreatedAt.Format("2006-01-02 15:04:05")

	return nil
}
