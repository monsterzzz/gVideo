package api

import (
	"encoding/gob"
	"gVideoServer/middleware"
	"gVideoServer/model"
	v1 "gVideoServer/routers/api/v1"

	"github.com/gin-contrib/sessions"

	"github.com/gin-contrib/sessions/cookie"

	"github.com/gin-gonic/gin"
)

func InitRouters() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	r.LoadHTMLGlob("dist/*.html")        // 添加入口index.html
	r.LoadHTMLFiles("static/*/*")        // 添加资源路径
	r.Static("/static", "./dist/static") // 添加资源路径
	//r.StaticFile("/index.html", "dist/index.html") //前端接口
	r.StaticFile("/", "dist/index.html") //前端接口

	gob.Register(model.User{})

	store := cookie.NewStore([]byte("monster"))

	r.Use(sessions.Sessions("gid", store))

	//r.Static("/static", "./static")
	r.Static("/static1", "./videoRepo")

	apiV1 := r.Group("api/v1")
	{

		apiV1.POST("/login", v1.UserLogin)
		apiV1.POST("/signIn", v1.UserSign)
		apiV1.Use(middleware.AuthRequired())

		userGroup := apiV1.Group("user")
		{
			userGroup.POST("/nickname", v1.UpdateNickName)
			userGroup.POST("/description", v1.UpdateDescription)
			userGroup.GET("/videorecords", v1.GetVideoRecords)

		}

		imageGroup := apiV1.Group("image")
		{
			//imageGroup.POST("/avatar")
			//imageGroup.POST("/video")
			imageGroup.POST("/video/upload", v1.UploadImage)
			imageGroup.POST("/avatar/upload", v1.AvatarUpload)
		}
		videoGroup := apiV1.Group("video")
		{
			videoGroup.GET("/get", v1.GetAllVideo)
			videoGroup.POST("/upload", v1.UploadVideo)
			videoGroup.GET("/item/:uuid", v1.TestVideo)
		}
		videoFormGroup := apiV1.Group("videoform")
		{
			videoFormGroup.POST("/upload", v1.VideoFormUpload)
			videoFormGroup.POST("/checkmd5/:md5", v1.CheckMd5)
		}
		commentGroup := apiV1.Group("comment")
		{
			commentGroup.POST("/add", v1.AddComment)
			commentGroup.GET("/:uuid", v1.GetComments)
		}
	}
	return r

}
