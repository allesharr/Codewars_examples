package bitcounting

import (
	"strconv"
)

//Write a function that takes an integer as input, and returns the number of bits that are equal to one in the binary representation of that number. You can guarantee that input is non-negative.

//Example: The binary representation of 1234 is 10011010010, so the function should return 5 in this case

func Count(num uint64) int {
	var count int = 0
	s := strconv.FormatUint(num, 2)
	for _, element := range s {
		if element == '1' {
			count++
		}

	}
	return count
}

func CountInt(num int64) int {
	var count int = 0
	s := strconv.FormatInt(num, 2)
	for _, element := range s {
		if element == '1' {
			count++
		}

	}
	return count
}
