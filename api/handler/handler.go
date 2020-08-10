package handler

import (
	"gin/api/model"
	"gin/api/service"
	_ "gin/api/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

//数据查询 http://localhost:8000/user/query
func Query(c *gin.Context) {
	service := new(service.Service)
	result, err :=service.Query()
	if err!= nil || len(result) == 0{
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "抱歉未找到相关信息",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data":   result,
	})
}

//模糊查询 http://localhost:8000/user/vague/query
func VagueQuery(c *gin.Context) {
	username := c.Request.FormValue("username")
	service := new(service.Service)
	result, err :=service.VagueQuery(username)

	if err!= nil || len(result) == 0{
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "抱歉未找到相关信息",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data":   result,
	})
}

//排序查询 http://localhost:8000/user/sort/query
func SortQuery(c *gin.Context) {
	sort := c.Request.FormValue("sort")
	service := new(service.Service)
	result, err :=service.SortQuery(sort)

	if err!= nil || len(result) == 0{
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "抱歉未找到相关信息",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data":   result,
	})
}

//单个查询  http://localhost:8000/user/get:id
func Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Request.FormValue("id"))
	service := new(service.Service)
	result, err := service.Get(id)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "抱歉未找到相关信息",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data":   result,
	})
}

//添加数据 http://localhost:8000/user/add
func Insert(c *gin.Context) {
	var order model.DemoOrder
	c.BindJSON(&order)
	order.Time = time.Now().Format("2006-01-02 15:04:05")
	service := new(service.Service)
	id, err  := service.Insert(&order)

	if err != nil || id == 0{
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "添加失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":  1,
		"message": "添加成功",
	})
}

//修改数据  http://localhost:8000/user/update
func Update(c *gin.Context) {
	var order model.DemoOrder
	c.BindJSON(&order)
	service := new(service.Service)
	result, err := service.Update(&order)

	if err != nil || result.Id == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "修改失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":  1,
		"message": "修改成功",
	})
}

//上传更新  http://localhost:8000/upload
func UpLoad(c *gin.Context) {
	service := new(service.Service)
	err :=service.UpLoad(c)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "上传更新失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":  1,
		"message": "上传更新成功",
	})
}

//下载文件 http://localhost:8000/download
func DownLoad(c *gin.Context) {
	service := new(service.Service)
	service.DownLoad(c)
}

//下载文件 http://localhost:8000/excel/download
func ExcelDownLoad(c *gin.Context) {
	service := new(service.Service)
	err := service.ExcelDownLoad(c)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "下载失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":  1,
		"message": "下载成功",
	})
}



