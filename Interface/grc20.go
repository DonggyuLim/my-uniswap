package Interface

import (
	"bytes"
	"encoding/gob"

	"github.com/DonggyuLim/grc20/db"
)

type GRC20 interface {
	GetName() string
	GetSymbol() string
	GetTotalSupply() uint64
	BalanceOf(account string) uint64
	GetDecimal() uint8
	Transfer(from, to string, amount uint64) error
	TransferFrom(from, to, sepnder string, amount uint64) error
	Allowance(owner, spender string) uint64
	Approve(owner, spender string, amount uint64) error
	Mint(account string, amount uint64)
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
