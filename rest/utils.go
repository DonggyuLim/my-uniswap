package rest

import (
	"github.com/DonggyuLim/erc20/Interface"
	"github.com/DonggyuLim/erc20/grc20"
)

func SaveToken(name string, token Interface.GRC20) {
	Interface.SaveToken(name, token)
}

func GetToken(tokenName string) (*grc20.Token, error) {
	t, err := grc20.GetToken(tokenName)

	return t, err
}
