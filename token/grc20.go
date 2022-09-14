package token

import "github.com/shopspring/decimal"

type Token struct {
	Name        string
	Symbol      string
	Decimal     uint8
	TotalSupply decimal.Decimal
	Balance     map[string]decimal.Decimal
	Allowances  map[string]decimal.Decimal
}

// type address string

// func StringToAddress(data string) address {
// 	return address(data)
// }

// /////////////////
// ////////////////
// ////////////////
// Query
func (t *Token) GetName() string {
	return t.Name
}
func (t *Token) GetSymbol() string {
	return t.Symbol
}
func (t *Token) GetTotalSupply() decimal.Decimal {
	return t.TotalSupply
}

func (t *Token) BalanceOf(account string) decimal.Decimal {
	return t.Balance[account]
}

func (t *Token) GetDecimal() uint8 {
	return t.Decimal
}

func (t *Token) Allowance(owner, spender string) decimal.Decimal {
	return t.allowance(owner, spender)
}
func (t *Token) allowance(owner, spender string) decimal.Decimal {
	key := owner + ":" + spender
	return t.Allowances[key]
}

// /////////////////////////////////////////////////
// /////////////////////////////////////////////////
// /////////////////////////////////////////////////
// /////////////////////////////////////////////////
// Mutate

func (t *Token) Transfer(from, to string, amount decimal.Decimal) error {
	err := t.checkBalance(from, amount)
	if err != nil {
		return err
	}
	t.transfer(from, to, amount)
	return nil
}
func (t *Token) transfer(from, to string, amount decimal.Decimal) {

	fromBalance := t.Balance[from]
	t.Balance[from] = fromBalance.Sub(amount)
	toBalance := t.Balance[to]
	t.Balance[to] = toBalance.Add(amount)
}

func (t *Token) Approve(owner, spender string, amount decimal.Decimal) error {
	if err := t.checkBalance(owner, amount); err != nil {
		return err
	}
	t.approve(owner, spender, amount)
	return nil
}

func (t *Token) approve(owner, spender string, amount decimal.Decimal) error {

	key := owner + ":" + spender
	currentBalance := t.Allowances[key]
	t.Allowances[key] = currentBalance.Add(amount)
	return nil
}

func (t *Token) TransferFrom(from, to, spender string, amount decimal.Decimal) error {
	if err := t.checkspendAllowance(from, spender, amount); err != nil {
		return err
	}
	t.transfer(from, to, amount)
	return nil
}

func (t *Token) Mint(account string, amount decimal.Decimal) {
	t.mint(account, amount)
}

func (t *Token) mint(address string, amount decimal.Decimal) {

	t.TotalSupply = t.TotalSupply.Add(amount)
	currentBalance := t.BalanceOf(address)
	// newBalance := currentBalance + amount
	newBalance := currentBalance.Add(amount)
	t.Balance[address] = newBalance
}

func (t *Token) Burn(address string, amount decimal.Decimal) {
	t.burn(address, amount)
}

func (t *Token) burn(address string, amount decimal.Decimal) {
	currentBalance := t.BalanceOf(address)
	newBalance := currentBalance.Sub(amount)
	t.Balance[address] = newBalance
}

// ///////////////////////////////////
// //////////////////////////////////
