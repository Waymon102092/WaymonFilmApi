package services

import (
	"Waymon_api/pkg/aliyun"
	"Waymon_api/pkg/e"
	"Waymon_api/pkg/res"
	"context"
	"fmt"
	"mime/multipart"
	"path"
	"strconv"
	"sync"
	"time"
)

type UploadService struct {
	Path string `json:"path" form:"path"`
}

func (service *UploadService) Upload(ctx context.Context, files []*multipart.FileHeader) res.Response {
	code := e.Success
	wg := new(sync.WaitGroup)
	wg.Add(len(files))
	for _, file := range files {
		tmp, _ := file.Open()
		ex := path.Ext(file.Filename)
		t := time.Now()
		pathName := fmt.Sprintf("code/%s%s%s%s%s%s%s",
			strconv.Itoa(t.Year()),
			strconv.Itoa(int(t.Month())),
			strconv.Itoa(t.Day()),
			strconv.Itoa(t.Hour()),
			strconv.Itoa(t.Minute()),
			strconv.Itoa(t.Second()),
			strconv.Itoa(int(t.UnixNano())))
		fileName := pathName + ex
		path, err := aliyun.UploadFile(tmp, fileName)
		if err != nil {
			code = e.UploadError
			return res.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Data:   err.Error(),
			}
		}
		service.Path = path
		wg.Done()
	}
	wg.Wait()
	return res.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   service.Path,
	}
}
