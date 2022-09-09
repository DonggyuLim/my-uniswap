package router

import (
	db "github.com/DonggyuLim/erc20/db"
	"github.com/DonggyuLim/erc20/grc20"
	"github.com/DonggyuLim/erc20/utils"
	"github.com/gin-gonic/gin"
)

// rest body json -> data:struct bind
func BindJson(c *gin.Context, data interface{}) {
	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.String(400, err.Error())
	}
}

// tokenName -> db -> token
func GetToken(c *gin.Context, token string) *grc20.Token {
	value, ok := db.Get(token)
	if !ok {
		c.String(400, "Not Address")
	}

	t := utils.ByteToToken(value)
	return t
}
