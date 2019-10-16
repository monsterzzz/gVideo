package model

import "github.com/jinzhu/gorm"

type Video struct {
	gorm.Model  `json:"-"`
	UserID      uint      `json:"-"`
	Title       string    `json:"title"`
	Uuid        string    `json:"uuid"`
	Path        string    `json:"-"`
	Description string    `json:"description"`
	Md5         string    `json:"-"`
	Picture     string    `json:"picture"`
	PlayCount   int       `gorm:"default:0" json:"play_count"`
	LikesCount  int       `gorm:"default:0" json:"likes_count"`
	IsComplete  int       `gorm:"default:0" json:"-"`
	Comment     []Comment `json:"comment"`
	Time        string    `gorm:"-" json:"time"`
}

func (v *Video) AfterFind() {
	v.Time = v.CreatedAt.Format("2006-01-02 15:04:05")
}
