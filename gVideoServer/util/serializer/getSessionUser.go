package serializer

import (
	"gVideoServer/model"

	"github.com/gin-contrib/sessions"

	"github.com/gin-gonic/gin"
)

func GetSessionUser(c *gin.Context) *model.User {
	sessionInterface, _ := c.Get("session")
	session, _ := sessionInterface.(sessions.Session)
	user, _ := session.Get("user").(model.User)
	return &user
}
