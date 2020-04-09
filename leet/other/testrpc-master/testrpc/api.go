package testrpc

import (
	"errors"
	"net"
	"routine/leet"
	"sync"
)

func NewServer() *main.Server {
	return &main.Server{
		ServiceMap:  make(map[string]map[string]*main.Service),
		serviceLock: sync.Mutex{}}
}

func NewClient(conn net.Conn) *main.Client {
	return &main.Client{
		conn: conn}
}

func Dial(network, address string) (*main.Client, error) {
	if network != "tcp" {
		return nil, errors.New("Unsupported protocol")
	}

	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}

	return NewClient(conn), nil
}
