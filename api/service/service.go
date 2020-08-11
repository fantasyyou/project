package service

import (
	"encoding/json"
	"fmt"
	"gin/api/dao"
	"gin/api/model"
	"github.com/gin-gonic/gin"
	"github.com/tealeg/xlsx"
	_ "io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Service struct {
	service *Service
	dao     *dao.Dao
}

//数据查询 http://localhost:8000/user/query
func(service *Service) Query() (orders []model.DemoOrder, err error){
	orders,err = service.dao.Query()
	return
}

//模糊查询 http://localhost:8000/user/vague/query
func(service *Service) VagueQuery(username string) (orders []model.DemoOrder, err error){
	orders, err = service.dao.VagueQuery(username)
	if err != nil {
		return
	}
	return
}

//排序查询 http://localhost:8000/user/sort/query
func(service *Service) SortQuery(sort string) (orders []model.DemoOrder, err error){
	orders, err = service.dao.SortQuery(sort)
	if err != nil {
		return
	}
	return
}

//单个查询  http://localhost:8000/user/get:id
func(service *Service) Get(id int) (order model.DemoOrder, err error){
	order,err = service.dao.Get(id)
	if err != nil {
		return
	}
	return
}

//添加数据 http://localhost:8000/user/add
func(service *Service) Insert(order *model.DemoOrder) (id int64,err error){
	id, err = service.dao.Insert(order)
	if err != nil {
		return
	}
	return
}

//修改数据  http://localhost:8000/user/update
func(service *Service) Update(order *model.DemoOrder) (updateOrder model.DemoOrder, err error){
	updateOrder, _ = service.dao.Update(order)
	if err != nil {
		return
	}
	return
}

//上传更新  http://localhost:8000/upload
func(service *Service) UpLoad(c *gin.Context) (err error){
	// 获取上传文件，返回的是multipart.FileHeader对象，代表一个文件，里面包含了文件名之类的详细信息
	// file是表单字段名字
	file, _ := c.FormFile("file")
	// 打印上传的文件名
	log.Println(file.Filename)
	// 将上传的文件，保存到路径文件中
	c.SaveUploadedFile(file, "Ceshi.json")
	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))

	files, err := os.Open("Ceshi.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer files.Close()

	fileinfo, err := files.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}

	filesize := fileinfo.Size()
	buffer := make([]byte, filesize)

	bytesread, err := files.Read(buffer)
	if err != nil {
		fmt.Println(err)
		return
	}
	result := string(buffer)
	fmt.Println("bytes read: ", bytesread)

	//反序列化
	var list model.DemoOrderList
	err = json.Unmarshal([]byte(result), &list)
	if err != nil {
		fmt.Println("Can't decode json message", err)
	}

	for i := 0; i < len(list.Matches); i++ {
		service.dao.Update(&list.Matches[i])
	}
	return
}

//下载文件 http://localhost:8000/download
func(service *Service) DownLoad(c *gin.Context) {
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", "Ceshi.txt"))
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.File("Ceshi.txt")
}

//下载文件 http://localhost:8000/excel/download
func(service *Service) ExcelDownLoad(c *gin.Context) (err error){
	result, err := service.dao.Query()
	if err != nil {
		fmt.Println("查询结果为空!")
	}
	file := xlsx.NewFile()
	sheet, _ := file.AddSheet("Sheet1")
	row := sheet.AddRow()
	cell := row.AddCell()
	cell.Value = "Id"
	cell = row.AddCell()
	cell.Value = "Order_No"
	cell = row.AddCell()
	cell.Value = "User_Name"
	cell = row.AddCell()
	cell.Value = "Amount"
	cell = row.AddCell()
	cell.Value = "Status"
	cell = row.AddCell()
	cell.Value = "File_Url"
	cell = row.AddCell()
	cell.Value = "Time"
	for i:=0 ; i <len(result); i++ {
		row := sheet.AddRow()
		cell := row.AddCell()
		cell.Value = strconv.FormatInt(result[i].Id,10)
		cell = row.AddCell()
		cell.Value = result[i].Orderno
		cell = row.AddCell()
		cell.Value = result[i].Username
		cell = row.AddCell()
		cell.Value = strconv.FormatFloat(float64(result[i].Amount), 'E', -1, 32)
		cell = row.AddCell()
		cell.Value = result[i].Status
		cell = row.AddCell()
		cell.Value = result[i].Fileurl
		cell = row.AddCell()
		cell.Value = result[i].Time
	}
	err = file.Save("file.xlsx")
	if err != nil {
		fmt.Println("保存文件失败!")
	}
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", "file.xlsx"))
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.File("file.xlsx")
	return
}






