package router

import (
	. "gin/api/handler"
	"github.com/gin-gonic/gin"
)


func InitRouter() *gin.Engine {
	router := gin.Default()
	// 设置文件上传大小限制，默认是32m
	router.MaxMultipartMemory = 64 << 20  // 64 MiB

	router.GET("/user/query", Query)
	router.GET("/user/vague/query", VagueQuery)
	router.GET("/user/sort/query", SortQuery)
	router.GET("/user/get", Get)
	router.GET("/user/add", Insert)
	router.GET("/user/update", Update)
	router.POST("/upload", UpLoad)
	router.POST("/download", DownLoad)
	router.POST("/excel/download", ExcelDownLoad)
	return router
}
