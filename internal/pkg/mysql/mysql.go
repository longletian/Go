package mysql

import (
	"demoProject/config"
	"github.com/jinzhu/gorm"
)

func InitMysqlData()  *gorm.DB {
	var   db  *gorm.DB
	config,err:=config.GetConfig()
	if err ==nil{
		sqlcon:=config.MySqlConfig.Username+
			":"+config.MySqlConfig.Password+
			"@tcp("+config.MySqlConfig.Hostname+
			")/"+config.MySqlConfig.Database
		db,err=gorm.Open("mysql",sqlcon)
		if err!=nil {
			//异常抛出和记录异常
			panic("failed to connect database")
		}
		defer  db.Close()
	}
	return db
}

func  InitMigrate(data ...interface{})  {
	///初始化数据库
    db:=InitMysqlData()
	db.AutoMigrate(data)
}

func MysqlClose() {
	db:=InitMysqlData()
	db.Close()
}
