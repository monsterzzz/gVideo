package model

func Migrate() {
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Video{})
	DB.AutoMigrate(&Comment{})
}
