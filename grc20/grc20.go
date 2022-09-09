package grc20

type Token struct {
	Name        string
	Symbol      string
	Decimal     uint8
	TotalSupply uint64
	Balance     map[string]uint64
	Allowances  map[string]uint64
}

type address string

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

func StringToAddress(data string) address {
	return address(data)
}

func (t *Token) GetName() string {
	return t.Name
}
func (t *Token) GetSymbol() string {
	return t.Symbol
}
func (t *Token) GetTotalSupply() uint64 {
	return t.TotalSupply
}

func (t *Token) BalanceOf(account string) uint64 {
	return t.Balance[account]
}

func (t *Token) GetDecimal() uint8 {
	return t.Decimal
}

func (t *Token) Transfer(to, from string, amount uint64) {
	t.Balance[from] = t.Balance[from] - amount
	t.Balance[to] = t.Balance[to] + amount
}

//transferFrom 생성해야함.

func (t *Token) Allowance(owner, spender string) uint64 {
	return t.allowance(owner, spender)
}
func (t *Token) allowance(owner, spender string) uint64 {
	key := owner + ":" + spender
	return t.Allowances[key]
}

func (t *Token) Approve(owner, spender string, amount uint64) {
	t.approve(owner, spender, amount)
}

func (t *Token) approve(owner, spender string, amount uint64) {
	key := owner + ":" + spender
	t.Allowances[key] = amount
}

func (t *Token) Mint(to string, amount uint64) {
	t.mint(to, amount)
}

func (t *Token) mint(address string, amount uint64) {
	t.TotalSupply += amount
	currentBalance := t.BalanceOf(address)
	newBalance := currentBalance + amount
	t.Balance[address] = newBalance
}
