package main

import (
	"regexp"
)

/*
 Validates payment card like MasterCard, Visa, etc.
*/
func IsPaymentCard(value string) bool {
	re, _ := regexp.Compile(`^\d{13,19}$`)

	if !re.MatchString(value) {
		return false
	}

	var sum uint8
	var digit uint8
	var isOddCardLength = len(value) % 2 == 0
	var isOdd bool

	for i, c := range value {
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
