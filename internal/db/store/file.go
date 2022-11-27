package store

import "file-store/internal/db/model"

type File interface {
	OnFileUploadFinished(filehash, filename, fileaddr string, filesize int64) error
	GetFileMeta(filehash string) (*model.File, error)
	List(page, limit int) ([]model.File, error)
}
