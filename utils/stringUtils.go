package utils

import (
	"database/sql"
	"regexp"
	"strings"
)

// StringUtils : Transform NULL and NONEs into nil and replace "," in decimal values with "."
func StringUtils(s string) sql.NullString {

	// Clean all the number to see if there's an "," left, if there's it's probably a decimal value
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
