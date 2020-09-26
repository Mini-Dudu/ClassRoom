package database

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

//打开并连接数据库
func OpenDb() (string,error) {

	str := ""

	config := beego.AppConfig
	dbDriver := config.String("db_Driver")
	//fmt.Println("驱动",dbDriver)
	dbUser := config.String("db_User")
	//fmt.Println(dbUser)
	dbPassword := config.String("db_Password")
	dbIp := config.String("db_Ip")
	dbName := config.String("db_Name")

	connUrl := dbUser + ":" + dbPassword + "@tcp(" + dbIp + ")/" + dbName + "?charset=utf8"
	//fmt.Println(connUrl)

	//Time := []byte(time.Now().String())
	//fmt.Println(Time)
	////user.Time = time
	//
	//fmt.Printf("时间类型%T,%s", []byte(time.Now().String()), []byte(time.Now().String()))
	db, err := sql.Open(dbDriver,connUrl)
	if err != nil {
		str = "数据库连接出错，请重试！"
		fmt.Println(err)
		return str,err
	}

	Db = db
	if Db != nil {
		str = "已连接到Mysql数据库"
		return str,nil
	}else {
		str = "打开数据库失败"
		return str,nil
	}

	//var user models.User
	//Inssert_user(user)
	//
	//return str,nil
}

//关闭数据库
func CloseDb() error {
	//如果Db 不等于空，说明开启了数据库
	if Db != nil {
		err := Db.Close()
		return err
	}
	return nil
}