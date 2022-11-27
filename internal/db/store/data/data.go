package data

import (
	"file-store/internal/config"
	"file-store/internal/db/store"
	postgres "file-store/internal/db/store/data/pg"
	"sync"
)

type dataCenter struct {
	pg *postgres.Store
}

func (d *dataCenter) File() store.File {
	return newFile(d)
}

var (
	datacFactory store.Factory
	once         sync.Once
)

func GetStoreDBFactory() (store.Factory, error) {
	once.Do(func() {
		pg, err := postgres.GetPGFactory(config.PgConfig)
		if err != nil {
			panic(err)
		}

		datacFactory = &dataCenter{pg}
	})

	return datacFactory, nil
}

func SetStoreDBFactory() {
	factory, err := GetStoreDBFactory()
	if err != nil {
		return
	}

	store.SetFactory(factory)
}
