package routers

import (
	"github.com/gin-gonic/gin"
	"xqj/docs"
	v2 "xqj/routers/api/v2"

	"net/http"
	"xqj/middleware/jwt"
	"xqj/pkg/export"
	"xqj/pkg/qrcode"
	"xqj/pkg/upload"
	"xqj/routers/api"

	"github.com/dchest/captcha"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	v1 "xqj/routers/api/v1"
)

var captchaHandler = captcha.Server(100, 40)

func InitRouter() *gin.Engine {
	r := gin.New()

	docs.SwaggerInfo.BasePath = "/"

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	r.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.GET("/captcha/*", api.CreateUser)

	r.GET("/createuser", api.CreateUser)
	r.GET("/createuser", api.CreateUser)
	r.GET("/auth", api.GetAuth)
	//上传图片
	r.POST("/upload", api.UploadImage)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
		//导出标签
		r.POST("/tags/export", v1.ExportTag)
		//导入标签
		r.POST("/tags/import", v1.ImportTag)
		//获取文章列表
		apiv1.GET("/articles", v1.GetArticles)
		//获取指定文章
		apiv1.GET("/articles/:id", v1.GetArticle)
		//新建文章
		apiv1.POST("/articles", v1.AddArticle)
		//更新指定文章
		apiv1.PUT("/articles/:id", v1.EditArticle)
		//删除指定文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
		//生成二维码
		apiv1.POST("/articles/poster/generate", v1.GenerateArticlePoster)
	}

	apiv2 := r.Group("/api/v2")
	apiv2.Use(jwt.JWT())
	{
		//apiv2.GET("/profiles/", v2.GetProfiles)
		apiv2.POST("/profile/modify", v2.ModifyProfile)

		//更新相亲对象信息
		apiv2.POST("/profile/add", v2.AddProfile)

		//获取所有相亲对象列表信息
		apiv2.GET("/profile/getall", v2.GetProfiles)

		//根据地区获取所有相亲对象列表信息
		apiv2.GET("/profile/getbycity", v2.GetProfiles)

	}

	return r
}
