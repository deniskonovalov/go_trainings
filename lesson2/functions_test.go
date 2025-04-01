package lesson2

import "testing"

func benchmarkIsBinaryPalindromeReverse(n int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsBinaryPalindrome(n)
	}
}

func benchmarkIsBinaryPalindromeTwoIndexes(n int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsBinaryPalindromeByTwoIndexes(n)
	}
}

func BenchmarkIsBinaryPalindromeReverse5(b *testing.B) {
	benchmarkIsBinaryPalindromeReverse(5, b)
}

func BenchmarkIsBinaryPalindromeReverse45(b *testing.B) {
	benchmarkIsBinaryPalindromeReverse(45, b)
}

func BenchmarkIsBinaryPalindromeReverse585(b *testing.B) {
	benchmarkIsBinaryPalindromeReverse(585, b)
}

func BenchmarkIsBinaryPalindromeTwoIndexes5(b *testing.B) {
	benchmarkIsBinaryPalindromeTwoIndexes(5, b)
}

func BenchmarkIsBinaryPalindromeTwoIndexes45(b *testing.B) {
	benchmarkIsBinaryPalindromeTwoIndexes(45, b)
}

func BenchmarkIsBinaryPalindromeTwoIndexes585(b *testing.B) {
	benchmarkIsBinaryPalindromeTwoIndexes(585, b)
}
