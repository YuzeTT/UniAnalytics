package util

import (
	"log"
	"os"
)

var data_str = `title: 安全中心
card_title: 出站提示
card_text: 您即将离开本站，请注意您的帐号和财产安全。`

func InitConfig() {
	isFile := FileExist("./config.yaml")
	if !isFile {
		log.Println("未找到配置文件，已自动创建")
		file, err := os.Create("./config.yaml")
		if err != nil {
			log.Fatalf("配置文件有误，示例配置文件[config.yaml]：\n%v\n更多细节请学习Toml语法", data_str)
			log.Fatalln(err)
		} else {
			file.WriteString(data_str)
			log.Println("配置文件已创建，请关闭程序编辑 config.yaml 后重新启动")
		}
	}
}
