package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type Goods struct {
	Id      int
	Title   string
	Price   float32
	Content string
}

type AddGoodsRes struct {
	Success bool
	Message string
}

func (this Goods) AddGoods(req Goods, res *AddGoodsRes) error {
	fmt.Println(req)
	*res = AddGoodsRes{
		Success: true,
		Message: "Success",
	}
	return nil
}

type GetGoodsReq struct {
	Id int
}
type GetGoodsRes struct {
	Id      int
	Title   string
	Price   float32
	Content string
}

func (this Goods) GetGoods(req GetGoodsReq, res *GetGoodsRes) error {
	fmt.Printf("Start to query id %#v\n", req)
	*res = GetGoodsRes{
		Id:      23,
		Title:   "毛巾被",
		Price:   32.3,
		Content: "服务器返回",
	}
	return nil

}

func main() {
	err1 := rpc.RegisterName("goods", new(Goods))
	if err1 != nil {
		fmt.Println(err1)
	}

	listener, err2 := net.Listen("tcp", ":8082")
	if err2 != nil {
		fmt.Println(err2)
	}

	defer listener.Close()
	for {
		fmt.Println("Accept connection")
		conn, err3 := listener.Accept()
		if err3 != nil {
			fmt.Println(err3)
		}
		rpc.ServeConn(conn)
	}

}
