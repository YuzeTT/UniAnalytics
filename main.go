package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"hsott.cn/UniShortLink/util"
)

func init() {
	fmt.Println("=== UniAnalytics v0.1.0 ===")
	log.Println("初始化中...")
	t := time.Now()
	util.InitSql()
	log.Println("初始化完成，耗时", time.Since(t))
}

func main() {
	log.Println("启动服务中...")
	t := time.Now()

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		j := ctx.Query("j")
		if j == "" {
			ctx.String(http.StatusOK, "参数错误")
			return
		}
		ctx.Redirect(http.StatusMovedPermanently, j)
	})

	log.Println("启动服务完成，耗时", time.Since(t))
	log.Println("启动成功: 127.0.0.1:8080 ")
	r.Run(":8080")
}
