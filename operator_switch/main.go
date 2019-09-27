package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	rand.Seed(5)
	randomNumber := rand.Intn(5)
	fmt.Print("The number random is ")
	switch randomNumber {
	case 5:
		fmt.Println("five.")
	case 4:
		fmt.Println("four.")
	case 3:
		fmt.Println("three.")
	case 2:
		fmt.Println("two.")
	case 1:
		fmt.Println("one.")
	}

	fmt.Print("My system family is ")
	switch runtime.GOOS {
	case "linux":
		fmt.Printf("linux ")
		fallthrough
	default:
		fmt.Println("because I like.")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Ahaul! Weekend is now!")
	case time.Friday:
		fmt.Println("Ok, today is friday, my friend.")
	default:
		fmt.Println("Let's to work.")
	}

	switch {
	case randomNumber == 1:
		fmt.Println("Random number is 1.")
	case randomNumber > 1:
		fmt.Println("Random number is greater than 1.")
	}
}
