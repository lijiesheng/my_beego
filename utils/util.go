package utils

import (
	"fmt"
	"github.com/astaxie/beego"
	"os"
	"path/filepath"
	"strings"
)

const (
	Version = "1.0"
	StoreLocal string = "local"
	StoreOss string = "oss"
)

var (
	BasePath, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	StoreType = beego.AppConfig.String("store_type") // 存储类型
)

// 操作图片
// 如果用的是 oos 存储，这 style 是avatar、cover 可选项
func ShowImg(img string, style ...string) (url string) {
	if strings.HasPrefix(img, "https://") || strings.HasPrefix(img, "http://") {
		return img
	}
	img = "/" + strings.TrimLeft(img, "./")  //将字符串最前面的空格修整掉
	switch StoreType {
	case StoreOss:
		s := ""
		if len(style) > 0 && strings.TrimSpace(style[0]) != "" {
			s = "/" + style[0]
		}
		url = strings.TrimRight(beego.AppConfig.String("oss::Domain"), "/ ") + img + s
	case StoreLocal:
		url = img
	}
	fmt.Println("img===>",img)
	fmt.Println("url===>", url)
	return url
}