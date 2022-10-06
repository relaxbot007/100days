package main

import (
	td "100days/go_micro/protobuf/service"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

const (
	post = "127.0.0.1:8081"
)

type TradeDataServer struct {
}

func (this *TradeDataServer) GetTradeData(ctx context.Context, in *td.GetTradeDataRequest) (*td.GetTradeDataReply, error) {
	return &td.GetTradeDataReply{Message: "Hello" + in.Name}, nil
}

func main() {
	listener, err1 := net.Listen("tcp", post)
	if err1 != nil {
		fmt.Println("Network Error", err1)
		return
	}
	defer listener.Close()

	srv := grpc.NewServer()
	td.RegisterTradeDataServiceServer(srv, &TradeDataServer{})

	for {
		fmt.Println("Waiting for connect")
		err2 := srv.Serve(listener)
		if err2 != nil {
			fmt.Println("Network Error")
		}
	}
}
