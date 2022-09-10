package router

import (
	"github.com/DonggyuLim/erc20/Interface"
	"github.com/DonggyuLim/erc20/grc20"
	"github.com/gin-gonic/gin"
)

func SaveToken(name string, token Interface.GRC20) {
	Interface.SaveToken(name, token)
}

func GetToken(c *gin.Context, tokenName string) *grc20.Token {
	t := grc20.GetToken(c, tokenName)
	return t
}
