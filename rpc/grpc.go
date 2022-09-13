package rpc

import (
	"log"
	"net"
	"sync"

	rpc "github.com/DonggyuLim/erc20/protos/RPC"
	"google.golang.org/grpc"
)

const (
	RPCPort = "9000"
)

type RPCServer struct {
	rpc.RPCServer
}

// func (r *RPCServer) Transfer(ctx context.Context, req *rpc.TransferRequest) (*rpc.TransferResponse, error) {
// 	tokenName := req.TokenName

// }

func Grpc(wg *sync.WaitGroup) {
	lis, err := net.Listen("tcp", ":"+RPCPort)
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}
	grpcServer := grpc.NewServer()
	rpc.RegisterRPCServer(grpcServer, &RPCServer{})

	log.Printf("start gRPC server on %s port", RPCPort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to server:%s", err)
	}
	defer wg.Done()
}
