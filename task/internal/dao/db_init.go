package dao

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"strings"
	"time"
)

var DB *gorm.DB

func InitDB() {
	host := viper.GetString("mysql.host")
	port := viper.GetString("mysql.port")
	database := viper.GetString("mysql.database")
	username := viper.GetString("mysql.username")
	password := viper.GetString("mysql.password")
	charset := viper.GetString("mysql.charset")
	dsn := strings.Join([]string{username, ":", password, "@tcp(", host, ":", port, ")/", database, "?charset=" + charset + "&parseTime=true"}, "")
	err := Database(dsn)
	if err != nil {
		panic(err)
	}
}

func Database(dsn string) error {
	var ormLogger logger.Interface
	if gin.Mode() == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default
	}
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                      dsn,
		DefaultStringSize:        256,
		DisableDatetimePrecision: true, //禁用datetime的精度
		DontSupportRenameIndex:   true, //不支持重命名索引
		DontSupportRenameColumn:  true,
	}), &gorm.Config{
		Logger: ormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		}, //单数的命名策略，生成的表没有s后缀
	})
	if err != nil {
		return err
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxIdleTime(time.Second * 30)
	DB = db
	err = migration()
	if err != nil {
		return err
	}
	return nil
}
