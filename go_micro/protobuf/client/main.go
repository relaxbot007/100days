package main

import (
	td "100days/go_micro/protobuf/service"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	post = "127.0.0.1:8081"
)

func main() {
	conn, err1 := grpc.Dial(post, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err1 != nil {
		fmt.Println("Connect to the server error")
	}
	defer conn.Close()

	c := td.NewTradeDataServiceClient(conn)
	r1, err2 := c.GetTradeData(context.Background(), &td.GetTradeDataRequest{Name: "Linux"})
	if err2 != nil {
		fmt.Println("could not get trade data server", err2)
	}
	fmt.Println("From thee server ", r1.Message)
}
