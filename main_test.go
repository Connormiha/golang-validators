package main

import (
	"testing"
)

func TestIsPaymentCard(t *testing.T) {
	var result bool
	var invalidValues = []string{"1", "99990", "5360455961388881", "5360455961388882", "5844097252647043"}

	for _, value := range invalidValues {
		result = IsPaymentCard(value)

		if result {
			t.Error("Should be invalid ", value)
		}
	}
	
	var validValues = []string{
		// Master card
		"5360455961388880", "2720994098260002", "5387855391779104",
		// Visa
		"4916571723272436", "4716921009028345", "4556686337398685366",
		// Visa Electron
		"4508913026101389", "4917585515144008", "4844097252647043",
		// American Express (AMEX)
		"341768978049054", "372583415581029", "344528240888797",
		// Maestro
		"6763952592774687", "5018039841847830", "6759341529211235",
	}
	
	for _, value := range validValues {
		result = IsPaymentCard(value)

		if !result {
			t.Error("Should be valid ", value)
		}
	}
}

func BenchmarkIsPaymentCard(b *testing.B) {
	var validValues = []string{
		// Invalid
		"1", "99990", "5360455961388881", "5360455961388882", "5844097252647043",
		// Master card
		"5360455961388880", "2720994098260002", "5387855391779104",
		// Visa
		"4916571723272436", "4716921009028345", "4556686337398685366",
		// Visa Electron
		"4508913026101389", "4917585515144008", "4844097252647043",
		// American Express (AMEX)
		"341768978049054", "372583415581029", "344528240888797",
		// Maestro
		"6763952592774687", "5018039841847830", "6759341529211235",
	}

	b.Logf("b.N is %d\n", b.N)

	for i := 0; i < b.N; i++ {
		for _, value := range validValues {
			IsPaymentCard(value)
		}
	}
}