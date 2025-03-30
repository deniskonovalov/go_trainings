package lesson2

import (
	"fmt"
	"strconv"
)

func FibonacciIterative(n int) int {
	// If n is less than or equal to 0, return 0
	if n <= 0 {
		return 0
	}
	// The first Fibonacci number for 1 is 1
	if n == 1 {
		return 1
	}

	// Init the first two Fibonacci numbers.
	prev := 0
	curr := 1

	// Iterate from 2 to n to compute subsequent Fibonacci numbers.
	for i := 2; i <= n; i++ {
		// Save the current value in a temporary variable.
		temp := curr
		// Calculate the new current value as the sum of the previous two.
		curr += prev
		// Update the previous value to the old current value.
		prev = temp
	}

	// Return the computed Fibonacci number for the given n.
	return curr
}

func FibonacciRecursive(n int) int {
	// If n is less than or equal to 0, return 0
	if n <= 0 {
		return 0
	}
	// The first Fibonacci number for 1 is 1
	if n == 1 {
		return 1
	}

	// Recursively calculate the Fibonacci number by summing the two preceding values.
	return FibonacciRecursive(n-1) + FibonacciRecursive(n-2)
}

func IsPrime(n int) bool {
	// If n is less or equal 1 it is not prime
	if n <= 1 {
		return false
	}

	// In the loop check if n has at least 1 divider more than 1 and less than n
	// then n is not prime
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}

	// Then n is prime
	return true
}

func IsBinaryPalindrome(n int) bool {
	// Convert int to binary representation
	numInBinary := strconv.FormatInt(int64(n), 2)
	// Use bytes not runes because byte is enough for our case
	bytes := []byte(numInBinary)

	// Revers string in the loop
	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}

	reversed := string(bytes)

	// Check if reversed string is equal to original string
	// then num is palindrome, else - no
	if reversed == numInBinary {
		return true
	}

	return false
}

func ValidParentheses(s string) bool {
	// Stack for storing open parentheses
	stack := []rune{}

	// Map of open-closed parentheses
	brackets := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	// Iterate each char in string
	for _, char := range s {
		// If char is open parentheses, then add it to stack
		if char == '(' || char == '[' || char == '{' {
			stack = append(stack, char)
		} else if char == ')' || char == ']' || char == '}' {
			// If it's closed parentheses then check if stack is empty
			// or if we don't have open parentheses for it
			if len(stack) == 0 || stack[len(stack)-1] != brackets[char] {
				return false
			}
			// Use slice to remove last item from our stack
			stack = stack[:len(stack)-1]
		}
	}

	// At the end stack should be empty
	return len(stack) == 0
}

func Increment(num string) int {
	// Convert binary string to the decimal number
	decimalNum, err := strconv.ParseInt(num, 2, 64)

	// If input was invalid then return error
	if err != nil {
		fmt.Println(err)
		return 0
	}

	// Increment the decimal number
	decimalNum++

	// Return incremented decimal
	return int(decimalNum)
}
