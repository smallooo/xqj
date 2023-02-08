package util

import "xqj/pkg/setting"

// Setup Initialize the utils
func Setup() {
	jwtSecret = []byte(setting.AppSetting.JwtSecret)
}
