package models

type Auth struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Login struct {
	Uid       int    `gorm:"primary_key" json:"id"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Passcode  string `json:"passcode"`
	Passwd    string `json:"passwd"`
	LoginIp   string `json:"login_ip"`
	LoginTime string `json:"login_time"`
}

func CheckAuth(username, password string) bool {
	var auth Auth
	db.Select("id").Where(Auth{Username: username, Password: password}).First(&auth)
	if auth.ID > 0 {
		return true
	}

	return false
}

func LoginRecord(login Login) bool {
	var auth Auth
	var record = Login{Uid: login.Uid, Email: login.Email, Username: login.Username, Passcode: login.Passcode, Passwd: login.Passwd, LoginIp: login.LoginIp, LoginTime: login.LoginTime}

	db.Save(&record)
	if auth.ID > 0 {
		return true
	}

	return false
}
