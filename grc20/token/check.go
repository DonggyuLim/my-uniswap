package token

import (
	"errors"

	"github.com/shopspring/decimal"
)

// check
func (t *Token) checkBalance(owner string, amount decimal.Decimal) error {
	balance := t.BalanceOf(owner)
	if balance.Cmp(amount) == -1 {
		return errors.New("amount is biger than owner balance")
	}
	return nil
}

func (t *Token) checkspendAllowance(owner, spender string, amount decimal.Decimal) error {
	Allowance := t.allowance(owner, spender)
	if Allowance.Cmp(amount) == -1 {
		return errors.New("amount is biger than allowance")
	}
	return nil
}
