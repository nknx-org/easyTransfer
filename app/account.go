package app

import (
	"fmt"

	nkn "github.com/nknorg/nkn-sdk-go"
)

// InitializeAccount will create an account file and return it. If there is already an account in place it will just return the existing account.
func InitializeAccount() *nkn.Account {
	account, err := nkn.NewAccount(nil)

	if err != nil {
		fmt.Println(err)
	}
	return account
}
