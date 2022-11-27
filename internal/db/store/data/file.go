package data

import (
	"file-store/internal/db/model"
	postgres "file-store/internal/db/store/data/pg"
)

type File struct {
	pg *postgres.Store
}

func (f File) OnFileUploadFinished(filehash, filename, fileaddr string, filesize int64) error {
	return f.pg.OnFileUploadFinished(filehash, filename, fileaddr, filesize)
}

func (f File) GetFileMeta(filehash string) (*model.File, error) {
	return f.pg.GetFileMeta(filehash)
}

func (f File) List(page, limit int) ([]model.File, error) {
	return f.pg.List(page, limit)
}

func newFile(store *dataCenter) *File {
	return &File{store.pg}
}
