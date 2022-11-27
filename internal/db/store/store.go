package store

import "fmt"

type Factory interface {
	File() File
}

var storePointer Factory

func SetFactory(store Factory) {
	storePointer = store
}

func GetFactory() (Factory, error) {
	if storePointer == nil {
		return nil, fmt.Errorf("数据层未初始化")
	}
	return storePointer, nil
}
