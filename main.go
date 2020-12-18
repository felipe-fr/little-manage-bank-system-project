package main

import (
	"fmt"

	"bank/accounts"
	"bank/users"
)

func main() {
	userFelipe := users.Owner{Name: "Felipe", Id: "45236987415", Occupation: "Developer"}
	felipeAccount := accounts.CurrentAccount{Owner: userFelipe,
		AgencyNumber: 222, AccountNumber: 1111222}
	felipeAccount.Deposit(1000)
	userAnna := users.Owner{Name: "Anna", Id: "87541965874", Occupation: "Teacher"}
	annaAccount := accounts.CurrentAccount{Owner: userAnna,
		AgencyNumber: 223, AccountNumber: 1111223}
	annaAccount.Deposit(400)
	fmt.Println(felipeAccount)
	fmt.Println(annaAccount)
	fmt.Println("********************")
	fmt.Println(felipeAccount.BankStatement())
	fmt.Println(felipeAccount.Withdraw(1000))
	fmt.Println(felipeAccount.BankStatement())
	fmt.Println(felipeAccount.Deposit(600))
	fmt.Println(felipeAccount.BankStatement())
	fmt.Println("**************************")
	fmt.Println("Anna:", annaAccount.BankStatement())
	fmt.Println("Felipe:", felipeAccount.BankStatement())
	fmt.Println(felipeAccount.Transfer(300, &annaAccount))
	fmt.Println("Anna:", annaAccount.BankStatement())
	fmt.Println("Felipe:", felipeAccount.BankStatement())
}
