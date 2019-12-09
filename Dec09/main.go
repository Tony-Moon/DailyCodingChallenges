package main

import (
	"fmt"
	"strconv"
)

// Set global values for roman numerals using the ANSCII integers for letters
var romVal = map[byte]int{
	73: 1,    // I
	86: 5,    // V
	88: 10,   // X
	76: 50,   // L
	67: 100,  // C
	68: 500,  // D
	77: 1000, // M
}

func main() {
	var romNum string
	fmt.Println("Welcome to the Roman Numeral converter.")
	fmt.Println("Roman Numeral: ")
	fmt.Scan(&romNum)

	numArray := str2Array(romNum)
	finalSum := compareSum(numArray)
	fmt.Println("In Decimal, your Roman Numeral is: " + strconv.Itoa(finalSum))
}

func str2Array(romNum string) []int {
	strLen := len(romNum)
	var numArray = make([]int, strLen)
	for i := range numArray {
		numArray[i] = romVal[romNum[i]]
	}
	return numArray
}

func compareSum(numArray []int) int {
	var righ, left, summ int
	fmt.Println(len(numArray) - 1)

	for i := len(numArray) - 1; i >= 0; i-- {

		if i == (len(numArray) - 1) {
			righ = 0
			left = numArray[i]
		} else {
			righ = numArray[i+1]
			left = numArray[i]
		}

		if i == len(numArray) {
			// Last Number
			summ = left
		} else if left >= righ {
			summ += left
		} else if left < righ {
			summ -= left
		}
	}
	return summ
}
