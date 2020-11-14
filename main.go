package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"questionnaire_api/api/ioc"
	_ "questionnaire_api/docs"
	"questionnaire_api/middleware/cors"
	"questionnaire_api/orm"
	"questionnaire_api/router"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @host 127.0.0.1:8089
// @BasePath /api/v1
func main() {
	r := gin.Default()
	InitConfig()
	r.Use(cors.Cors())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.InitRouter(r)
	ioc.InitIoc()
	if isSuccess := orm.GetInstance().InitDataPool(); !isSuccess {
		panic("数据库初始化错误")
	}

	if err := r.Run(fmt.Sprintf(":%d", viper.Get("port"))); err != nil {
		panic(err)
	}
}

//初始化配置文件
func InitConfig() {
	viper.SetConfigName("debug_config") // 指定配置文件名称（不需要带后缀）
	viper.SetConfigType("yaml")         // 指定配置文件类型
	viper.AddConfigPath("./config/")    // 指定查找配置文件的路径（这里使用相对路径）
	err := viper.ReadInConfig()         // 读取配置信息
	if err != nil {                     // 读取配置信息失败
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

