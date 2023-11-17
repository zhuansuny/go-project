package main

import (
	"fmt"
	"net"
)

var (
	localIPArray []string
)

func init() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(fmt.Sprintf("get local ip failed, %v", err))
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				localIPArray = append(localIPArray, ipnet.IP.String())
			}
		}
	}
	localIPArray = append(localIPArray, "127.0.0.1")
	fmt.Println(localIPArray)
}
