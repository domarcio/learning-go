package main

import (
	"fmt"

	"github.com/domarcio/basic_functions/math"
)

func main() {
	result := math.Calculate(math.Multiplier, 5, 5)
	fmt.Printf("Mutiplier 5 * 5: %d\n", result)

	result = math.Add(5, 7)
	fmt.Printf("Sum 5 + 7: %d\n", result)

	result, remainder := math.Divider(12, 5)
	fmt.Printf("The division result is %d and your remainder is %d\n", result, remainder)
}
