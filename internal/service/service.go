package service

import (
	"file-store/internal/db/store"
)

type Service interface {
	File() FileSrv
}

type service struct {
	store store.Factory
}

func (s service) File() FileSrv {
	return newFile(s.store)
}

func NewSerivce(store store.Factory) Service {
	return &service{store}
}
