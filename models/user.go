package models

//用户注册信息
type User struct {
	UserId int `json:"user_id"`
	Name string `json:"name"`
	Password string `json:"password"`
	Time string `json:"time"`
}

//用户登入信息
type UserLogin struct {
	UserId int
	PassWord string
}


