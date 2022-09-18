package token

import (
	"bytes"
	"encoding/gob"
	"errors"

	"github.com/DonggyuLim/grc20/db"
)

// byte -> Token
func ByteToToken(data []byte) *Token {
	var token *Token
	decoder := gob.NewDecoder(bytes.NewBuffer(data))
	err := decoder.Decode(&token)
	if err != nil {
		panic(err)
	}
	return token
}

// tokenName -> db -> token
func GetToken(tokenName string) (*Token, error) {
	value, ok := db.Get(tokenName)
	if !ok {
		return &Token{}, errors.New("don't Get token")

	}
	t := ByteToToken(value)

	return t, nil
}
