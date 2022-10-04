package main

import (
	"fmt"
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

type GetGoodsReq struct {
	Id int
}
type GetGoodsRes struct {
	Id      int
	Title   string
	Price   float32
	Content string
}

func main() {
	conn, err1 := rpc.Dial("tcp", "127.0.0.1:8082")
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	defer conn.Close()

	var reply AddGoodsRes
	err2 := conn.Call("goods.AddGoods", Goods{
		Id:      10,
		Title:   "毛巾",
		Price:   10,
		Content: "不掉毛",
	}, &reply)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println("Add Goods Reply:", reply)

	var getGoodsReply GetGoodsRes
	conn.Call("goods.GetGoods", GetGoodsReq{
		Id: 12,
	}, &getGoodsReply)
	fmt.Println("Get Goods Reply:", getGoodsReply)
}
