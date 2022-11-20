package handler

import (
	"file-store/core"
	"file-store/errors"
	"file-store/log"
	"file-store/meta"
	"file-store/tools"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"go.uber.org/zap"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	zap.L().Sugar().Infof("sssssss")
	if r.Method == "GET" {

	} else if r.Method == "POST" {
		// 接收文件流幷存放到本地目錄
		file, head, err := r.FormFile("file")
		if err != nil {
			log.Log().Error("接收文件失敗", zap.Any("err", err))
			return
		}
		defer file.Close()

		fileMeta := &meta.FileMeta{
			FileName: head.Filename,
			Location: "/tmp/" + head.Filename,
			UploadAt: time.Now().Format("2006-01-02 15:04:05"),
		}

		newFile, err := os.Create(fileMeta.Location)
		if err != nil {
			log.Log().Error("創建文件失敗", zap.Any("err", err))
			return
		}
		defer newFile.Close()

		fileMeta.FileSize, err = io.Copy(newFile, file)
		if err != nil {
			log.Log().Error("保存文件失敗", zap.Any("err", err))
			return
		}

		newFile.Seek(0, 0)
		fileMeta.Sha256 = tools.FileSha256(newFile)

		meta.UpdateFileMeta(fileMeta)

		core.WriteResp(w, fileMeta, errors.Success("上傳成功"))
	}
}

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fsha256 := r.Form.Get("filehash")
	fm := meta.GetFileMeta(fsha256)

	file, err := os.Open(fm.Location)
	if err != nil {
		core.WriteResp(w, nil, errors.Failed(http.StatusInternalServerError, "服務器錯誤"))
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		core.WriteResp(w, nil, errors.Failed(http.StatusInternalServerError, "服務器錯誤"))
		return
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment;filename=\""+fm.FileName+"\"")
	w.Write(data)
}

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	opType := r.Form.Get("op")
	fileSha256 := r.Form.Get("filehash")
	newFileName := r.Form.Get("filename")

	if opType == "0" {
		core.WriteResp(w, nil, errors.Failed(http.StatusForbidden, ""))
	}

	if r.Method != "POST" {
		core.WriteResp(w, nil, errors.Failed(http.StatusMethodNotAllowed, ""))
	}

	curFileMeta := meta.GetFileMeta(fileSha256)
	curFileMeta.FileName = newFileName
	meta.UpdateFileMeta(curFileMeta)

	core.WriteResp(w, curFileMeta, errors.Success(""))
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	filesha256 := r.Form.Get("filehash")

	meta.DelFileMeta(filesha256)
	core.WriteResp(w, nil, errors.Success(""))
}
