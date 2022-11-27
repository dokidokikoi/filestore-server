package routes

import (
	"file-store/internal/controller"
	"file-store/internal/db/store/data"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	store, _ := data.GetStoreDBFactory()

	storeGroup := r.Group("store")
	{
		fileController := controller.NewFileController(store)
		storeGroup.POST("file", fileController.Upload)
		storeGroup.GET("file", fileController.Download)
		storeGroup.PATCH("file", fileController.Update)
		storeGroup.DELETE("file", fileController.Delete)
		storeGroup.GET("file/list", fileController.List)
	}
}
