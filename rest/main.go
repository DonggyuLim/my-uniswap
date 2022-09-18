package rest

import (
	"sync"

	"github.com/gin-gonic/gin"
)

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
	r.POST("/transferFrom", TransferFrom)
	r.POST("/approve", Approve)
	r.POST("/burn", Burn)
	r.Run(":8000")
	defer wg.Done()
}
