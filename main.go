package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/liubaninc/sdk-go/client"
	_ "github.com/liubaninc/sdk-go/docs"
	"github.com/liubaninc/sdk-go/syncer"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/tendermint/tendermint/libs/log"
	"net/http"
	"os"
)

// @title 区块链浏览器
// @version 1.0
// @description 区块链浏览器后端API接口文档

// @contact.name API Support
// @contact.url http://www.swagger.io/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8080
// @BasePath
func main() {
	client := client.MustNew("http://localhost:26657")

	db, err := syncer.NewDB("sqlite", "synced.db")
	if err != nil {
		panic(fmt.Errorf("failed to open database"))
	}

	syncer.AutoMigrate(db)

	logger := log.NewTMLogger(os.Stdout)
	api := syncer.NewApi(logger, db)
	syncer := syncer.NewHander(logger, db, client)

	start := syncer.GetLastBlockHeight() + 1
	go client.OffChainSyncing(context.Background(), start, syncer)

	router := gin.Default()
	router.Use(func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type, Authorization")
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	})
	router.MaxMultipartMemory = 8 << 20 // 8 MiB

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	api.RegisterApi(router)

	if err := router.Run(fmt.Sprintf(":%d", 8080)); err != nil {
		panic(err)
	}

}
