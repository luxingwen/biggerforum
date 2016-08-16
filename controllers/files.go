package controllers

import (
	"io"
	"os"
	"strconv"
	"time"
)

type FileController struct {
	BaseController
}

func (this *FileController) Post() {
	req := this.Ctx.Request
	f, h, err := req.FormFile("editormd-image-file")
	if err != nil {
		failer(this)
	}
	upload := "upload/" + strconv.Itoa(time.Now().Year()) + "/" + time.Now().Month().String() + "/" + strconv.Itoa(time.Now().Day())
	err = os.MkdirAll(upload, 0777)
	if err != nil {
		failer(this)
	}
	fs, err := os.Create(upload + "/" + h.Filename)
	io.Copy(fs, f)
	this.Data["json"] = map[string]interface{}{"success": 1, "message": "上传成功", "url": "http://192.168.1.50:9999/" + upload + "/" + h.Filename}
	this.ServeJson()
}

func failer(this *FileController) {
	this.Data["json"] = map[string]interface{}{"success": 0, "message": "上传失败"}
	this.ServeJson()
}
