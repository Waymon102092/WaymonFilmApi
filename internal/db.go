package internal

import (
	"Waymon_api/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"strings"
	"time"
)

var (
	DB        *gorm.DB
	WriteConn string
)

func InitDB() {
	host := viper.GetString("mysql.host")
	port := viper.GetString("mysql.port")
	database := viper.GetString("mysql.database")
	username := viper.GetString("mysql.username")
	password := viper.GetString("mysql.password")
	charset := viper.GetString("mysql.charset")
	WriteConn = strings.Join([]string{username, ":", password, "@tcp(", host, ":", port, ")/", database, "?charset=", charset, "&parseTime=true"}, "")

	var gormLogger logger.Interface
	if gin.Mode() == "debug" {
		gormLogger = logger.Default.LogMode(logger.Info)
	} else {
		gormLogger = logger.Default
	}
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       WriteConn,
		DefaultStringSize:         256,   //数据库字符串最大长度
		DisableDatetimePrecision:  true,  //禁用datetime精度
		DontSupportRenameIndex:    true,  //重命名索引
		DontSupportRenameColumn:   true,  //重命名列 兼容低版本
		SkipInitializeWithVersion: false, //根据版本自动配置
	}), &gorm.Config{
		Logger: gormLogger, //输出日志
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //table不加s
		},
	})
	if err != nil {
		fmt.Println("db Open err:", err)
	}
	sqlDb, _ := db.DB()
	sqlDb.SetMaxIdleConns(20)  //连接池
	sqlDb.SetMaxOpenConns(100) //连接数
	sqlDb.SetConnMaxLifetime(time.Second * 30)
	DB = db
	//迁移
	//migration()
	fmt.Println("mysql 初始化成功")
}

func migration() {
	err := DB.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(
		&model.Amount{},
		&model.Admin{},
		&model.Amount{},
		&model.Apply{},
		&model.Banner{},
		&model.Brand{},
		&model.Config{},
		&model.File{},
		&model.Custom{},
		&model.Help{},
		&model.Media{},
		&model.Member{},
		&model.Menu{},
		&model.Message{},
		&model.Order{},
		&model.Poster{},
		&model.Promote{},
		&model.Report{},
		&model.ReportCategory{},
		&model.ReportImg{},
		&model.Role{},
		&model.RoleMenu{},
		&model.Withdraw{},
		&model.City{},
		&model.Hot{},
	)
	if err != nil {
		fmt.Println("migration err:", err)
	}
}
