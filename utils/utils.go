package utils

import (
	"bytes"
	"encoding/gob"

	"strconv"

	"github.com/DonggyuLim/grc20/Interface"
	"github.com/DonggyuLim/grc20/token"
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

// GRC20struct -> byte
func DataToByte(data Interface.GRC20) []byte {
	var result bytes.Buffer
	enc := gob.NewEncoder(&result)
	err := enc.Encode(data)
	if err != nil {
		panic(err)
	}
	return result.Bytes()
}

func SaveToken(name string, token Interface.GRC20) {
	Interface.SaveToken(name, token)
}

func GetToken(tokenName string) (*token.Token, error) {
	t, err := token.GetToken(tokenName)

	return t, err
}
