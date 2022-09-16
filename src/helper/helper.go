package helper

import (
	"os"
	"path"
	"strings"
)

var sBuild strings.Builder

//高效拼接字符串
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
