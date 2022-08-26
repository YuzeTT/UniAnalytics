package main

import (
	"embed"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
	"hsott.cn/UniAnalytics/util"
)

//go:embed html
var html embed.FS

type Config struct {
	Title      string `yaml:"title"`
	Card_title string `yaml:"card_title"`
	Card_text  string `yaml:"card_text"`
}

func init() {
	fmt.Println("===== UniAnalytics v0.2.2 =====")
	log.Println("初始化中...")
	t := time.Now()
	util.InitConfig()
	log.Println("===========================")
	util.InitSql()
	log.Println("初始化完成，耗时", time.Since(t))
	log.Println("===========================")
}

func main() {
	log.Println("启动服务中...")
	t := time.Now()

	//读取配置文件
	conf := &Config{}
	data, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		return
	}
	if err := yaml.Unmarshal([]byte(data), &conf); err != nil {
		return
	}

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	// r.LoadHTMLGlob("html/*")
	template, _ := template.ParseFS(html, "html/*.html")
	r.SetHTMLTemplate(template)

	r.GET("/", func(ctx *gin.Context) {
		j := ctx.Query("j")
		jumpType := ctx.Query("type")
		if j == "" {
			ctx.String(http.StatusOK, "参数错误")
			return
		}
		if jumpType == "" {
			ctx.HTML(http.StatusOK, "jump.html", gin.H{
				"url":        j,
				"title":      conf.Title,
				"card_title": conf.Card_title,
				"card_text":  conf.Card_text,
			})
			return
		}
		util.AddSql(j, ctx.ClientIP())
		ctx.Redirect(http.StatusMovedPermanently, j)
	})

	log.Println("启动服务完成，耗时", time.Since(t))
	log.Println("启动成功: http://127.0.0.1:8080 ")
	log.Println("===========================")
	r.Run(":8080")
}
