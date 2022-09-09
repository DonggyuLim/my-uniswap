package router

import (
	"github.com/DonggyuLim/erc20/db"
	"github.com/DonggyuLim/erc20/grc20"
	"github.com/DonggyuLim/erc20/utils"
	"github.com/gin-gonic/gin"
)

type deployRequest struct {
	TokenName   string `json:"tokenName"`
	Symbol      string `json:"symbol"`
	TotalSupply uint64 `json:"totalSupply"`
	Decimal     uint8  `json:"decimal"`
	Account     string `json:"account"`
}

// 배포
func Deploy(c *gin.Context) {
	r := deployRequest{}
	BindJson(c, r)

	t := grc20.NewToken(r.TokenName, r.Symbol, r.Decimal)

	t.Mint(r.Account, r.TotalSupply)

	db.Add(r.TokenName, utils.StructToByte(t))
	c.JSON(200, gin.H{
		"message":     success,
		"name":        r.TokenName,
		"symbol":      r.Symbol,
		"totalSupply": r.TotalSupply,
		"decimal":     r.Decimal,
	})
}

type trasferRequest struct {
	TokenName string
	To        string
	From      string
	Amount    uint64
}

func Transfer(c *gin.Context) {
	r := trasferRequest{}
	BindJson(c, r)
	t := GetToken(c, r.TokenName)
	t.Transfer(r.To, r.From, r.Amount)
	db.Add(r.TokenName, utils.StructToByte(t))
	c.JSON(200, gin.H{
		"account": r.From,
		"balance": t.BalanceOf(r.From),
	})
}

type approveRequest struct {
	tokenName string
	owner     string
	spender   string
	amount    uint64
}

func Approve(c *gin.Context) {
	r := approveRequest{}
	BindJson(c, r)
	t := GetToken(c, r.tokenName)
	t.Approve(r.owner, r.spender, r.amount)
	c.JSON(200, gin.H{
		"message":         success,
		"allowanceAmount": t.Allowance(r.owner, r.spender),
	})
}
