package main

import (
	"unicode"
)

/*
 Validates payment card like MasterCard, Visa, etc.
*/
func IsPaymentCard(value string) bool {
	var strLenght = len(value)
	
	if strLenght < 13 || strLenght > 19 {
		return false
	}

	var sum uint8
	var digit uint8
	var isOddCardLength = strLenght % 2 == 0
	var isOdd bool

	for i, c := range value {
		if !unicode.IsDigit(c) {
			return false
		}

		digit = uint8(c-'0')
		isOdd = i % 2 == 0

		if (isOddCardLength && isOdd) || (!isOddCardLength && !isOdd) {
			digit*= 2

			if digit > 9 {
				digit-= 9
			}
		}

		sum+= digit
	}

	return sum % 10 == 0
}
