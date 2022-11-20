package meta

import (
	"sync"
)

// 文件元信息
type FileMeta struct {
	Sha256   string
	FileName string
	FileSize int64
	Location string
	UploadAt string
}

var (
	fileMetas map[string]*FileMeta

	lock sync.RWMutex
)

func init() {
	fileMetas = make(map[string]*FileMeta)
}

func UpdateFileMeta(fmeta *FileMeta) {
	lock.Lock()
	defer lock.Unlock()

	fileMetas[fmeta.Sha256] = fmeta
}

func GetFileMeta(sha256 string) *FileMeta {
	lock.RLock()
	defer lock.RUnlock()

	return fileMetas[sha256]
}

func DelFileMeta(sha256 string) {
	lock.Lock()
	defer lock.Unlock()

	delete(fileMetas, sha256)
}
