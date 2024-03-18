package utl

import (
	"fmt"
	"strings"
)

// Case insensitive substring check
func SubString(large, small string) bool {
	return strings.Contains(strings.ToLower(large), strings.ToLower(small))
}

// Split the string and return last element
func LastElem(s, splitter string) string {
	split := strings.Split(s, splitter)
	return split[len(split)-1]
}

// Return first N chars of string n
func FirstN(s string, n int) string {
	if len(s) <= n {
		return s
	}
	i := 0
	for j := range s {
		if i == n {
			return s[:j]
		}
		i++
	}
	return s
}

// Return the best printable string value for given x variable
func Str(x interface{}) string {
	if x == nil {
		return ""
	}
	switch GetType(x) {
	case "bool":
		return fmt.Sprintf("%t", x)
	case "string":
		return x.(string)
	default:
		return "" // Blank for other types
	}
}

// Returns string value of unknown variable, wrapped in single quotes. This is a special
// function to lookout for leading '*', which YAML does not allow and must be single-quoted
func StrSingleQuote(x interface{}) string {
	s := Str(x)
	if strings.HasPrefix(s, "*") {
		return "'" + s + "'"
	}
	return s
}

// Converts any value to its string representation using default formatting.
func ToStr(value interface{}) string {
	return fmt.Sprintf("%v", value)
}

// ItemInList checks if a given string (arg) is present in a list of strings (argList).
// Returns true if found, false otherwise.
func ItemInList(arg string, argList []string) bool {
	for _, value := range argList {
		if value == arg {
			return true
		}
	}
	return false
}

// Return string of spaces for padded printing. Needed when printing terminal colors.
// Colorize output uses % sequences that conflict with Printf's own formatting with %
func PadSpaces(targetWidth, stringWidth int) string {
	padding := targetWidth - stringWidth
	if padding > 0 {
		return fmt.Sprintf("%*s", padding, " ")
	} else {
		return ""
	}
}

// Return value as a string, padded with leading spaces, totalling width size wide. This is
// needed when printing terminal colors, because they conflict with Printf's own '%' formatting
func PreSpc(value interface{}, width int) string {
	str := ToStr(value)
	padding := width - len(str)
	if padding > 0 {
		return fmt.Sprintf("%*s%s", padding, " ", str)
	} else {
		return str
	}
}

// Return value as a string, padded with trailing spaces, totalling width size wide. This is
// needed when printing terminal colors, because they conflict with Printf's own '%' formatting
func PostSpc(value interface{}, width int) string {
	str := ToStr(value)
	padding := width - len(str)
	if padding > 0 {
		return fmt.Sprintf("%s%*s", str, padding, " ")
	} else {
		return str
	}
}
