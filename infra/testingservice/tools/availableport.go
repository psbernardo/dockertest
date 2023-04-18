package tools

import (
	"net"
)

func GetAvailablePort() int {
	l, err := net.Listen("tcp", ":0")

	if err != nil {
		panic(err)
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port
}
