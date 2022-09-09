package router

import (
	db "github.com/DonggyuLim/erc20/db"
	"github.com/DonggyuLim/erc20/grc20"
	"github.com/DonggyuLim/erc20/utils"
	"github.com/gin-gonic/gin"
)

// rest body json -> data:struct bind

// tokenName -> db -> token
func GetToken(c *gin.Context, token string) *grc20.Token {
	value, ok := db.Get(token)
	if !ok {
		c.String(400, "Don't exsits")

	}
	t := utils.ByteToToken(value)

	return t
}

// save data
func SaveToken(name string, token *grc20.Token) {
	db.Add(name, utils.StructToByte(token))
}
