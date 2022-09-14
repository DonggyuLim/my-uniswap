package Interface

import (
	"bytes"
	"encoding/gob"

	"github.com/DonggyuLim/grc20/db"
	"github.com/shopspring/decimal"
)

type GRC20 interface {
	GetName() string
	GetSymbol() string
	GetTotalSupply() decimal.Decimal
	BalanceOf(account string) decimal.Decimal
	GetDecimal() uint8
	Transfer(from, to string, amount decimal.Decimal) error
	TransferFrom(from, to, sepnder string, amount decimal.Decimal) error
	Allowance(owner, spender string) decimal.Decimal
	Approve(owner, spender string, amount decimal.Decimal) error
	Mint(account string, amount decimal.Decimal)
}

// GRC20struct -> byte
func StructToByte(data GRC20) []byte {
	var result bytes.Buffer
	enc := gob.NewEncoder(&result)
	err := enc.Encode(data)
	if err != nil {
		panic(err)
	}
	return result.Bytes()
}

// save Token data
func SaveToken(name string, token GRC20) {
	db.Add(name, StructToByte(token))
}
