package rpc

import (
	"context"
	"errors"
	"log"
	"net"
	"sync"

	rpc "github.com/DonggyuLim/erc20/protos/RPC"

	u "github.com/DonggyuLim/erc20/utils"
	"google.golang.org/grpc"
)

const (
	RPCPort = "9000"
)

type RPCServer struct {
	rpc.RPCServer
}

func (r *RPCServer) Transfer(ctx context.Context, req *rpc.TransferRequest) (*rpc.TransferResponse, error) {
	tokenName := req.TokenName
	to := req.To
	from := req.From
	amount := req.Amount
	t, err := u.GetToken(tokenName)
	if err != nil {
		return &rpc.TransferResponse{
			Success:     false,
			ToBalance:   0,
			FromBalance: 0,
		}, errors.New("invalid token")
	}
	t.Transfer(from, to, amount)
	u.SaveToken(tokenName, t)
	return &rpc.TransferResponse{
		Success:     true,
		ToBalance:   t.BalanceOf(to),
		FromBalance: t.BalanceOf(from),
	}, nil
}

func (r *RPCServer) Approve(ctx context.Context, req *rpc.ApproveRequest) (*rpc.ApproveResponse, error) {
	tokenName := req.TokenName
	owner := req.Owner
	spender := req.Spender
	amount := req.Amount
	t, err := u.GetToken(tokenName)
	if err != nil {
		return &rpc.ApproveResponse{
			Success: false,
			Balance: 0,
		}, errors.New("invalid token")
	}
	t.Approve(owner, spender, amount)

	u.SaveToken(tokenName, t)
	return &rpc.ApproveResponse{
		Success: true,
		Balance: t.Allowance(owner, spender),
	}, nil
}

func (r *RPCServer) GetBalance(ctx context.Context, req *rpc.GetBalanceRequest) (*rpc.GetBalanceResponse, error) {

	t, err := u.GetToken(req.GetTokenName())
	if err != nil {
		return &rpc.GetBalanceResponse{
			Success: false,
			Balance: 0,
		}, errors.New("invalid token")
	}
	balance := t.BalanceOf(req.GetAccount())
	return &rpc.GetBalanceResponse{
		Success: true,
		Balance: balance,
	}, nil
}

func (r *RPCServer) GetTokenInfo(ctx context.Context, req *rpc.TokenInfoRequest) (*rpc.TokenInfoResponse, error) {
	t, err := u.GetToken(req.GetTokenName())
	if err != nil {
		return &rpc.TokenInfoResponse{
			Success:     false,
			TokenName:   "",
			Symbol:      "",
			Decimal:     0,
			TotalSupply: 0,
		}, errors.New("invalid token")
	}
	return &rpc.TokenInfoResponse{
		Success:     true,
		TokenName:   t.GetName(),
		Symbol:      t.GetSymbol(),
		Decimal:     uint32(t.GetDecimal()),
		TotalSupply: t.GetTotalSupply(),
	}, nil
}

func (r *RPCServer) GetAllowance(ctx context.Context, req *rpc.AllowanceRequest) (*rpc.AllowanceResponse, error) {
	t, err := u.GetToken(req.GetTokenName())
	if err != nil {
		return &rpc.AllowanceResponse{
			Success: false,
			Balance: 0,
		}, errors.New("invalid token")
	}
	allowance := t.Allowance(req.GetOwner(), req.GetSpender())
	return &rpc.AllowanceResponse{
		Success: true,
		Balance: allowance,
	}, nil
}

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
