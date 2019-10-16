package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"

	"os"
)

type Claims struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func ClaimsMaker(ID int, username, password string) Claims {
	return Claims{
		ID,
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + 10,
			Issuer:    "monster",
		},
	}
}

func GenerateToken(claims Claims) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		log.Println(err)
		log.Println("token generate error")
	}
	return ss
}

func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return &Claims{}, err
	}
	claims, ok := token.Claims.(*Claims)
	if !ok && !token.Valid {
		fmt.Println("!!")
		return &Claims{}, err
	}
	return claims, nil

}
