package main

import (
	"sync"

	"github.com/DonggyuLim/erc20/db"
	"github.com/DonggyuLim/erc20/rest"
	"github.com/DonggyuLim/erc20/rpc"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(2)
	db.NewDB()
	defer db.Close()
	go rest.Rest(wg)
	go rpc.Grpc(wg)
	wg.Wait()
}
