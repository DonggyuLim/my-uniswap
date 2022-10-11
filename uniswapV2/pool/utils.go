package pool

import (
	"errors"

	"github.com/DonggyuLim/uniswap/client"
	"github.com/shopspring/decimal"
)

func grc20TwoBalanceCheck(address string, tokenA, tokenB Token) error {
	err := grc20CheckBalance(tokenA.Name, address, tokenA.GetBalance())
	if err != nil {
		return err
	}
	err = grc20CheckBalance(tokenB.Name, address, tokenB.GetBalance())
	if err != nil {
		return err
	}

	return nil
}

func grc20CheckBalance(tokenName, address string, amount decimal.Decimal) error {
	cli := client.GetClient()
	balance, err := cli.GetBalance(tokenName, address)

	if err != nil {
		return err
	}
	if balance.Cmp(amount) == -1 {
		return errors.New("balance isn't enough")
	}
	return nil
}
