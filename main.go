package main

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	_ "javbus-api-go/docs" // 这里导入你的 Swagger 文档生成的 package
)

// @title Your API Title
// @version 1.0
// @description Your API Description
// @termsOfService https://example.com/terms
// @host example.com
// @basePath /api/v1
// @schemes http https
func main() {
	r := gin.New()

	// Your routes go here

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}
