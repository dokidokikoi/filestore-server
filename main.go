package main

import (
	"file-store/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/file/upload", handler.UploadHandler)
	http.HandleFunc("/file/download", handler.DownloadHandler)
	http.HandleFunc("/file/upload", handler.UpdateHandler)
	http.HandleFunc("/file/delete", handler.DeleteHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
