package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	success = "success"
	fail    = "fail"
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
			Method:      "GET",
			Description: "See Documentation",
		},
		{
			URL:         "/deploy",
			Method:      "POST",
			Description: "Contract deploy",
			Example:     `{name:token-name,symbol:token-symbol,totalSupply:token-supply(uint)}`,
		},
		{
			URL:         "/mint",
			Method:      "POST",
			Description: "Contract Mint",
			Example:     `{name:token-nam",account:Contract-address,balance:"account-balance}`,
		},
		{
			URL:         "/balance",
			Method:      "GET",
			Description: "Query Balance of Account",
			Example:     "http://localhost:8080/balance?name=0xb",
		},
		{
			URL:         "/transfer",
			Method:      "POST",
			Description: "Your Balance of Account Send To Account",
			Example:     `{to:0xa,from:0xb,amount:100}`,
		},
	}
	c.JSON(http.StatusOK, data)
}

// type deploy struct {
// 	Name        string `json:"name"`
// 	Symbol      string `json:"symbol"`
// 	TotalSupply uint64 `json:"totalSupply"`
// }

// // 배포
// func Deploy(c *gin.Context) {
// 	d := deploy{}
// 	err := c.ShouldBindJSON(&d)
// 	if err != nil {
// 		c.String(400, err.Error())
// 	}
// 	db.ChangeDBName(d.Name)
// 	db.NewDB()
// 	db.Add("name", []byte(d.Name))
// 	db.Add("symbol", []byte(d.Symbol))
// 	db.Add("totalSupply", []byte(strconv.FormatUint(d.TotalSupply, 10)))

// 	c.JSON(200, gin.H{
// 		"message":     success,
// 		"name":        d.Name,
// 		"symbol":      d.Symbol,
// 		"totalSupply": d.TotalSupply,
// 	})
// }

// type mint struct {
// 	Name    string `json:"name"`
// 	Account string `json:"account"`
// 	Balance uint64 `json:"balance"`
// }

// // 민팅
// func Mint(c *gin.Context) {
// 	m := mint{}
// 	err := c.ShouldBindJSON(&m)
// 	if err != nil {
// 		c.String(400, err.Error())
// 	}
// 	db.Add(m.Account, []byte(strconv.FormatUint(m.Balance, 10)))
// 	c.JSON(200, gin.H{
// 		"message": success,
// 		"account": m.Account,
// 		"balance": strconv.FormatUint(m.Balance, 10),
// 	})
// }

// func BalanceOf(c *gin.Context) {
// 	account := c.Query("name")
// 	value, ok := db.Get(account)
// 	if !ok {
// 		c.String(400, "Not have balance")
// 	}

// 	c.JSON(200, gin.H{
// 		"account": u.ByteToString(value),
// 	})
// }

// type trasfer struct {
// 	To     string
// 	From   string
// 	Amount uint64
// }

// func Transfer(c *gin.Context) {
// 	t := trasfer{}
// 	err := c.ShouldBindJSON(&t)
// 	if err != nil {
// 		c.String(400, err.Error())
// 	}
// 	err = transfer(t.From, t.To, t.Amount)
// 	if err != nil {
// 		c.String(400, fmt.Sprint(err))
// 	}
// 	fb, ok := db.Get(t.From)
// 	if !ok {
// 		c.String(400, fmt.Sprint(err))
// 	}

// 	c.JSON(200, gin.H{
// 		"message": "success",
// 		"balance": u.ByteToString(fb),
// 	})
// }

// func transfer(from, to string, amount uint64) error {
// 	//from balance
// 	fb, ok := db.Get(from)

// 	if !ok {
// 		return errors.New("not have balance")
// 	}
// 	//to balance
// 	tb, ok := db.Get(to)
// 	if !ok {
// 		return errors.New("don't exsits account")
// 	}

// 	newfb := u.StringToUint(u.ByteToString(fb)) - amount
// 	newtb := u.StringToUint(u.ByteToString(tb)) + amount
// 	db.Add(from, []byte(strconv.FormatUint(newfb, 10)))
// 	db.Add(to, []byte(strconv.FormatUint(newtb, 10)))
// 	return nil
// }
