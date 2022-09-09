package router

import (
	db "github.com/DonggyuLim/erc20/db"
	"github.com/DonggyuLim/erc20/utils"
	"github.com/gin-gonic/gin"
)

type BalanceRequest struct {
	Name    string
	Account string
}

func BalanceOf(c *gin.Context) {
	b := BalanceRequest{}
	BindJson(c, b)
	t := GetToken(c, b.Name)
	balance := t.BalanceOf(b.Account)
	db.Add(b.Name, utils.StructToByte(t))

	c.JSON(200, gin.H{
		"account": b.Account,
		"balance": balance,
	})
}

type tokenInfoRequest struct {
	TokenName string
}

// 토큰 정보
func TokenInfo(c *gin.Context) {
	r := tokenInfoRequest{}
	BindJson(c, r)
	t := GetToken(c, r.TokenName)

	c.JSON(200, gin.H{
		"tokenName":   t.GetName(),
		"symbol":      t.GetSymbol(),
		"decimal":     t.GetDecimal(),
		"totalSupply": t.GetTotalSupply(),
	})
}

type allowanceRequest struct {
	tokenName string
	owner     string
	spender   string
}

func Allowance(c *gin.Context) {
	r := allowanceRequest{}
	BindJson(c, r)
	t := GetToken(c, r.tokenName)
	t.Allowance(r.owner, r.spender)
	c.JSON(200, gin.H{
		"allowance": t.Allowance(r.owner, r.spender),
	})
}
