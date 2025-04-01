package main

import (
	"fmt"
	"learningGo/lesson2"
)

func main() {

	// FibonacciIterative tests
	fmt.Printf("FibonacciIterative(1): %d\n", lesson2.FibonacciIterative(1))
	fmt.Printf("FibonacciIterative(8): %d\n", lesson2.FibonacciIterative(8))
	fmt.Printf("FibonacciIterative(77): %d\n", lesson2.FibonacciIterative(11))

	// FibonacciRecursive tests
	fmt.Printf("FibonacciRecursive(1): %d\n", lesson2.FibonacciRecursive(1))
	fmt.Printf("FibonacciRecursive(8): %d\n", lesson2.FibonacciRecursive(8))
	fmt.Printf("FibonacciRecursive(77): %d\n", lesson2.FibonacciRecursive(11))

	// IsPrime tests
	fmt.Printf("IsPrime(1): %t\n", lesson2.IsPrime(1))
	fmt.Printf("IsPrime(2): %t\n", lesson2.IsPrime(2))
	fmt.Printf("IsPrime(5): %t\n", lesson2.IsPrime(5))
	fmt.Printf("IsPrime(8): %t\n", lesson2.IsPrime(8))
	fmt.Printf("IsPrime(15): %t\n", lesson2.IsPrime(15))
	fmt.Printf("IsPrime(151): %t\n", lesson2.IsPrime(151))
	fmt.Printf("IsPrime(6123842): %t\n", lesson2.IsPrime(6123842))

	//IsBinaryPalindrome tests
	fmt.Printf("lesson2.IsBinaryPalindrome(4): %t\n", lesson2.IsBinaryPalindrome(4))
	fmt.Printf("lesson2.IsBinaryPalindrome(5): %t\n", lesson2.IsBinaryPalindrome(5))
	fmt.Printf("lesson2.IsBinaryPalindrome(9): %t\n", lesson2.IsBinaryPalindrome(9))
	fmt.Printf("lesson2.IsBinaryPalindrome(19): %t\n", lesson2.IsBinaryPalindrome(19))
	fmt.Printf("lesson2.IsBinaryPalindrome(585): %t\n", lesson2.IsBinaryPalindrome(585))

	fmt.Printf("lesson2.IsBinaryPalindromeByTwoIndexes(4): %t\n", lesson2.IsBinaryPalindromeByTwoIndexes(4))
	fmt.Printf("lesson2.IsBinaryPalindromeByTwoIndexes(5): %t\n", lesson2.IsBinaryPalindromeByTwoIndexes(5))
	fmt.Printf("lesson2.IsBinaryPalindromeByTwoIndexes(9): %t\n", lesson2.IsBinaryPalindromeByTwoIndexes(9))
	fmt.Printf("lesson2.IsBinaryPalindromeByTwoIndexes(19): %t\n", lesson2.IsBinaryPalindromeByTwoIndexes(19))
	fmt.Printf("lesson2.IsBinaryPalindromeByTwoIndexes(585): %t\n", lesson2.IsBinaryPalindromeByTwoIndexes(585))

	// ValidParentheses tests
	fmt.Printf("lesson2.ValidParentheses('()[]}{}'): %t\n", lesson2.ValidParentheses("()[]}{}"))
	fmt.Printf("lesson2.ValidParentheses('((one)[two]{three}'): %t\n", lesson2.ValidParentheses("(one)[two]{three}"))
	fmt.Printf("lesson2.ValidParentheses('(one[two]){three}'): %t\n", lesson2.ValidParentheses("(one[two]){three}"))
	fmt.Printf("lesson2.ValidParentheses('(one[two]{three}'): %t\n", lesson2.ValidParentheses("(one[two]{three}"))

	// Increment test
	fmt.Printf("lesson2.Increment('101'): %d\n", lesson2.Increment("101"))
	fmt.Printf("lesson2.Increment('1001'): %d\n", lesson2.Increment("1001"))
	fmt.Printf("lesson2.Increment('10010'): %d\n", lesson2.Increment("10010"))
}
