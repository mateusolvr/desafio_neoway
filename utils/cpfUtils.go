package utils

import (
	"fmt"
	"regexp"
)

// CleanCPF : Clean the CPF numbers
func CleanCPF(cpf string) string {

	clean := regexp.MustCompile(`[^\d]`)
	cleanCPF := clean.ReplaceAllString(cpf, "")
	// Pad zeros to the left
	cleanCPF = fmt.Sprintf("%011v", cleanCPF)

	return cleanCPF

}
