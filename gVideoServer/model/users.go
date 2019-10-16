package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model  `json:"-"`
	Nickname    string    `gorm:"type:varchar(10);not null" json:"nickname"`
	Username    string    `gorm:"type:varchar(20);not null" json:"-"`
	Password    string    `gorm:"type:varchar(100);not null" json:"-"`
	Status      string    `gorm:"type:varchar(20);" json:"-"`
	Description string    `gorm:"size:1000;type:varchar(20);";json:"description"`
	Avatar      string    `gorm:"size:1000" json:"avatar"`
	Video       []Video   `json:"-"`
	Comment     []Comment `json:"-"`
}

func basicNameExists(query interface{}, name string) (*User, error) {
	var user User
	result := DB.Where(query, name).First(&user)
	fmt.Println(query, name)
	fmt.Println(user)
	if user.ID > 0 {
		return &user, nil
	}
	return nil, result.Error
}

func ExistsUserName(name string) (*User, error) {
	return basicNameExists("username = ?", name)
}

func ExistsNickname(name string) (*User, error) {
	return basicNameExists("nickname = ?", name)
}

func (u *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}
	u.Password = string(bytes)
	return nil
}

func (u *User) CheckPassword(password string) (err error) {
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err == nil {
		return nil
	}
	return err
}

func (u *User) SetNickName(name string) error {
	u.Nickname = name
	return DB.Save(u).Error
}

func GetCurrentUserById(id interface{}) (User, error) {
	var user User
	result := DB.First(&user, id)
	return user, result.Error
}

func GetCurrentUserByUsername(name string) (User, error) {
	var user User

	result := DB.Where("username = ?", name)
	return user, result.Error
}

func CreateUser(u *User) error {
	return DB.Create(&u).Error
}
