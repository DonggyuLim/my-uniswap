package utils

import (
	"bytes"
	"encoding/gob"
	"strconv"

	"github.com/DonggyuLim/erc20/grc20"
)

// byte -> string
func ByteToString(data []byte) string {
	return string(data[:])
}

// string -> uint
func StringToUint(data string) uint64 {
	uint, err := strconv.ParseUint(data, 10, 32)
	if err != nil {
		panic(err)
	}
	return uint
}

// uint -> string
func UintToString(data uint64) string {
	return strconv.FormatUint(data, 10)
}

// struct -> byte
func StructToByte(data interface{}) []byte {
	var result bytes.Buffer
	enc := gob.NewEncoder(&result)
	err := enc.Encode(data)
	if err != nil {
		panic(err)
	}
	return result.Bytes()
}

// byte -> Token
func ByteToToken(data []byte) *grc20.Token {
	var token *grc20.Token
	encoder := gob.NewDecoder(bytes.NewBuffer(data))
	err := encoder.Decode(&token)
	if err != nil {
		panic(err)
	}
	return token
}
