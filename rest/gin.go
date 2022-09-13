package rest

import (
	"sync"

	"github.com/gin-gonic/gin"
)

const RestPort = "8080"

func Rest(wg *sync.WaitGroup) {
	r := gin.Default()
	//GET
	r.GET("/", Document)
	r.GET("/balance", BalanceOf)
	r.GET("/allowance", Allowance)
	r.GET("/token/:name", TokenInfo)

	//POST
	r.POST("/deploy", Deploy)
	r.POST("/mint", Mint)
	r.POST("/transfer", Transfer)
	r.POST("/approve", Approve)
	r.Run(RestPort)
	defer wg.Done()
}
