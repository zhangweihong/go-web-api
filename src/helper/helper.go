package helper

import (
	"fmt"
	"os"
	"path"
	"strings"
)

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
	var p = fmt.Sprintf("%v/%v%v", curDir, upLoadDir, relPath)
	p = strings.ReplaceAll(p, "\\", "/")
	dir := path.Dir(p)
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		os.MkdirAll(dir, 0777)
	}
	return p
}
