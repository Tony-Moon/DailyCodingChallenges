package main

import (
	"fmt"
	"log"
)

// N is the number of stairs in string form
var N string

// X is the number of steps you can take
var X = [2]int{1, 2}

func main() {
	fmt.Println("What would you like the number of stairs to be?")
	_, err := fmt.Scanln(&N)
	if err != nil {
		log.Fatal("Error scanning input:", err)
	}
	fmt.Println("There are " + N + " stairs.")
}
