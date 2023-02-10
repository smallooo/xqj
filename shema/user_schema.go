package shema

type ActionRecordReq struct {
	// action
	Action string `validate:"required,oneof=login e_mail find_pass" form:"action"`
	IP     string `json:"-"`
}

type ActionRecordResp struct {
	CaptchaID  string `json:"captcha_id"`
	CaptchaImg string `json:"captcha_img"`
	Verify     bool   `json:"verify"`
}
