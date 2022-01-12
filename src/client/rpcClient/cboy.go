package rpcClient

import "net/rpc"

const RpcServiceName = "CboyService"
type CboyServiceClient struct {
	*rpc.Client
}
func RpcServiceStart(network, address string) (*CboyServiceClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &CboyServiceClient{Client: c}, nil
}
func (c *CboyServiceClient) AddPre(request string, reply *string) error {
	return c.Client.Call(RpcServiceName+".AddPre", request, reply)
}
func (c *CboyServiceClient)  Add(num1,num2 int,ans *int)error{
	return  c.Client.Call(RpcServiceName+".Add",[]int{num1,num2},ans)
}
func (c *CboyServiceClient)  Mul(num1,num2 int,ans *int)error{
	return  c.Client.Call(RpcServiceName+".Mul",[]int{num1,num2},ans)
}