package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
	"hsott.cn/UniShortLink/util"
)

var data_str = `title: 安全中心
card_title: 出站提示
card_text: 您即将离开本站，请注意您的帐号和财产安全。`

type Config struct {
	Title      string `yaml:"title"`
	Card_title string `yaml:"card_title"`
	Card_text  string `yaml:"card_text"`
}

func FileExist(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func init() {
	fmt.Println("=== UniAnalytics v0.1.0 ===")
	log.Println("初始化中...")
	t := time.Now()

	// 检查配置文件
	isFile := FileExist("./config.yaml")
	if !isFile {
		log.Println("未找到数据库文件，开始自动创建")
		file, err := os.Create("./config.yaml")
		if err != nil {
			log.Fatalf("配置文件有误，示例配置文件[data.toml]：\n%v\n更多细节请学习Toml语法", data_str)
			log.Fatalln(err)
		} else {
			file.WriteString(data_str)
			log.Println("未检测到配置文件，已自动创建")
		}
	}

	util.InitSql()
	log.Println("初始化完成，耗时", time.Since(t))
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
	fmt.Printf("--- t:\n%v\n\n", conf.Title)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.LoadHTMLGlob("html/*")

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
	log.Println("启动成功: 127.0.0.1:8080 ")

	r.Run(":8080")
}
