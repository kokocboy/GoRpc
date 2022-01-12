package rpcServer

import (
	"log"
	"net"
	"net/rpc"
)

const CboyServiceName = "CboyService"
func RpcCboyServiceStart(){
	rpc.RegisterName(CboyServiceName, new(CboyService))
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}else{
			log.Println("sucessful")
		}
		go rpc.ServeConn(conn)
	}
}

type CboyService struct {}

func (c *CboyService) AddPre(request string, reply *string) error {
	*reply = "cboy:" + request
	return nil
}
func (c *CboyService)Add(nums []int,ans *int) error  {
	*ans=nums[0]+nums[1]
	return nil
}
func (c *CboyService)Mul(nums []int,ans *int) error  {
	*ans=nums[0]*nums[1]
	return nil
}