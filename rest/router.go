package rest

import (
	"net/http"

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
