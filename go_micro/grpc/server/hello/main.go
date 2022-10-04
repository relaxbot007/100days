package main

import (
	"fmt"
	"net"
	"net/rpc"
)

// 定义结构体
type Hello struct {
}

// 函数名大写 public 函数
// req res 不能是channel func类型 不能序列话
func (this Hello) SayHello(req string, res *string) error {
	*res = "Hello" + req
	return nil
}

func main() {
	//register rpc service
	err1 := rpc.RegisterName("hello", new(Hello))
	if err1 != nil {
		fmt.Println(err1)
	}

	//监听端口
	listener, err2 := net.Listen("tcp", ":8080")
	if err2 != nil {
		fmt.Println(err2)
	}
	defer listener.Close()

	for {
		fmt.Println("Wait to connect...")
		conn, err3 := listener.Accept()
		if err3 != nil {
			fmt.Println(err3)
		}
		rpc.ServeConn(conn)
	}
}
