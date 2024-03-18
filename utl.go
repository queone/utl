// Package utl is a library of very common housekeeping functions. Many of these
// functions seem very simple and almost unnecessary, but they can make code
// much easier to read and digest.
package utl

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"sort"

	"github.com/google/uuid"
)

// Same as regular Printf function but always exits, with a return code of 1
func Die(format string, args ...interface{}) {
	fmt.Printf(format, args...)
	os.Exit(1)
}

// Returns returns a string showing the current filepath line number and function name.
// See https://stackoverflow.com/questions/25927660/how-to-get-the-current-function-name
func Trace() string {
	progCounter, fp, ln, ok := runtime.Caller(1)
	if !ok {
		return fmt.Sprintf("%s\n    %s:%d\n", "?", "?", 0)
	}
	funcPointer := runtime.FuncForPC(progCounter)
	if funcPointer == nil {
		return fmt.Sprintf("%s\n    %s:%d\n", "?", fp, ln)
	}
	return fmt.Sprintf("%s\n    %s:%d\n", funcPointer.Name(), fp, ln)
}

// Returns true if given string is a valid UUID number. False otherwise.
func ValidUuid(s string) bool {
	_, err := uuid.Parse(s)
	return err == nil
}

// Check if two variables are of the same type
func SameType(a, b interface{}) bool {
	a_type := fmt.Sprintf("%T", a)
	b_type := fmt.Sprintf("%T", b)
	return a_type == b_type
}

func GetType(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

// Checks if given rune is an alphabetic character (either uppercase or lowercase).
// Returns true if it is, false otherwise.
func IsAlpha(c rune) bool {
	if ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z') {
		return true
	}
	return false
}

// Returns true if rune is a numerical digit. False otherwise.
func IsDigit(c rune) bool {
	if '0' <= c && c <= '9' {
		return true
	}
	return false
}

// Returns true if rune is a hexadeximal digit. False otherwise.
func IsHexDigit(c rune) bool {
	if ('0' <= c && c <= '9') || ('a' <= c && c <= 'f') || ('A' <= c && c <= 'F') {
		return true
	}
	return false
}

// Return the map string object's keys sorted
func SortMapStringKeys(obj map[string]string) (sortedKeys []string) {
	sortedKeys = make([]string, 0, len(obj))
	for k := range obj {
		sortedKeys = append(sortedKeys, k)
	}
	sort.Strings(sortedKeys)
	return sortedKeys
}

// Return the object's keys sorted
func SortObjStringKeys(obj map[string]interface{}) (sortedKeys []string) {
	sortedKeys = make([]string, 0, len(obj))
	for k := range obj {
		sortedKeys = append(sortedKeys, k)
	}
	sort.Strings(sortedKeys)
	return sortedKeys
}

// TODO: Combine above SortMapStringKeys and SortObjStringKeys using interfaces
// SortStringKeys()?

// Print prompt message and return single rune character input
func PromptMsg(msg string) rune {
	fmt.Print(Yel(msg))
	reader := bufio.NewReader(os.Stdin)
	confirm, _, err := reader.ReadRune()
	if err != nil {
		fmt.Println(err)
	}
	return confirm
}
