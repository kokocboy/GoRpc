package main

import (
	"fmt"
	"log"
	"rpcClient/rpcClient"
)
func main() {
	client, err := rpcClient.RpcServiceStart("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	var option int
	fmt.Scan(&option)
	if option==1 {
		var ans int
		err = client.Mul(3,2, &ans)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(ans)
	}else{
		var ans string
		err = client.AddPre("test", &ans)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(ans)
	}
}