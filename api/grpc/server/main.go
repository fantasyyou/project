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

func (s *server) Insert(ctx context.Context,req *simple.InsertOrder) (*simple.InsertReplay, error){
	var order model.DemoOrder
	order.Orderno  = req.Orderno
	order.Username = req.Username
	order.Amount   = req.Amount
	order.Status   = req.Status
	order.Fileurl  = req.Fileurl
	order.Time     = time.Now().Format("2006-01-02 15:04:05")
	service := new(service.Service)
	result, _  := service.Insert(&order)
	return &simple.InsertReplay{Id:result},nil
}


func (s *server) Query(ctx context.Context,req *simple.Request) (*simple.OrderList, error){
	service := new(service.Service)
	result, _  := service.Query()
	var order []int64
	order = make([]int64,5,10)
	for i:=0; i < len(result); i++ {
		order[i] = result[i].Id
	}
	//var order []*simple.Order
	//var str []string
	return &simple.OrderList{Id:order},nil
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

