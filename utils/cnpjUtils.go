package utils

import (
	"regexp"
	"strconv"
)

// CNPJUtils : Valida e higieniza o número de CNPJ
func CNPJUtils(cnpj string) (string, bool) {

	clean := regexp.MustCompile(`[^\d]`)
	cleanCNPJ := clean.ReplaceAllString(cnpj, "")

	// Retorn falso se o CNPJ não possuir 14 dígitos
	if len(cleanCNPJ) != 14 {
		return cleanCNPJ, false
	}

	// Conferência do primeiro dígito verificador
	mult := []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	sumFirstDigit := 0
	for j := 0; j < 12; j++ {
		digit, _ := strconv.Atoi(cleanCNPJ[j : j+1])
		sumFirstDigit = sumFirstDigit + digit*(mult[j])
	}

	resto := sumFirstDigit % 11
	realFirstVerifyingDigit := 0

	if resto >= 2 {
		realFirstVerifyingDigit = 11 - resto
	}

	firstVerifyingDigit, _ := strconv.Atoi(cleanCNPJ[12:13])
	firstCheck := realFirstVerifyingDigit == firstVerifyingDigit

	if firstCheck == false {
		return cleanCNPJ, false
	}

	// Conferência do segundo dígito verificador
	mult2 := []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	sumSecDigit := 0
	for j := 0; j < 13; j++ {
		digit, _ := strconv.Atoi(cleanCNPJ[j : j+1])
		sumSecDigit = sumSecDigit + digit*(mult2[j])
	}

	resto2 := sumSecDigit % 11
	realSecVerifyingDigit := 0

	if resto2 >= 2 {
		realSecVerifyingDigit = 11 - resto2
	}

	secVerifyingDigit, _ := strconv.Atoi(cleanCNPJ[13:14])
	secCheck := realSecVerifyingDigit == secVerifyingDigit

	if secCheck == false {
		return cleanCNPJ, false
	}

	return cleanCNPJ, true

}
