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
	TokenName string          `json:"tokenName"`
	Account   string          `json:"account"`
	Balance   decimal.Decimal `json:"balance"`
	Allowance decimal.Decimal `json:"Allowance"`
}

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
		t := token.NewToken(r.TokenName, r.Symbol, r.Decimal)

		// t.Mint(r.Account, r.TotalSupply)
		t.Mint(r.Account, u.UintToDecimal(r.TotalSupply))

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

	t, err := u.GetToken(r.TokenName)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	t.Mint(r.Account, u.UintToDecimal(r.Amount))
	u.SaveToken(r.TokenName, t)
	c.JSON(200, gin.H{
		"message":     "success",
		"totalSupply": t.GetTotalSupply(),
		"account": AccountResponse{
			TokenName: t.GetName(),
			Account:   r.Account,
			Balance:   t.BalanceOf(r.Account),
		},
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
	t, err := u.GetToken(r.TokenName)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	err = t.Transfer(r.From, r.To, u.UintToDecimal(r.Amount))
	if err != nil {
		c.String(400, err.Error())
		return
	}
	u.SaveToken(r.TokenName, t)
	c.JSON(200, gin.H{
		"from": AccountResponse{
			TokenName: t.GetName(),
			Account:   r.From,
			Balance:   t.BalanceOf(r.From),
		},
		"to": AccountResponse{
			TokenName: t.GetName(),
			Account:   r.To,
			Balance:   t.BalanceOf(r.To),
		},
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
	t, err := u.GetToken(r.TokenName)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	err = t.Approve(r.Owner, r.Spender, u.UintToDecimal(r.Amount))
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

	Spender string `json:"spender"`
	To      string `json:"to"`
	Amount  uint64 `json:"amount"`
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
	err = t.TransferFrom(r.Onwer, r.Spender, r.To, u.UintToDecimal(r.Amount))
	if err != nil {
		c.String(400, err.Error())
		return
	}
	u.SaveToken(r.TokenName, t)
	c.JSON(200, gin.H{
		"message": success,
		"from": AccountResponse{
			TokenName: t.GetName(),
			Account:   r.From,
			Balance:   t.BalanceOf(r.From),
		},
		"to": AccountResponse{
			TokenName: t.GetName(),
			Account:   r.To,
			Balance:   t.BalanceOf(r.To),
		},
	})
}

type burnRequest struct {
	TokenName string          `json:"tokenName"`
	Account   string          `json:"account"`
	Amount    decimal.Decimal `json:"amount"`
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
	t.Burn(r.Account, r.Amount)
	u.SaveToken(r.TokenName, t)
	c.JSON(200, gin.H{
		"message":     success,
		"totalSupply": t.GetTotalSupply(),
		"account": AccountResponse{
			Account: r.Account,
			Balance: t.BalanceOf(r.Account),
		},
	})
}
