package utils

import (
	"fmt"
	"regexp"
)

// CleanCNPJ : Clean the CNPJ numbers
func CleanCNPJ(cnpj string) string {

	clean := regexp.MustCompile(`[^\d]`)
	cleanCNPJ := clean.ReplaceAllString(cnpj, "")
	// Pad zeros to the left
	if cleanCNPJ != "" {
		cleanCNPJ = fmt.Sprintf("%014v", cleanCNPJ)
	}

	return cleanCNPJ

}
