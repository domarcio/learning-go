package main

import "fmt"

func main() {
	foo := false

	if !foo {
		fmt.Printf("Foo is not true\n")
	}

	if description, status := isDelinquencyCustomer(60); status {
		fmt.Printf("What is your situation?: %s\n", description)
	}
}

func isDelinquencyCustomer(days int) (description string, status bool) {
	if days <= 60 {
		description = "Hn, ok. You are delinquency customer :("
		status = true
		return
	}

	description = "Hey! Be happy! You are ok with your finances! :D"
	status = false
	return
}
