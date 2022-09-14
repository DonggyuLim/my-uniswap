package token

func NewToken(name, symbol string, decimal uint8) *Token {
	bm := make(map[string]uint64)
	am := make(map[string]uint64)

	t := &Token{
		Name:        name,
		Symbol:      symbol,
		Decimal:     decimal,
		TotalSupply: 0,
		Balance:     bm,
		Allowances:  am,
	}
	return t
}
