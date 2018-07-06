package models

import (
	"net/rpc"
	"net"
	"fmt"
	"net/http"
)

func StartRPCServer(port int) error {
	rpc.HandleHTTP()
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	return http.Serve(listener, nil)
}
