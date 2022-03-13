package tool

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitMysql() error {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/dou_ban?charset=utf8")
	if err != nil {
		panic("the config of mysql is wrong!")
		return err
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic("can not link the mysql!")
		return err
	}
	DB = db
	fmt.Println("linking to mysql...")
	return nil
}

//提交失误，补充注释作为修改，再次提交
