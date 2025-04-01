package lesson2

import (
	"fmt"
	"math"
	"strconv"
)

func FibonacciIterative(n int) int {

	if n <= 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	prev := 0
	curr := 1

	for i := 2; i <= n; i++ {
		prev, curr = curr, prev+curr
	}

	return curr
}

func FibonacciRecursive(n int) int {
	if n <= 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return FibonacciRecursive(n-1) + FibonacciRecursive(n-2)
}

func IsPrime(n int) bool {
	switch {
	case n <= 1:
		return false
	case n == 2:
		return true
	case n%2 == 0:
		return false
	}

	sqrt := int(math.Sqrt(float64(n)))

	for i := 3; i <= sqrt; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func IsBinaryPalindrome(n int) bool {
	numInBinary := strconv.FormatInt(int64(n), 2)
	bytes := []byte(numInBinary)

	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}

	reversed := string(bytes)

	if reversed == numInBinary {
		return true
	}

	return false
}

func IsBinaryPalindromeByTwoIndexes(n int) bool {
	numInBinary := strconv.FormatInt(int64(n), 2)

	i, j := 0, len(numInBinary)-1
	for i < j {
		if numInBinary[i] != numInBinary[j] {
			return false
		}

		i++
		j--
	}
	return true
}

func ValidParentheses(s string) bool {
	stack := []rune{}

	brackets := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, char := range s {

		switch char {
		case '(', '[', '{':
			stack = append(stack, char)
		case ')', ']', '}':

			if len(stack) == 0 || stack[len(stack)-1] != brackets[char] {

				return false
			}

			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

func Increment(num string) int {
	decimalNum, err := strconv.ParseInt(num, 2, 64)

	if err != nil {
		fmt.Println(err)
		return 0
	}

	decimalNum++

	return int(decimalNum)
}
