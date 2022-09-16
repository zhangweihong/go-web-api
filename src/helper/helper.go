package helper

import (
	"os"
	"path"
	"strings"
)

func Try(fun func(), catch func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			catch(err)
		}
	}()
	fun()
}

func GetUploadsFilePath(relPath string) string {
	curDir, _ := os.Getwd()
	var p = curDir + "/upload/" + relPath
	p = strings.ReplaceAll(p, "\\", "/")
	dir := path.Dir(p)
	os.MkdirAll(dir, 0777)
	return p
}

var log_file_name = "project-logs.log"
var log_file_path = "/logs"

func GetLogFilePath() string {
	curDir, _ := os.Getwd()
	var p = curDir + "/" + log_file_path
	p = path.Join(p, log_file_name)
	p = strings.ReplaceAll(p, "\\", "/")
	dir := path.Dir(p)
	os.MkdirAll(dir, 0777)
	return p
}
