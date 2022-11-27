package postgres

import (
	"file-store/internal/db/model"
	"fmt"
)

func (s Store) OnFileUploadFinished(filehash, filename, fileaddr string, filesize int64) error {
	meta := &model.File{
		FileSha256: filehash,
		FileName:   filename,
		FileAddr:   fileaddr,
		FileSize:   filesize,
		Status:     model.Active,
	}

	if err := s.db.Create(meta).Error; err != nil {
		fmt.Println("failed to create meta data, err:", err)
		return err
	}

	return nil
}

func (s Store) GetFileMeta(filehash string) (*model.File, error) {
	meta := &model.File{FileSha256: filehash}

	if err := s.db.First(meta).Error; err != nil {
		fmt.Println("failed to create meta data, err:", err)
		return nil, err
	}

	return meta, nil
}

func (s Store) List(page, limit int) ([]model.File, error) {
	files := &[]model.File{}
	if err := s.db.Limit(limit).Offset((page - 1) * limit).Find(files).Error; err != nil {
		return nil, err
	}

	return *files, nil
}

func (s Store) Count() (int64, error) {
	var total int64
	s.db.Model(&model.File{}).Count(&total)
	return total, nil
}
