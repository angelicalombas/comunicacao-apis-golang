package utils

import (
	"regexp"
	"strconv"
	"unicode"
)

// isValidCPF verifies if a CPF number is valid
func IsValidCPF(cpf string) (bool, string) {
	// Verifica se contém alguma letra
	if match, _ := regexp.MatchString("[a-zA-Z]", cpf); match {
		return false, "Invalid format"
	}

	// Remove non-digit characters
	var numbers []int
	for _, r := range cpf {
		if unicode.IsDigit(r) {
			n, _ := strconv.Atoi(string(r))
			numbers = append(numbers, n)
		}
	}

	if len(numbers) != 11 {
		return false, "Invalid length"
	}

	// Check if all digits are equal
	allSame := true
	for i := 1; i < 11 && allSame; i++ {
		if numbers[i] != numbers[0] {
			allSame = false
		}
	}
	if allSame {
		return false, "All digits are the same"
	}

	// Validate first check digit
	sum := 0
	for i := 0; i < 9; i++ {
		sum += numbers[i] * (10 - i)
	}
	firstCheckDigit := (sum * 10 % 11) % 10
	if firstCheckDigit != numbers[9] {
		return false, "Invalid check digit"
	}

	// Validate second check digit
	sum = 0
	for i := 0; i < 10; i++ {
		sum += numbers[i] * (11 - i)
	}
	secondCheckDigit := (sum * 10 % 11) % 10
	if secondCheckDigit != numbers[10] {
		return false, "Invalid check digit"
	}

	cpfInt := ""
	for _, n := range numbers {
		cpfInt += strconv.Itoa(n)
	}

	return true, cpfInt
}
