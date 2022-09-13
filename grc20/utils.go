package grc20

import (
	"bytes"
	"encoding/gob"
	"errors"

	"github.com/DonggyuLim/erc20/db"
)

func NewToken(name, symbol string, decimal uint8) *Token {
	bm := make(map[string]uint64)
	am := make(map[string]uint64)

	t := &Token{
		Name:        name,
		Symbol:      symbol,
		Decimal:     decimal,
		TotalSupply: 0,
		Balance:     bm,
		Allowances:  am,
	}
	return t
}

// byte -> Token
func ByteToToken(data []byte) *Token {
	var token *Token
	encoder := gob.NewDecoder(bytes.NewBuffer(data))
	err := encoder.Decode(&token)
	if err != nil {
		panic(err)
	}
	return token
}

// tokenName -> db -> token
func GetToken(tokenName string) (*Token, error) {
	value, ok := db.Get(tokenName)
	if !ok {
		return &Token{}, errors.New("invalid token")

	}
	t := ByteToToken(value)

	return t, nil
}
