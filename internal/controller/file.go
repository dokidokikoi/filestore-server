package controller

import (
	"file-store/internal/db/store"
	"file-store/internal/meta"
	"file-store/internal/service"
	"file-store/pkg/core"
	"file-store/pkg/errors"
	"file-store/pkg/log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type FileController struct {
	service service.Service
}

func (f FileController) Upload(ctx *gin.Context) {
	// 接收文件流幷存放到本地目錄
	head, err := ctx.FormFile("file")
	if err != nil {
		log.Log().Error("接收文件失敗", zap.Any("err", err))
		core.WriteResp(ctx, nil, errors.Failed(http.StatusBadRequest, "上傳失败"))
		return
	}

	fileMeta, err := f.service.File().SaveFile(head)
	core.WriteResp(ctx, fileMeta, errors.Success("上傳成功"))
}

func (f FileController) Download(ctx *gin.Context) {
	fsha256 := ctx.Query("filehash")

	fm, data, err := f.service.File().GetFile(fsha256)
	if err != nil {
		core.WriteResp(ctx, nil, errors.Failed(http.StatusServiceUnavailable, err.Error()))
	}

	w := ctx.Writer
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment;filename=\""+fm.FileName+"\"")
	w.Write(data)
}

func (f FileController) Update(ctx *gin.Context) {
	type req struct {
		FileHash string `json:"filehash"`
		FileName string `json:"filename"`
	}

	r := &req{}

	if err := ctx.ShouldBindJSON(r); err != nil {
		core.WriteResp(ctx, nil, errors.Failed(http.StatusBadRequest, "数据格式有误"))
		return
	}

	curFileMeta, _ := f.service.File().UpdateMeta(r.FileHash, r.FileName)

	core.WriteResp(ctx, curFileMeta, errors.Success(""))
}

func (f FileController) Delete(ctx *gin.Context) {
	type req struct {
		FileHash string `json:"filehash"`
		FileName string `json:"filename"`
	}

	r := &req{}
	if err := ctx.ShouldBindJSON(r); err != nil {
		core.WriteResp(ctx, nil, errors.Failed(http.StatusBadRequest, "数据格式有误"))
		return
	}

	meta.DelFileMeta(r.FileHash)
	core.WriteResp(ctx, nil, errors.Success(""))
}

func (f FileController) List(ctx *gin.Context) {
	type req struct {
		Page  int `json:"page"`
		Limit int `json:"limit"`
	}

	r := &req{}
	if err := ctx.ShouldBindJSON(r); err != nil {
		core.WriteResp(ctx, nil, errors.Failed(http.StatusBadRequest, "数据格式有误"))
		return
	}

	metas, _ := f.service.File().List(r.Page, r.Limit)
	core.WriteResp(ctx, metas, errors.Success(""))
}

func NewFileController(store store.Factory) *FileController {
	return &FileController{
		service: service.NewSerivce(store),
	}
}
