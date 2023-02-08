package v2

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
	"xqj/models"
	"xqj/pkg/e"
	"xqj/pkg/setting"
	"xqj/utils"
)

// 获取多个相亲对象信息
func GetProfiles(c *gin.Context) {
	areaCode := c.Query("areaCode")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if areaCode != "" {
		maps["name"] = areaCode
	}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	code := e.SUCCESS

	data["lists"], _ = models.GetProfiles(utils.GetPage(c), setting.AppSetting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// 添加相亲对象信息
func AddProfile(c *gin.Context) {
	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	cover_image_url := c.Query("cover_image_url")
	creater_id := c.Query("creater_id")
	create_by := c.Query("create_by")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if title != "" {
		maps["title"] = title
	}
	if desc != "" {
		maps["desc"] = desc
	}
	if content != "" {
		maps["content"] = content
	}
	if cover_image_url != "" {
		maps["cover_image_url"] = cover_image_url
	}
	if creater_id != "" {
		maps["creater_id"] = creater_id
	}
	if create_by != "" {
		maps["create_by"] = create_by
	}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}
	code := e.SUCCESS
	models.AddProfile(title, desc, content, cover_image_url, creater_id, create_by)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// 设置相亲对象信息可见性
func ModifyProfile(c *gin.Context) {
	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	cover_image_url := c.Query("cover_image_url")
	creater_id := c.Query("creater_id")
	create_by := c.Query("create_by")

	isVisiable := c.Query("isVisiable")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if isVisiable != "" {
		maps["isVisiable"] = isVisiable
	}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}
	code := e.SUCCESS
	models.ModifyProfile(title, desc, content, cover_image_url, creater_id, create_by)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// 设置相亲对象信息
func EditProfile(c *gin.Context) {
	name := c.Query("name")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	code := e.SUCCESS

	data["lists"], _ = models.GetTags(utils.GetPage(c), setting.AppSetting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
