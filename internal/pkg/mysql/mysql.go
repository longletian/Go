package mysql

import "github.com/jinzhu/gorm"

func InitMysqlData() {
	sqlconstr:="root:9624968Mi@(localhost)/goblog?charset=utf8mb4&parseTime=True&loc=Local"
  	db,err:=gorm.Open("mysql",sqlconstr)
  	if err!=nil {
		panic("failed to connect database")
	}
  	defer  db.Close()
}

func   InitAutoMgirate(...interface{}) {
	gorm.DB.AutoMigrate()
}


func MysqlClose() {
	gorm.DB.Close()
}
