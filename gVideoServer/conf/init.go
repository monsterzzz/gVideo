package conf

import (
	"gVideoServer/model"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()

	model.DbConnect()

}
