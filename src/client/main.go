package main

import (
	"fmt"
	"log"
	"rpcClient/rpcClient"
)

var address = "172.25.144.119"
//var address = "localhost"
func main() {
	client, err := rpcClient.RpcServiceStart("tcp", address+":1002")
	defer client.Close()
	if err != nil {
		log.Fatal("dialing:", err)
	}
		var ans int
		err = client.Mul(3,2, &ans)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(ans)
		var str string
		err = client.AddPre("test", &str)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(str)
}