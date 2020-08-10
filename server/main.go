package main

import (
	"context"
	"encoding/json"
	"fmt"
	"gin/api/model"
	"gin/api/service"
	simple "gin/proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"strconv"
	"strings"
	"time"
)

const(
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
}

func (s *server) mustEmbedUnimplementedSimpleServer() {
	panic("implement me")
}

func (s *server) Query(ctx context.Context,req *simple.HelloRequest) (*simple.HelloReplay, error){
	service := new(service.Service)
	result, _ :=service.Query()
	buf, _ := json.Marshal(result)
	return &simple.HelloReplay{Message:string(buf)},nil
}

func (s *server) VagueQuery(ctx context.Context,req *simple.HelloRequest) (*simple.HelloReplay, error){
	service := new(service.Service)
	result, _ :=service.VagueQuery(req.Name)
	buf, _ := json.Marshal(result)
	return &simple.HelloReplay{Message:string(buf)},nil
}

func (s *server) SortQuery(ctx context.Context,req *simple.HelloRequest) (*simple.HelloReplay, error){
	service := new(service.Service)
	result, _ :=service.SortQuery(req.Name)
	buf, _ := json.Marshal(result)
	return &simple.HelloReplay{Message:string(buf)},nil
}

func (s *server) Get(ctx context.Context,req *simple.HelloRequest) (*simple.HelloReplay, error){
	id, _ := strconv.Atoi(req.Name)
	service := new(service.Service)
	result, _ :=service.Get(id)
	buf, _ := json.Marshal(result)
	return &simple.HelloReplay{Message:string(buf)},nil
}

func (s *server) Insert(ctx context.Context,req *simple.HelloRequest) (*simple.HelloReplay, error){
	var order model.DemoOrder
	str := strings.Split(req.Name,",")
	order.Orderno = str[0]
	order.Username = str[1]
	order.Amount, _ = strconv.ParseFloat(str[2],64)
	order.Status = str[3]
	order.Fileurl = str[4]
	order.Time = time.Now().Format("2006-01-02 15:04:05")
	service := new(service.Service)
	result, _  := service.Insert(&order)
	return &simple.HelloReplay{Message:string(result)},nil
}

func (s *server) Update(ctx context.Context,req *simple.HelloRequest) (*simple.HelloReplay, error){
	var order model.DemoOrder
	str := strings.Split(req.Name,",")
	order.Id, _ = strconv.Atoi(str[0])
	order.Amount, _ = strconv.ParseFloat(str[1],64)
	order.Status = str[2]
	order.Fileurl = str[3]
	service := new(service.Service)
	result, _  := service.Update(&order)
	buf, _ := json.Marshal(result)
	return &simple.HelloReplay{Message:string(buf)},nil
}

func (s *server) UpLoad(ctx context.Context,req *simple.HelloRequest) (*simple.HelloReplay, error){
	var c *gin.Context
	service := new(service.Service)
	service.UpLoad(c)
	return &simple.HelloReplay{Message:"上传成功"},nil
}

func (s *server) DownLoad(ctx context.Context,req *simple.HelloRequest) (*simple.HelloReplay, error){
    var c *gin.Context
	service := new(service.Service)
	service.DownLoad(c)
	return &simple.HelloReplay{Message:"下载成功"},nil
}

func (s *server) ExcelDownLoad(ctx context.Context,req *simple.HelloRequest) (*simple.HelloReplay, error){

	var c *gin.Context
	service := new(service.Service)
	err := service.ExcelDownLoad(c)
	if err != nil {
		fmt.Println("下载失败")
	}
	return &simple.HelloReplay{Message:"下载成功"},nil
}

func main(){
	lis,err := net.Listen("tcp",port)

	if err != nil {
		log.Fatal("fail to listen")
	}

	s := grpc.NewServer()

	simple.RegisterSimpleServer(s,&server{})

	reflection.Register(s)

	if err:= s.Serve(lis);err != nil{
		log.Fatal("fail to server")
	}
}

