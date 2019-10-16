package middleware

import (
	"gVideoServer/util/serializer"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {

		session := sessions.Default(c)
		user := session.Get("user")
		if user == nil {
			c.JSON(200, &serializer.Response{
				Status: 99999,
				Data:   "must login",
			})

			c.Abort()
		}
		c.Set("session", session)
		c.Next()
	}
}
