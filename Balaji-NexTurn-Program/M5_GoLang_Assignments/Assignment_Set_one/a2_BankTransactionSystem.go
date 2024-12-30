// Case Study:
// You need to simulate a bank transaction system with the following features:
// 1. Account Management: Each account has an ID, name, and balance. Store the
// accounts in a slice.
// 2. Deposit Function: A function to deposit money into an account. Validate if the
// deposit amount is greater than zero.
// 3. Withdraw Function: A function to withdraw money from an account. Ensure the
// account has a sufficient balance before proceeding. Return appropriate errors
// for invalid amounts or insufficient balance.
// 4. Transaction History: Maintain a transaction history for each account as a string
// slice. Use a loop to display the transaction history when requested.
// 5. Menu System: Implement a menu-driven program where users can choose
// actions like deposit, withdraw, view balance, or exit. Use constants for menu
// options and break the loop to exit.

package main

import (
	"errors"
	"fmt"
)

type Account struct {
	ID              int
	Name            string
	Balance         float64
	TransactionHist []string
}

var accounts []Account

func findAccountByID(id int) (*Account, error) {
	for i := range accounts {
		if accounts[i].ID == id {
			return &accounts[i], nil
		}
	}
	return nil, errors.New("account not found")
}

func deposit(accountID int, amount float64) error {
	if amount <= 0 {
		return errors.New("deposit amount must be greater than zero")
	}
	account, err := findAccountByID(accountID)
	if err != nil {
		return err
	}
	account.Balance += amount
	account.TransactionHist = append(account.TransactionHist, fmt.Sprintf("Deposited: %.2f", amount))
	return nil
}

func withdraw(accountID int, amount float64) error {
	if amount <= 0 {
		return errors.New("withdrawal amount must be greater than zero")
	}
	account, err := findAccountByID(accountID)
	if err != nil {
		return err
	}
	if account.Balance < amount {
		return errors.New("insufficient balance")
	}
	account.Balance -= amount
	account.TransactionHist = append(account.TransactionHist, fmt.Sprintf("Withdrew: %.2f", amount))
	return nil
}

func viewTransactionHistory(accountID int) ([]string, error) {
	account, err := findAccountByID(accountID)
	if err != nil {
		return nil, err
	}
	return account.TransactionHist, nil
}

func main() {
	const (
		OptionDeposit  = 1
		OptionWithdraw = 2
		OptionViewHist = 3
		OptionViewBal  = 4
		OptionExit     = 5
	)

	// Seed accounts
	accounts = append(accounts, Account{ID: 1, Name: "Alice", Balance: 1000})
	accounts = append(accounts, Account{ID: 2, Name: "Bob", Balance: 500})

	for {
		fmt.Println("\nBank Transaction System")
		fmt.Println("1. Deposit")
		fmt.Println("2. Withdraw")
		fmt.Println("3. View Transaction History")
		fmt.Println("4. View Balance")
		fmt.Println("5. Exit")
		fmt.Print("Choose an option: ")

		var option int
		fmt.Scan(&option)

		switch option {
		case OptionDeposit:
			var id int
			var amount float64
			fmt.Print("Enter account ID: ")
			fmt.Scan(&id)
			fmt.Print("Enter deposit amount: ")
			fmt.Scan(&amount)
			if err := deposit(id, amount); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Deposit successful")
			}

		case OptionWithdraw:
			var id int
			var amount float64
			fmt.Print("Enter account ID: ")
			fmt.Scan(&id)
			fmt.Print("Enter withdrawal amount: ")
			fmt.Scan(&amount)
			if err := withdraw(id, amount); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Withdrawal successful")
			}

		case OptionViewHist:
			var id int
			fmt.Print("Enter account ID: ")
			fmt.Scan(&id)
			history, err := viewTransactionHistory(id)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Transaction History:")
				for _, record := range history {
					fmt.Println(record)
				}
			}

		case OptionViewBal:
			var id int
			fmt.Print("Enter account ID: ")
			fmt.Scan(&id)
			account, err := findAccountByID(id)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("Current Balance: %.2f\n", account.Balance)
			}

		case OptionExit:
			fmt.Println("Exiting... Goodbye!")
			return

		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
