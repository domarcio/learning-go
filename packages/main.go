package main

import (
	"fmt"

	"github.com/domarcio/packages/age"
	"github.com/domarcio/packages/preffix"
)

// Username is the name of the Xerify
var Username = "Woody"

func main() {
	completeUsername := preffix.Preffix + " " + Username
	fmt.Printf("Username: %s\n", completeUsername)
	fmt.Printf("Age: %d\n", age.Age)
}
