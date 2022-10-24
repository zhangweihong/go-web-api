package helper

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"errors"
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
func GetUploadsFilePath(relPath string) (abPath string, rPath string) {
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

//pkcs7Padding 填充
func pkcs7Padding(data []byte, blockSize int) []byte {
	//判断缺少几位长度。最少1，最多 blockSize
	padding := blockSize - len(data)%blockSize
	//补足位数。把切片[]byte{byte(padding)}复制padding个
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

//pkcs7UnPadding 填充的反向操作
func pkcs7UnPadding(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, errors.New("加密字符串错误！")
	}
	//获取填充的个数
	unPadding := int(data[length-1])
	return data[:(length - unPadding)], nil
}

//AES 加密 iv长度16位
func AES_CBC_PK7Encrypt(strData string, strKey string, strIV string) (base64 string, err error) {
	//创建加密实例
	key := []byte(strKey)
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	iv := []byte(strIV)
	data := []byte(strData)
	//判断加密快的大小
	blockSize := block.BlockSize()
	//填充
	encryptBytes := pkcs7Padding(data, blockSize)
	//初始化加密数据接收切片
	crypted := make([]byte, len(encryptBytes))
	//使用cbc加密模式
	blockMode := cipher.NewCBCEncrypter(block, iv)
	//执行加密
	blockMode.CryptBlocks(crypted, encryptBytes)

	return Base64Encode(string(crypted)), nil
}

//AesDecrypt 解密 iv长度16位
func AES_CBC_PK7Decrypt(base64Str string, strKey string, strIV string) (deStr string, err error) {
	key := []byte(strKey)
	//创建实例
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	iv := []byte(strIV)
	data := []byte(Base64Decode(base64Str))

	//使用cbc
	blockMode := cipher.NewCBCDecrypter(block, iv)
	//初始化解密数据接收切片
	crypted := make([]byte, len(data))
	//执行解密
	blockMode.CryptBlocks(crypted, data)
	//去除填充
	crypted, err = pkcs7UnPadding(crypted)
	if err != nil {
		return "", err
	}
	return string(crypted), nil
}
