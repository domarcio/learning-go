package main

import "fmt"

func main() {
	printFibonacciSequence()
	fmt.Printf("\n")

	infinityLoop()
	fmt.Printf("\n")

	printLetterByLetter()
}

func printFibonacciSequence() {
	fmt.Println("### Fibonacci sequence ###")

	var fibonacciNumber int = 1
	var lastResult int = 1
	var next int = 0

	for x := 1; x <= 5; x++ {
		fmt.Println("Number:", fibonacciNumber)
		next = fibonacciNumber + lastResult
		lastResult = fibonacciNumber
		fibonacciNumber = next
	}
}

func infinityLoop() {
	fmt.Println("### Infinity loop ###")

	value := 0
	condition := true

	for condition {
		value++
		if value%5 == 0 {
			condition = false
		}
		fmt.Println("Increment:", value)
	}
	fmt.Printf("--\n")
	for {
		fmt.Println("Decrement:", value)
		value--
		if value == 0 {
			break
		}
	}
}

func printLetterByLetter() {
	fmt.Println("### Show letters ###")

	message := "I love write code using Go"

	for key, letter := range message {
		fmt.Printf("Key[%d] - Letter: %q\n", key, letter)
	}
}
