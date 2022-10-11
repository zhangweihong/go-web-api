package helper

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"gin-framework/basic/src/common"
	"gin-framework/basic/src/tool"
	"io"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

//高效拼接字符串
var sBuild strings.Builder

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
func GetUploadsFilePath(relPath string) (string, string) {
	curDir, _ := os.Getwd()
	//相对的路径
	newRelPath := tool.Splicing(upLoadDir, GetLocalShortTime(), "/", relPath)
	//绝对路径
	p := tool.Splicing(curDir, "/", newRelPath)
	p = strings.ReplaceAll(p, "\\", "/")
	dir := path.Dir(p)
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		os.MkdirAll(dir, 0777)
	}
	return p, newRelPath
}

//删除文件
func DelUploadsFile(relPath string) {
	curDir, _ := os.Getwd()
	newRelPath := tool.Splicing(upLoadDir, relPath)
	p := tool.Splicing(curDir, "/", newRelPath)
	_, err := os.Stat(p)

	if os.IsExist(err) {
		os.Remove(p)
	} else {
		common.Logger.Warn(err.Error())
	}
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

//获取当前时间戳
func GetTimeUnix() string {
	now := time.Now() //获取当前时间
	return strconv.FormatInt(now.UnixMilli(), 10)
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
		common.Logger.Error(err)
	}
	return hex.EncodeToString(h.Sum(nil))
}

//获取唯一ID
func NewUUID() string {
	u1, _ := uuid.NewUUID()
	return u1.String()
}

//解密base64
func Base64Decode(base string) string {
	data, err := base64.StdEncoding.DecodeString(base)
	if err != nil {
		common.Logger.Error(err)
		return ""
	}
	return string(data)
}

//加密base64
func Base64Encode(str string) string {
	data := base64.StdEncoding.EncodeToString([]byte(str))
	return data
}
