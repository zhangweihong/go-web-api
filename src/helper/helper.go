package helper

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"gin-framework/basic/src/middleware"
	"io"
	"os"
	"path"
	"strings"
	"time"
)

//高效拼接字符串
var sBuild strings.Builder

func Splicing(strs ...string) string {
	sBuild.Reset()
	len := len(strs)
	for i := 0; i < len; i++ {
		sBuild.WriteString(strs[i])
	}
	return sBuild.String()
}

//捕获异常
func Try(fun func(), catch func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			catch(err)
		}
	}()
	fun()
}

var upLoadDir = "upload/"

//获取上传的文件路径
func GetUploadsFilePath(relPath string) string {
	curDir, _ := os.Getwd()
	p := Splicing(curDir, "/", upLoadDir, relPath)
	// var p = Splicing("%v/%v%v", curDir, upLoadDir, relPath)
	p = strings.ReplaceAll(p, "\\", "/")
	dir := path.Dir(p)
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		os.MkdirAll(dir, 0777)
	}
	return p
}

//获取当前时间 2019-12-12 12:33:19
func GetLocalTime() string {
	now := time.Now()      //获取当前时间
	year := now.Year()     //年
	month := now.Month()   //月
	day := now.Day()       //日
	hour := now.Hour()     //小时
	minute := now.Minute() //分钟
	second := now.Second() //秒
	return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", year, month, day, hour, minute, second)
}

//获取当前简单时间 2019-12-12
func GetLocalShortTime() string {
	now := time.Now()    //获取当前时间
	year := now.Year()   //年
	month := now.Month() //月
	day := now.Day()     //日
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

//获取MD5
func Md5(value string) string {
	if value == "" {
		return ""
	}
	data := md5.Sum([]byte(value))
	return hex.EncodeToString(data[:])
}

//获取文件Md5
func Md5File(file *os.File) string {
	h := md5.New()
	_, err := io.Copy(h, file)
	if err != nil {
		middleware.Logger.Error(err)
	}
	return hex.EncodeToString(h.Sum(nil))
}
