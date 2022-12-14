package client

import (
	"context"
	"log"
	"sync"
	"time"

	pb "github.com/DonggyuLim/grc20/protos/RPC"
	"github.com/shopspring/decimal"
	"google.golang.org/grpc"
)

const (
	serviceHost = "localhost:9001"
	zero        = "0"
)

var (
	once sync.Once
)

type client struct {
	client pb.RPCClient
	conn   *grpc.ClientConn
}

var cli *client

func StartClient(wg *sync.WaitGroup) {
	once.Do(func() {
		conn, err := grpc.Dial(serviceHost, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		c := pb.NewRPCClient(conn)
		cli = &client{
			conn:   conn,
			client: c,
		}

	})

	defer wg.Done()

}

func GetClient() *client {
	return cli
}

func (c *client) CloseConn() {
	c.conn.Close()
}

func (c *client) GetBalance(tokenName, account string) (decimal.Decimal, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := c.client.GetBalance(ctx, &pb.GetBalanceRequest{
		TokenName: tokenName,
		Account:   account,
	})
	if err != nil {
		return decimal.Zero, err
	}
	balance := decimal.RequireFromString(res.GetBalance())
	return balance, nil
}

func (c *client) TokenInfo(tokenName string) (*pb.TokenInfoResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := c.client.GetTokenInfo(ctx, &pb.TokenInfoRequest{
		TokenName: tokenName,
	})
	if err != nil {
		return &pb.TokenInfoResponse{}, err
	}
	return res, nil
}

func (c *client) Allowance(tokenName, owner, spender string) (decimal.Decimal, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := c.client.GetAllowance(ctx, &pb.AllowanceRequest{
		TokenName: tokenName,
		Owner:     owner,
		Spender:   spender,
	})
	if err != nil {
		return decimal.Zero, err
	}
	allowance := decimal.RequireFromString(res.GetAllowance())
	return allowance, nil
}

func (c *client) Approve(tokenName, owner, spender string, amount string) (*pb.ApproveResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := c.client.Approve(ctx, &pb.ApproveRequest{
		TokenName: tokenName,
		Owner:     owner,
		Spender:   spender,
		Amount:    amount,
	})
	if err != nil {
		return &pb.ApproveResponse{}, err
	}
	return res, nil
}

func (c *client) TransferFrom(tokenName, owner, spender, to string, amount string) (*pb.TransferFromResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := c.client.TransferFrom(ctx, &pb.TransferFromRequest{
		TokenName: tokenName,
		Spender:   spender,
		Owner:     owner,
		To:        to,
		Amount:    amount,
	})
	if err != nil {
		return &pb.TransferFromResponse{}, err
	}
	return res, nil
}

func (c *client) Transfer(tokenName, from, to string, amount string) (*pb.TransferResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := c.client.Transfer(ctx, &pb.TransferRequest{
		TokenName: tokenName,
		From:      from,
		To:        to,
		Amount:    amount,
	})

	if err != nil {
		return &pb.TransferResponse{}, err
	}
	return res, nil
}
