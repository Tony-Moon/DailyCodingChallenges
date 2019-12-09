package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
RomVal is a global map containing all the values of Roman Numerals from I (1) to
M (1000). They are saved as binary values from ANSCII.
*/
var RomVal = map[byte]int{
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

	numArray := Str2Array(strings.ToUpper(romNum)) // Ensure its uppercase!
	finalSum := CompareSum(numArray)
	fmt.Println("In Decimal, your Roman Numeral is: " + strconv.Itoa(finalSum))
}

/*
Str2Array converts a string of Roman Numerals to an array of the corropsonding values.
This happens by looking at the ANSCII value of each letter and finding the corresponding
"worth" in the romNum map.
*/
func Str2Array(romNum string) []int {
	strLen := len(romNum)
	var numArray = make([]int, strLen)
	for i := range numArray {
		numArray[i] = RomVal[romNum[i]]
	}
	return numArray
}

/*
CompareSum compares and adds the numbers of the array produced by the Str2Array function.
This function works from right to left, starting with the right-most value and deciding
wether to add or subtract left value. For example; XIV will start with 5, subtract 1 leaving
4, then add 10 for a total of 14.
*/
func CompareSum(numArray []int) int {
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
