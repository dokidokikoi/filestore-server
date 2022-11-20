package tools

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"hash"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

type Sha256Stream struct {
	hash hash.Hash
}

func (s *Sha256Stream) Update(data []byte) {
	if s.hash == nil {
		s.hash = sha256.New()
	}

	s.hash.Write(data)
}

func (s *Sha256Stream) Sum() string {
	return hex.EncodeToString(s.hash.Sum([]byte("")))
}

func Sha256(data []byte) string {
	hash := sha256.New()
	hash.Write(data)
	return hex.EncodeToString(hash.Sum([]byte("")))
}

func FileSha256(file *os.File) string {
	hash := sha256.New()
	io.Copy(hash, file)

	return hex.EncodeToString(hash.Sum(nil))
}

func MD5(data []byte) string {
	_md5 := md5.New()
	_md5.Write(data)

	return hex.EncodeToString(_md5.Sum([]byte("")))
}

func FileMD5(file *os.File) string {
	_md5 := md5.New()
	io.Copy(_md5, file)

	return hex.EncodeToString(_md5.Sum(nil))
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}

func GetFileSize(filename string) int64 {
	var result int64
	filepath.Walk(filename, func(path string, info fs.FileInfo, err error) error {
		result = info.Size()
		return nil
	})

	return result
}
