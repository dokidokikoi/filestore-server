package service

import (
	"errors"
	"file-store/internal/db/store"
	"file-store/internal/meta"
	"file-store/pkg/log"
	"file-store/tools"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"time"

	"go.uber.org/zap"
)

type FileSrv interface {
	SaveFile(head *multipart.FileHeader) (*meta.FileMeta, error)
	GetFile(fsha256 string) (*meta.FileMeta, []byte, error)
	UpdateMeta(fileSha256, newFileName string) (*meta.FileMeta, error)
	List(page, limit int) (*[]meta.FileMeta, error)

	UploadFileMetaDB(fileMeta *meta.FileMeta) error
}

type fileSrv struct {
	store store.Factory
}

func (f fileSrv) SaveFile(head *multipart.FileHeader) (*meta.FileMeta, error) {
	fileMeta := &meta.FileMeta{
		FileName: head.Filename,
		Location: "/tmp/" + head.Filename,
		UploadAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	newFile, err := os.Create(fileMeta.Location)
	if err != nil {
		log.Log().Error("創建文件失敗", zap.Any("err", err))
		return nil, err
	}
	defer newFile.Close()

	file, _ := head.Open()
	defer file.Close()
	fileMeta.FileSize, err = io.Copy(newFile, file)
	if err != nil {
		log.Log().Error("保存文件失敗", zap.Any("err", err))
		return nil, err
	}

	newFile.Seek(0, 0)
	fileMeta.Sha256 = tools.FileSha256(newFile)

	f.UploadFileMetaDB(fileMeta)

	return fileMeta, nil
}

func (f fileSrv) GetFile(fsha256 string) (*meta.FileMeta, []byte, error) {
	var fm *meta.FileMeta
	fm = meta.GetFileMeta(fsha256)
	if fm == nil {
		fileMeta, err := f.store.File().GetFileMeta(fsha256)
		if err != nil {
			return nil, nil, err
		}

		fm = &meta.FileMeta{
			FileName: fileMeta.FileName,
			FileSize: fileMeta.FileSize,
			Location: fileMeta.FileAddr,
			Sha256:   fileMeta.FileSha256,
		}

		meta.UpdateFileMeta(fm)
	}

	file, err := os.Open(fm.Location)
	if err != nil {
		log.Log().Error("打开文件失敗", zap.Any("err", err))
		return nil, nil, errors.New("服務器錯誤")
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Log().Error("读文件失敗", zap.Any("err", err))
		return nil, nil, errors.New("服務器錯誤")
	}

	return fm, data, nil
}

func (f fileSrv) UpdateMeta(fileSha256, newFileName string) (*meta.FileMeta, error) {
	curFileMeta := meta.GetFileMeta(fileSha256)
	curFileMeta.FileName = newFileName
	meta.UpdateFileMeta(curFileMeta)

	return curFileMeta, nil
}

func (f fileSrv) List(page, limit int) (*[]meta.FileMeta, error) {
	files, _ := f.store.File().List(page, limit)
	metas := []meta.FileMeta{}
	for _, f := range files {
		metas = append(metas, meta.FileMeta{
			FileName: f.FileName,
			FileSize: f.FileSize,
			Location: f.FileAddr,
			Sha256:   f.FileSha256,
		})
	}

	return &metas, nil
}

func (f fileSrv) UploadFileMetaDB(fileMeta *meta.FileMeta) error {
	return f.store.File().OnFileUploadFinished(fileMeta.Sha256, fileMeta.FileName, fileMeta.Location, fileMeta.FileSize)
}

func newFile(store store.Factory) FileSrv {
	return &fileSrv{store}
}
