package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"image/color"
	"log"
	"net/http"
	"xqj/pkg/e"
)

// configJsonBody json request body.
type configJsonBody struct {
	Id            string
	CaptchaType   string
	VerifyValue   string
	DriverAudio   *base64Captcha.DriverAudio
	DriverString  *base64Captcha.DriverString
	DriverChinese *base64Captcha.DriverChinese
	DriverMath    *base64Captcha.DriverMath
	DriverDigit   *base64Captcha.DriverDigit
}

var store = base64Captcha.DefaultMemStore

// base64Captcha create http handler
func SendTextMessage(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"code": "code",
		"msg":  "message",
		"data": "短信已发送，请注意查收。",
	})
}

// base64Captcha create http handler
func GenerateCaptchaHandler(c *gin.Context) {
	driverString := base64Captcha.DriverString{
		Height:          40,
		Width:           100,
		NoiseCount:      0,
		ShowLineOptions: 2 | 4,
		Length:          4,
		Source:          "1234567890qwertyuioplkjhgfdsazxcvbnm",
		BgColor:         &color.RGBA{R: 3, G: 102, B: 214, A: 125},
		Fonts:           []string{"wqy-microhei.ttc"},
	}

	driver := driverString.ConvertFonts()

	id, content, answer := driver.GenerateIdQuestionAnswer()
	item, err := driver.DrawCaptcha(content)

	if err != nil {
		log.Println(err)
	}

	captchaBase64 := item.EncodeB64string()
	code := e.SUCCESS
	log.Println(content)
	log.Println(answer)

	store.Set(id, answer)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": captchaBase64,
	})
}

// base64Captcha verify http handler
func captchaVerify(c *gin.Context) bool {

	store.Get("iji", false)

	//parse request json body
	//decoder := json.NewDecoder(r.Body)
	//var param configJsonBody
	//err := decoder.Decode(&param)
	//if err != nil {
	//	log.Println(err)
	//}
	//defer r.Body.Close()
	////verify the captcha
	//body := map[string]interface{}{"code": 0, "msg": "failed"}
	//if store.Verify(param.Id, param.VerifyValue, true) {
	//	body = map[string]interface{}{"code": 1, "msg": "ok"}
	//}
	////set json response
	//w.Header().Set("Content-Type", "application/json; charset=utf-8")
	//
	//json.NewEncoder(w).Encode(body)
	return true
}

// UserRegisterCaptcha godoc
// @Summary UserRegisterCaptcha
// @Description UserRegisterCaptcha
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} handler.RespBody{data=schema.GetUserResp}
// @Router /answer/api/v1/user/register/captcha [get]
func UserRegisterCaptcha(ctx *gin.Context) {
	// resp, err := action.UserRegisterCaptcha()

}
