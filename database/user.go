package database

import (
	"Beego_Web/models"
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

//用户注册，插入用户名和密码，以及注册时间
func Inssert_user(user models.User) (int64,error) {

	md5 := md5.New()

	md5.Write([]byte(user.Password))

	bytes := md5.Sum(nil)

	user.Password = hex.EncodeToString(bytes)
	fmt.Println(user.Password)
	//time := []byte(time.Now().String())
	//
	//user.Time = time
	//result,err := Db.Exec("insert into user(name,password,time) value (?,?,?)",user.Name,user.Password,time.Now())
	//if err != nil {
	//	return -1,err
	//}
	//
	//affectedNum,err := result.RowsAffected()

	return -2,nil
}

//用户的注销，
func Delete_user() {

}

//用户信息的修改
func Update_userInfo() {

}