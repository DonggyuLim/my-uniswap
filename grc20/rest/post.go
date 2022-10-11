package rest

import (
	"fmt"

	"github.com/DonggyuLim/grc20/Interface"
	"github.com/DonggyuLim/grc20/db"
	"github.com/DonggyuLim/grc20/token"
	u "github.com/DonggyuLim/grc20/utils"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

type AccountResponse struct {
	TokenName string `json:"tokenName"`
	Account   string `json:"account"`
	Balance   string `json:"balance"`
	Allowance string `json:"Allowance"`
}

type deployRequest struct {
	TokenName   string `json:"tokenName"`
	Symbol      string `json:"symbol"`
	TotalSupply string `json:"totalSupply"`
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
		t := token.NewToken(r.TokenName, r.Symbol, r.Decimal)
		totalSupply := decimal.RequireFromString(r.TotalSupply)
		// t.Mint(r.Account, r.TotalSupply)
		t.Mint(r.Account, totalSupply)

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
	Amount    string `json:"amount"`
}

func Mint(c *gin.Context) {
	r := mintRequest{}
	err := c.ShouldBindJSON(&r)

	if err != nil {
		c.String(400, err.Error())
		return
	}

	t, err := u.GetToken(r.TokenName)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	amount := decimal.RequireFromString(r.Amount)
	t.Mint(r.Account, amount)
	u.SaveToken(r.TokenName, t)
	c.JSON(200, gin.H{
		"message":     "success",
		"totalSupply": t.GetTotalSupply(),
		"account": AccountResponse{
			TokenName: t.GetName(),
			Account:   r.Account,
			Balance:   t.BalanceOf(r.Account).String(),
		},
	})
}

type trasferRequest struct {
	TokenName string
	To        string
	From      string
	Amount    string
}

func Transfer(c *gin.Context) {
	r := trasferRequest{}
	err := c.ShouldBindJSON(&r)

	if err != nil {
		c.String(400, err.Error())
		return
	}
	t, err := u.GetToken(r.TokenName)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	amount := decimal.RequireFromString(r.Amount)
	err = t.Transfer(r.From, r.To, amount)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	u.SaveToken(r.TokenName, t)
	c.JSON(200, gin.H{
		"from": AccountResponse{
			TokenName: t.GetName(),
			Account:   r.From,
			Balance:   t.BalanceOf(r.From).String(),
		},
		"to": AccountResponse{
			TokenName: t.GetName(),
			Account:   r.To,
			Balance:   t.BalanceOf(r.To).String(),
		},
		"balance": t.BalanceOf(r.From),
	})
}

type approveRequest struct {
	TokenName string `json:"tokenName"`
	Owner     string `json:"owner"`
	Spender   string `json:"spender"`
	Amount    string `json:"amount"`
}

func Approve(c *gin.Context) {
	r := approveRequest{}
	err := c.ShouldBindJSON(&r)

	if err != nil {
		c.String(400, err.Error())
		return
	}
	t, err := u.GetToken(r.TokenName)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	amount := decimal.RequireFromString(r.Amount)
	err = t.Approve(r.Owner, r.Spender, amount)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	u.SaveToken(r.TokenName, t)
	c.JSON(200, gin.H{
		"message": success,
		"account": t.Allowance(r.Owner, r.Spender),
	})
}

type transferFromRequest struct {
	TokenName string `json:"tokenName"`
	Onwer     string `json:"owner"`
	Spender   string `json:"spender"`
	To        string `json:"to"`
	Amount    string `json:"amount"`
}

func TransferFrom(c *gin.Context) {
	r := transferFromRequest{}
	err := c.ShouldBindJSON(&r)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	t, err := u.GetToken(r.TokenName)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	amount := decimal.RequireFromString(r.Amount)
	err = t.TransferFrom(r.Onwer, r.Spender, r.To, amount)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	u.SaveToken(r.TokenName, t)
	c.JSON(200, gin.H{
		"message": success,

		"to": AccountResponse{
			TokenName: t.GetName(),
			Account:   r.To,
			Balance:   t.BalanceOf(r.To).String(),
		},
	})
}

type burnRequest struct {
	TokenName string `json:"tokenName"`
	Account   string `json:"account"`
	Amount    string `json:"amount"`
}

func Burn(c *gin.Context) {
	r := burnRequest{}
	err := c.ShouldBindJSON(&r)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	t, err := u.GetToken(r.TokenName)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	amount := decimal.RequireFromString(r.Amount)
	t.Burn(r.Account, amount)
	u.SaveToken(r.TokenName, t)
	c.JSON(200, gin.H{
		"message":     success,
		"totalSupply": t.GetTotalSupply(),
		"account": AccountResponse{
			Account: r.Account,
			Balance: t.BalanceOf(r.Account).String(),
		},
	})
}
