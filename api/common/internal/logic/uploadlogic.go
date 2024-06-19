package logic

import (
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"

	"go-micro/api/common/internal/svc"

	"github.com/disintegration/imaging"
	"github.com/zeromicro/go-zero/core/logx"
)

type UploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadLogic {
	return &UploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadLogic) Upload(req *http.Request) (res map[string]interface{}, stat int, err error) {
	stat = 1
	dir := req.PostFormValue("dir")
	thumb_w, _ := strconv.Atoi(req.PostFormValue("thumb_w"))
	thumb_h, _ := strconv.Atoi(req.PostFormValue("thumb_h"))
	file_url := l.svcCtx.Config.FileUrl
	file, handler, err := req.FormFile("filename")
	if err != nil {
		return nil, -2, errors.New("请选择文件")
	}
	defer file.Close()

	//文件大小校验
	upload_max_size := l.svcCtx.Config.UploadMaxSize
	if int(handler.Size) > upload_max_size*1024*1024 {
		return nil, -3, errors.New("您上传的文件过大,最大值为" + strconv.Itoa(upload_max_size) + "MB")
	}

	//文件后缀过滤
	ext := path.Ext(handler.Filename) //输出.jpg
	allow_ext_map := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".gif":  true,
	}
	if _, ok := allow_ext_map[ext]; !ok {
		return nil, -4, errors.New("文件格式不正确")
	}

	//创建目录
	upload_dir := l.svcCtx.Config.UploadDir
	dir_num := time.Now().Format("20060102") //输出 20240404/
	err = os.MkdirAll(upload_dir+"/"+dir+"/"+dir_num, os.FileMode(0775))
	if err != nil {
		return nil, -5, errors.New("创建目录失败")
	}

	//构造文件名
	source := rand.NewSource(time.Now().UnixNano()) //这里用系统时间毫秒值当种子值
	rnd := rand.New(source)
	rand_num := fmt.Sprintf("%d", rnd.Intn(9999)+1000) //获取1000-9999随机数
	hash_name := md5.Sum([]byte(time.Now().Format("2006_01_02_15_04_05_") + rand_num))
	file_name := fmt.Sprintf("%x", hash_name) + ext //文件名 例 cf386af3f37962ad3769054f68d7a049.jpg
	path := upload_dir + "/" + dir + "/" + dir_num + "/" + file_name
	filelink := file_url + "/" + dir + "/" + dir_num + "/" + file_name
	//保存文件
	tempFile, err := os.Create(path)
	if err != nil {
		return nil, -6, errors.New("创建图片失败")
	}
	defer tempFile.Close() //关闭文件数据
	io.Copy(tempFile, file)

	//生成缩略图
	if thumb_w > 0 || thumb_h > 0 {
		src, err := imaging.Open(path)
		if err != nil {
			return nil, -7, errors.New("开启缩略图失败")
		}
		dsc := imaging.Resize(src, thumb_w, thumb_h, imaging.Lanczos)

		//构造缩略图文件名
		source := rand.NewSource(time.Now().UnixNano()) //这里用系统时间毫秒值当种子值
		r := rand.New(source)
		rand_num := fmt.Sprintf("%d", r.Intn(9999)+1000) //获取1000-9999随机数
		hash_name := md5.Sum([]byte(time.Now().Format("2006_01_02_15_04_05_") + rand_num))
		file_name = fmt.Sprintf("%x", hash_name) + ext //文件名 例 cf386af3f37962ad3769054f68d7a049.jpg
		path = upload_dir + "/" + dir + "/" + dir_num + "/" + file_name
		filelink = file_url + "/" + dir + "/" + dir_num + "/" + file_name

		err = imaging.Save(dsc, path)
		if err != nil {
			return nil, -8, errors.New("生成缩略图失败")
		}
	}

	//组装数据
	resp := make(map[string]interface{}) //创建1个空集合
	resp["realname"] = handler.Filename
	resp["filename"] = dir_num + "/" + file_name
	resp["filelink"] = filelink
	return resp, stat, nil
}
