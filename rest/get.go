package rest

import (
	"fmt"
	"net/http"

	u "github.com/DonggyuLim/grc20/utils"
	"github.com/gin-gonic/gin"
)

const (
	success = "success"
	fail    = "fail"
	GET     = "GET"
	POST    = "POST"
)

type Description struct {
	URL         string `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Example     string `json:"example,omitempty"`
}

func Document(c *gin.Context) {
	data := []Description{
		{
			URL:         "/",
			Method:      GET,
			Description: "See Documentation",
		},
		{
			URL:         "/deploy",
			Method:      POST,
			Description: "Contract deploy",
		},
		{
			URL:         "/balance",
			Method:      GET,
			Description: "balance is token amount of your account",
		},
		{
			URL:         "/allowance",
			Method:      GET,
			Description: "allowance of your account check ",
		},
		{
			URL:         "/mint",
			Method:      POST,
			Description: "mint append token of contract",
		},

		{
			URL:         "/transfer",
			Method:      POST,
			Description: "Your Balance of Account Send To Account",
		},
		{
			URL:         "/approve",
			Method:      POST,
			Description: "Balance Burn",
		},
		{
			URL:         "/transferfrom",
			Method:      POST,
			Description: "TransferFrom send balance of spender",
		},
	}
	c.JSON(http.StatusOK, data)
}

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
