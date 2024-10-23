/*
*

	@author: kiki
	@since: 2024/10/22
	@desc: //TODO

*
*/
package main

import (
	"fmt"
	"freelink/DB"
	"freelink/middlewares"
	"freelink/routes"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var (
	db                  *gorm.DB
	err                 error
	dbusername          string
	dbpassword          string
	dbhost              string
	dbport              string
	dbname              string
	dbcharset           string
	dbsecret            string
	tokenexpirationtime string
)

func init() {

	func() {
		for key, value := range map[string]string{"DB_HOST": "192.168.31.99", "DB_PORT": "3306", "DB_USERNAME": "root", "DB_PASSWORD": "Qaqaqa00.0", "DB_CHARSET": "utf8mb4", "DB_NAME": "freelink", "API_SECRET": "1qaz@wsx", "TOKEN_EXPIRATION_TIME": "14400"} {
			if err := os.Setenv(key, value); err != nil {
				fmt.Println(err.Error())
				return
			}
		}

		dbhost = os.Getenv("DB_HOST")
		dbport = os.Getenv("DB_PORT")
		dbname = os.Getenv("DB_NAME")
		dbsecret = os.Getenv("API_SECRET")
		dbcharset = os.Getenv("DB_CHARSET")
		dbusername = os.Getenv("DB_USERNAME")
		dbpassword = os.Getenv("DB_PASSWORD")
		tokenexpirationtime = os.Getenv("TOKEN_EXPIRATION_TIME")

		fmt.Println(dbhost, dbport, dbname, dbusername, dbpassword, dbcharset)

		if dbusername == "" || dbpassword == "" || dbhost == "" || dbport == "" || dbname == "" || dbcharset == "" || tokenexpirationtime == "" || dbsecret == "" {
			panic("DB_USERNAME, DB_PASSWORD, DB_NAME, DB_CHARSET, DB_HOST, DB_PORT, TOKEN_EXPIRATION_TIME, API_SECRET")
		}
	}()

	// connect db
	if db, err = gorm.Open(mysql.Open(fmt.Sprintf(`%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local`, dbusername, dbpassword, dbhost, dbport, dbname, dbcharset)), &gorm.Config{
		Logger: logger.New(log.New(os.Stdout, "\r\n", log.Flags()), logger.Config{SlowThreshold: time.Second, LogLevel: logger.Info, Colorful: true}),
	}); err != nil {
		panic(err)
	}

	// 设置连接池
	if dbobj, err := db.DB(); err != nil {
		panic(err)
	} else {
		dbobj.SetMaxIdleConns(10)
		dbobj.SetMaxOpenConns(100)
		dbobj.SetConnMaxLifetime(time.Hour)
	}

	// 设置自动迁移模式
	if err = db.AutoMigrate(DB.AutoMigrate()...); err != nil {
		panic(err)
	}

}

func main() {
	ginservices := gin.Default()

	// 跨域保护
	ginservices.Use(func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept, Authorization, X-CSRF-Token")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(200)
		} else {
			ctx.Next()
		}
	})

	ginservices.Use(middlewares.JwtAuthMiddleware())
	routes.Login(ginservices, db)

	fmt.Println("Gin service release port 80")
	if err = ginservices.Run(":80"); err != nil {
		panic(err)
	}
}
