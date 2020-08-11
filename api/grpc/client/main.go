package main

import (
	"context"
	pb "gin/api/grpc/proto"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	address  = "localhost:50051"
	orderno  = "500"
	username = "500"
	amount   =  500
	status   = "500"
	fileurl  = "500"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewSimpleClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	//r, err := c.Insert(ctx, &pb.InsertOrder{Orderno:orderno,Username:username,Amount:amount,Status:status,Fileurl:fileurl})
	r, err := c.Query(ctx, &pb.Request{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	//log.Printf("Greeting: %s", strconv.Itoa(int(r.GetId())))
	log.Printf("Greeting: %s", r.GetId())
}