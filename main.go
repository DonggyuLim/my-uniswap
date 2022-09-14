package main

import (
	"sync"

	"github.com/DonggyuLim/grc20/db"
	"github.com/DonggyuLim/grc20/rest"
	"github.com/DonggyuLim/grc20/rpc"
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
