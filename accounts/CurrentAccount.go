package accounts

import (
	"bank/users"
)

type CurrentAccount struct {
	Owner          users.Owner
	AgencyNumber   int
	AccountNumber  int
	accountBalance float64
}

func (c *CurrentAccount) Withdraw(withdrawValue float64) string {
	canWithdraw := withdrawValue > 0 && withdrawValue <= c.accountBalance
	if canWithdraw {
		c.accountBalance -= withdrawValue
		return "successful withdrawal"
	} else {
		return "not enought money"
	}
}

func (c *CurrentAccount) Deposit(depositValue float64) string {
	if depositValue > 0 {
		c.accountBalance += depositValue
		return "successful deposit"
	} else {
		return "Deposit value less than 0"
	}
}

func (c *CurrentAccount) Transfer(transferValue float64, targetAccount *CurrentAccount) bool {
	if transferValue < c.accountBalance && transferValue > 0 {
		c.accountBalance -= transferValue
		targetAccount.accountBalance += transferValue
		return true
	} else {
		return false
	}
}
func (c *CurrentAccount) BankStatement() float64 {
	return c.accountBalance
}
