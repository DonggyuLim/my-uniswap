package rest

import (
	"fmt"

	u "github.com/DonggyuLim/erc20/utils"
	"github.com/gin-gonic/gin"
)

type BalanceRequest struct {
	TokenName string `json:"tokenName"`
	Account   string `json:"account"`
}

func BalanceOf(c *gin.Context) {
	r := BalanceRequest{}
	err := c.ShouldBindJSON(&r)
	fmt.Println("r==========", r)
	if err != nil {
		c.String(400, err.Error())
	}
	t, err := u.GetToken(r.TokenName)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	balance := t.BalanceOf(r.Account)

	c.JSON(200, gin.H{
		"account": r.Account,
		"balance": balance,
	})
}

// 토큰 정보
func TokenInfo(c *gin.Context) {
	tn := c.Param("name")

	t, err := u.GetToken(tn)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	c.JSON(200, gin.H{
		"tokenName":   t.GetName(),
		"symbol":      t.GetSymbol(),
		"decimal":     t.GetDecimal(),
		"totalSupply": t.GetTotalSupply(),
	})
}

type allowanceRequest struct {
	TokenName string `json:"tokenName"`
	Owner     string `json:"owner"`
	Spender   string `json:"spender"`
}

func Allowance(c *gin.Context) {
	r := allowanceRequest{}
	err := c.ShouldBindJSON(&r)
	fmt.Println("r==========", r)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	t, err := u.GetToken(r.TokenName)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	t.Allowance(r.Owner, r.Spender)

	c.JSON(200, gin.H{
		"allowance": t.Allowance(r.Owner, r.Spender),
	})
}
