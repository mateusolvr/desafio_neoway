package utils

import (
	"database/sql"
	"regexp"
	"strings"
)

// StringUtils : Teste de biblioteca
func StringUtils(s string) sql.NullString {

	clean := regexp.MustCompile(`[\d]`)
	checkValue := clean.ReplaceAllString(s, "")

	s = strings.ToLower(s)
	if len(s) == 0 || s == "null" || s == "none" {
		return sql.NullString{}
	} else if checkValue == "," {
		s = strings.Replace(s, ",", ".", -1)
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}

}
