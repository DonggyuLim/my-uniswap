package grc20

import (
	"bytes"
	"encoding/gob"

	"github.com/DonggyuLim/erc20/db"
	"github.com/gin-gonic/gin"
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
func GetToken(c *gin.Context, tokenName string) *Token {
	value, ok := db.Get(tokenName)
	if !ok {
		c.String(400, "Don't exsits")

	}
	t := ByteToToken(value)

	return t
}
