package database

import (
	"Beego_Web/models"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)


//用户注册，保存用户名和密码，以及注册时间在数据库中
func Inssert_user(user models.User) (int,int64,error) {

	md5 := md5.New()

	md5.Write([]byte(user.Password))

	bytes := md5.Sum(nil)

	user.Password = hex.EncodeToString(bytes)
	fmt.Println(user.Password)

	//用于保存当前时间的[]byte类型
	var time_silce []byte
	time_silce = []byte(time.Now().String())

	//得到2020-09-26 14:37:36类型的格式的时间
	time := string(time_silce[:19])
	fmt.Println("时间",time)

	//用户注册时间，将得到的2020-09-26 14:37:36这种时间格式的时间赋值给user.Time
	user.Time = time

	result,err := Db.Exec("insert into user(name,password,time) value (?,?,?)",user.Name,user.Password,user.Time)
	if err != nil {
		return 0,-1,err
	}

	affectedNum,err := result.RowsAffected()

	rows,err := Db.Query("select userId from user where time=?",time)
	var userId int
	for rows.Next() {
		err = rows.Scan(&userId)
		if err != nil {
			//做出相应的处理
		}
	}
	return userId,-affectedNum,nil
}


//用户登入
func User_login(userId,userPwd string) (string,bool){


	var userName string		//保存按输入ID查出来的用户名
	var pwd string			//保存按输入ID查出来的密码hashMD5
	//fmt.Println("用户输入的UserId:",userId,"用户输入的密码:",userPwd)
	//fmt.Println()

	//1.拿到用户输入的用户名(数字Id)和密码,因为数据库中用户名没有设置唯一，而数字id是唯一的
	//userId = userId		//哈哈哈，，多余的可不写


	//2.按用户输入的Id查询用户是否存在，存在返回true，否则返回false，同时查出相应的密码
	rows,err := Db.Query("select username, password from user where id=?",userId)
	for rows.Next() {
		err = rows.Scan(&userName,&pwd)
		if err != nil {
			//做出相应的处理
		}
	}
	//fmt.Printf("查出的用户名:%s 秘密:%s",userName,pwd)
	if userName == "" {
		return "用户不存在,请注册后重试",false
	}

	//3.拿到用户名和密码后将用户输入的密码转成hashMD5
	md5 := md5.New()

	md5.Write([]byte(userPwd))

	bytes := md5.Sum(nil)

	userPwd = hex.EncodeToString(bytes)

	//fmt.Println("用户输入密码的hash值",userPwd)
	//fmt.Println("实际用户的密码hash值",pwd)
	//fmt.Println()
	//4.对比查出来的密码转成MD5的值与用户输入的密码转成MD5是否相等，
	//如果相等则返回登入成功和true，否则返回登入失败和false
	if userPwd != pwd {
		return "用户名或密码错误，请重试!",false
	}

	return "登入成功,"  + userName + "欢迎您!",true
}

//用户的注销，
func Delete_user(userId,userPwd string) (string,bool) {


	return "登入成功，欢迎您!",true
	//按相应的方法注销用户
}

//用户信息的修改
func Update_userInfo() {

}

//aaaaaaaa
func Update_userInfoaaaa() {

}