package token

import "github.com/shopspring/decimal"

func NewToken(name, symbol string, dec uint8) *Token {
	bm := make(map[string]decimal.Decimal)
	am := make(map[string]decimal.Decimal)

	t := &Token{
		Name:        name,
		Symbol:      symbol,
		Decimal:     dec,
		TotalSupply: decimal.NewFromInt(0),
		Balance:     bm,
		Allowances:  am,
	}
	return t
}
