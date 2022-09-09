package main

import (
	"github.com/DonggyuLim/erc20/db"
	"github.com/DonggyuLim/erc20/router"
	"github.com/gin-gonic/gin"
)

func main() {
	db.NewDB()
	defer db.Close()
	r := gin.Default()
	//GET
	r.GET("/", router.Document)
	r.GET("/balance", router.BalanceOf)
	r.GET("/allowance", router.Allowance)
	r.GET("/token", router.TokenInfo)
	//POST
	r.POST("/deploy", router.Deploy)
	r.POST("/transfer", router.Transfer)
	r.POST("/approve", router.Approve)
	r.Run()

}
