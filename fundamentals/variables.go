package main

import "fmt"

var (
	phone     = "999 999 999"
	address   string  // undeclared address is same as address = ""
	quantity  int     // like a 0
	purchased bool    // like a 0.00
	price     float64 // like a false
	words     rune    // special characters
)

// Variables like a first letter as uppercase are publics, we can access from other Go files.
// var Foo = 1.50

func main() {
	fmt.Printf("Address: %s\n", address)
	fmt.Printf("Quantity: %d\n", quantity)
	fmt.Printf("Price: %v\n", price)
	fmt.Printf("Purchased: %v\n", purchased)
	fmt.Printf("Words: %v\n", words)
}
