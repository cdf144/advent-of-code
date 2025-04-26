package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

// Init reads input lines from a specified file path and returns them along with a
// function to print the answer.
//
// The path parameter specifies the location of the input file. If the path is
// relative, it is resolved relative to the directory of the calling function's
// source file. If the path is absolute, it is used directly.
func Init(path string) ([]string, func(s any)) {
	if !filepath.IsAbs(path) {
		_, filename, _, ok := runtime.Caller(1)
		if !ok {
			panic("Could not get caller info")
		}
		dir := filepath.Dir(filename)
		path = filepath.Join(dir, path)
	}

	lines := readInputLines(path)
	printAnswer := makePrintAnswerFunc()

	return lines, printAnswer
}

// InitRaw reads the content from a specified file path and returns it along with a
// function to print the answer.
//
// The path parameter specifies the location of the input file. If the path is
// relative, it is resolved relative to the directory of the calling function's
// source file. If the path is absolute, it is used directly.
func InitRaw(path string) (string, func(s any)) {
	if !filepath.IsAbs(path) {
		_, filename, _, ok := runtime.Caller(1)
		if !ok {
			panic("Could not get caller info")
		}
		dir := filepath.Dir(filename)
		path = filepath.Join(dir, path)
	}

	content := readInputRaw(path)
	printAnswer := makePrintAnswerFunc()

	return content, printAnswer
}

// Filter takes a slice of any type and a function.
// It returns a new slice containing only the elements from the input slice for which the
// function returns true.
func Filter[T any](slice []T, f func(T) bool) []T {
	var result []T
	for _, v := range slice {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

// Atoi converts a string to an integer.
// It is a convenience wrapper around strconv.Atoi that panics if the string cannot be
// converted to an integer.
func Atoi(s string) int {
	return UnwrapInt(strconv.Atoi(s))
}

func UnwrapString(s string, err error) string {
	if err != nil {
		panic(err)
	}
	return s
}

func UnwrapInt(i int, err error) int {
	if err != nil {
		panic(err)
	}
	return i
}

func UnwrapBytes(b []byte, err error) []byte {
	if err != nil {
		panic(err)
	}
	return b
}

func makePrintAnswerFunc() func(s any) {
	return func(s any) {
		boldGreen := "\033[1;32m"
		resetBoldGreen := "\033[0m"
		fmt.Println(boldGreen + "Answer: " + resetBoldGreen + fmt.Sprint(s))
	}
}

func readInputLines(path string) []string {
	return strings.Split(
		strings.TrimSpace(readInputRaw(path)),
		"\n",
	)
}

func readInputRaw(path string) string {
	return string(
		UnwrapBytes(
			os.ReadFile(strings.TrimSpace(path)),
		),
	)
}
