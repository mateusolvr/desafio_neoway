package utils

import (
	"fmt"
	"regexp"
	"strconv"
)

// CPFUtils : Validate and clean the CPF numbers
func CPFUtils(cpf string) (string, bool) {

	clean := regexp.MustCompile(`[^\d]`)
	cleanCPF := clean.ReplaceAllString(cpf, "")
	// Pad zeros to the left
	cleanCPF = fmt.Sprintf("%011v", cleanCPF)

	// Return false if the CPF has more than 11 numbers
	if len(cleanCPF) > 11 {
		return cleanCPF, false
	}

	// Check the first verifying digit
	sumFirstDigit := 0
	for j := 0; j < 9; j++ {
		digit, _ := strconv.Atoi(cleanCPF[j : j+1])
		sumFirstDigit = sumFirstDigit + digit*(10-j)
	}

	firstVerifyingDigit, _ := strconv.Atoi(cleanCPF[9:10])
	firstCheck := ((sumFirstDigit * 10) % 11) == firstVerifyingDigit

	if firstCheck == false {
		return cleanCPF, false
	}

	// Check the second verifying digit
	sumSecDigit := 0
	for j := 0; j < 10; j++ {
		digit, _ := strconv.Atoi(cleanCPF[j : j+1])
		sumSecDigit = sumSecDigit + digit*(11-j)
	}

	secVerifyingDigit, _ := strconv.Atoi(cleanCPF[10:11])
	secCheck := ((sumSecDigit * 10) % 11) == secVerifyingDigit

	if secCheck == false {
		return cleanCPF, false
	}

	// Check if the numbers are all the same
	actualDigit := cleanCPF[0:1]
	actualCheck := true

	for i := 1; i < len(cleanCPF); i++ {
		// Compare each number on the CPF
		currentDigit := cleanCPF[i : i+1]
		currentCheck := currentDigit == actualDigit
		actualCheck = currentCheck && actualCheck
		actualDigit = currentDigit
	}

	if actualCheck == true {
		return cleanCPF, false
	}

	// If all checks are correct, return true
	return cleanCPF, true

}
