// 公共方法
package utils

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/mojocn/base64Captcha"
	"github.com/zeromicro/go-zero/rest/httpx"
	"golang.org/x/crypto/bcrypt"
)

// 验证码存放
var Store = base64Captcha.DefaultMemStore

type Body struct {
	Code      int         `json:"code"`
	Msg       string      `json:"msg"`
	Timestamp int         `json:"timestamp"`
	Data      interface{} `json:"data"`
}

func Response(w http.ResponseWriter, code int, msg string, data interface{}) {
	if data == nil || data == "" {
		data = struct{}{}
	}
	body := Body{
		Code:      code,
		Msg:       msg,
		Timestamp: Timestamp(),
		Data:      data,
	}
	httpx.OkJson(w, body)
}

func SuccessJson(w http.ResponseWriter, msg string, data interface{}) {
	if msg == "" {
		msg = "success"
	}
	Response(w, 0, msg, data)
}

func ErrorJson(w http.ResponseWriter, code int, msg string) {
	if code >= 0 {
		code = -1
	}
	Response(w, code, msg, nil)
}

// 生成Guid字串
func UniqueId() string {
	b := make([]byte, 48)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return Md5(base64.URLEncoding.EncodeToString(b))
}

func Md5(str string) string {
	hash := md5.New()
	hash.Write([]byte(str))
	//占位待%x为整型以十六进制方式显示
	return fmt.Sprintf("%x", hash.Sum(nil))
}

// 密码加密
func PasswordHash(pwd string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), err
}

// 获取时间戳
func Timestamp() int {
	t := time.Now().Unix()
	return int(t)
}

// 获取当前日期时间
func DateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// 获取当前日期
func Date() string {
	return time.Now().Format("2006-01-02")
}

// 时间戳转日期
func UnixToDateTime(timestramp int) string {
	t := time.Unix(int64(timestramp), 0)
	return t.Format("2006-01-02 15:04:05") //通用时间模板定义
}

// 时间戳转日期
func UnixToDate(timestramp int) string {
	t := time.Unix(int64(timestramp), 0)
	return t.Format("2006-01-02") //通用时间模板定义
}

// 日期转时间戳
func DateToUnix(str string) int {
	t, err := time.ParseInLocation("2006-01-02 15:04:05", str, time.Local)
	if err != nil {
		return 0
	}
	return int(t.Unix())
}

// 密码验证
func PasswordVerify(pwd string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
	return err == nil
}

// 获取图片完整地址
func DisplayImg(file_url string, filename string) string {
	if filename == "" {
		return ""
	}
	return file_url + "/image/" + filename
}

// 获取图片验证码
func GetCaptcha() (string, string, string, error) {
	driver := base64Captcha.DefaultDriverDigit
	captcha := base64Captcha.NewCaptcha(driver, Store)
	id, b64s, answer, err := captcha.Generate()
	return id, b64s, answer, err
}
