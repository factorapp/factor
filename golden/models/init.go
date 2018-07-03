package models

// factor.RPC is a generated typesafe thing that wraps net/rpc
var client factor.RPCClient

func init() {
	client = factor.NewRPCClient()
}
