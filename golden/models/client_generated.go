package models

import (
	"net/rpc"
)

type Client struct {
	
	TodoClient
	
}

func NewClient(port int) (*Client, error) {
	client, err := rpc.DialHTTP("tcp", ":1234")
	if err != nil {
		return nil, err
	}
	return &Client{
		
		TodoClient{RPC: client},
		
	}, nil
}
