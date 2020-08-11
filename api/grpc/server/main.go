package main

import (
	"context"
	simple "gin/api/grpc/proto"
	"gin/api/model"
	"gin/api/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"strconv"
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

func (s *server) Insert(ctx context.Context,req *simple.InsertRequest) (*simple.InsertReplay, error){
	var order model.DemoOrder
	order.Orderno = req.Orderno
	order.Username = req.Username
	order.Amount, _ = strconv.ParseFloat(req.Amount,64)
	order.Status = req.Status
	order.Fileurl = req.Fileurl
	order.Time = time.Now().Format("2006-01-02 15:04:05")
	service := new(service.Service)
	result, _  := service.Insert(&order)
	//return &simple.InsertReplay{Id:strconv.Itoa(result)},nil
	return &simple.InsertReplay{Id:strconv.Itoa(result)},nil
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

