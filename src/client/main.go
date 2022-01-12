package main

import (
	"fmt"
	"log"
	"rpcClient/rpcClient"
)
func main() {
	client, err := rpcClient.RpcServiceStart("tcp", "47.109.21.195:1002")
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