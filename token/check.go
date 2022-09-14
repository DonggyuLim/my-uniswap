package token

import (
	"errors"

	"github.com/shopspring/decimal"
)

// check
func (t *Token) checkBalance(owner string, amount uint64) error {
	balance := t.BalanceOf(owner)
	if balance.Cmp(decimal.New(int64(amount), 10)) == -1 {
		return errors.New("amount is biger than owner balance")
	}
	return nil
}

func (t *Token) checkspendAllowance(owner, spender string, amount uint64) error {
	Allowance := t.allowance(owner, spender)
	if Allowance.Cmp(decimal.New(int64(amount), 10)) == -1 {
		return errors.New("amount is biger than allowance")
	}
	return nil
}
