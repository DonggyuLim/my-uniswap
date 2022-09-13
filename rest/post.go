package rest

import (
	"fmt"

	"github.com/DonggyuLim/erc20/Interface"
	"github.com/DonggyuLim/erc20/db"
	"github.com/DonggyuLim/erc20/grc20"
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
	err := c.ShouldBindJSON(&r)
	fmt.Println("r==========", r)
	if err != nil {
		c.String(400, err.Error())
	}
	_, ok := db.Get(r.TokenName)
	fmt.Println(ok)
	if ok {
		c.String(400, "Exists Token")
		return
	} else {
		t := grc20.NewToken(r.TokenName, r.Symbol, r.Decimal)

		t.Mint(r.Account, r.TotalSupply)

		Interface.SaveToken(r.TokenName, t)
		c.JSON(200, gin.H{
			"message":     success,
			"name":        t.GetName(),
			"symbol":      t.GetSymbol(),
			"totalSupply": t.GetTotalSupply(),
			"decimal":     t.GetDecimal(),
		})
	}
}

type mintRequest struct {
	TokenName string `json:"tokenName"`
	Account   string `json:"account"`
	Amount    uint64 `json:"amount"`
}

func Mint(c *gin.Context) {
	r := mintRequest{}
	err := c.ShouldBindJSON(&r)

	if err != nil {
		c.String(400, err.Error())
		return
	}

	t, err := GetToken(r.TokenName)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	t.Mint(r.Account, r.Amount)
	SaveToken(r.TokenName, t)
	c.JSON(200, gin.H{
		"message":      "success",
		"totalBalance": t.GetTotalSupply(),
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
	err := c.ShouldBindJSON(&r)

	if err != nil {
		c.String(400, err.Error())
		return
	}
	t, err := GetToken(r.TokenName)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	t.Transfer(r.From, r.To, r.Amount)
	SaveToken(r.TokenName, t)
	c.JSON(200, gin.H{
		"account": r.From,
		"balance": t.BalanceOf(r.From),
	})
}

type approveRequest struct {
	TokenName string `json:"tokenName"`
	Owner     string `json:"owner"`
	Spender   string `json:"spender"`
	Amount    uint64 `json:"amount"`
}

func Approve(c *gin.Context) {
	r := approveRequest{}
	err := c.ShouldBindJSON(&r)

	if err != nil {
		c.String(400, err.Error())
		return
	}
	t, err := GetToken(r.TokenName)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	t.Approve(r.Owner, r.Spender, r.Amount)
	SaveToken(r.TokenName, t)
	c.JSON(200, gin.H{
		"message":         success,
		"allowanceAmount": t.Allowance(r.Owner, r.Spender),
	})
}
