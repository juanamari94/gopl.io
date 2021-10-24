// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package popcount_test

import (
	"testing"

	"gopl.io/ch2/popcount"
)

// -- Alternative implementations --

func BitCount(x uint64) int {
	// Hacker's Delight, Figure 5-2.
	x = x - ((x >> 1) & 0x5555555555555555)
	x = (x & 0x3333333333333333) + ((x >> 2) & 0x3333333333333333)
	x = (x + (x >> 4)) & 0x0f0f0f0f0f0f0f0f
	x = x + (x >> 8)
	x = x + (x >> 16)
	x = x + (x >> 32)
	return int(x & 0x7f)
}

func PopCountByClearing(x uint64) int {
	n := 0
	for x != 0 {
		x = x & (x - 1) // clear rightmost non-zero bit
		n++
	}
	return n
}

func PopCountByShifting(x uint64) int {
	n := 0
	for i := uint(0); i < 64; i++ {
		if x&(1<<i) != 0 {
			n++
		}
	}
	return n
}

// -- Benchmarks --

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountLoop(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountShift64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountShift64(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountKernighan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountKernighan(0x1234567890ABCDEF)
	}
}

func BenchmarkBitCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BitCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountByClearing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByClearing(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountByShifting(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByShifting(0x1234567890ABCDEF)
	}
}

type popCountPair struct {
	input    uint64
	expected int
}

func TestPopCount(t *testing.T) {
	cases := []popCountPair{{1, 1}, {2, 1}, {3, 2}, {127, 7}, {65535, 16}}
	for _, testCase := range cases {
		result := popcount.PopCount(testCase.input)
		if result != testCase.expected {
			t.Errorf("Expected %v got %v", testCase.expected, result)
		}
	}
}

func TestPopCountLoop(t *testing.T) {
	cases := []popCountPair{{1, 1}, {2, 1}, {3, 2}, {127, 7}, {65535, 16}}
	for _, testCase := range cases {
		result := popcount.PopCountLoop(testCase.input)
		if result != testCase.expected {
			t.Errorf("Expected %v got %v", testCase.expected, result)
		}
	}
}

func TestPopCountShift64(t *testing.T) {
	cases := []popCountPair{{1, 1}, {2, 1}, {3, 2}, {127, 7}, {65535, 16}}
	for _, testCase := range cases {
		result := popcount.PopCountShift64(testCase.input)
		if result != testCase.expected {
			t.Errorf("Expected %v got %v", testCase.expected, result)
		}
	}
}

func TestPopCountKernighan(t *testing.T) {
	cases := []popCountPair{{1, 1}, {2, 1}, {3, 2}, {127, 7}, {65535, 16}}
	for _, testCase := range cases {
		result := popcount.PopCountKernighan(testCase.input)
		if result != testCase.expected {
			t.Errorf("Expected %v got %v", testCase.expected, result)
		}
	}
}

// goos: darwin
// goarch: amd64
// cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
// BenchmarkPopCount-16                    1000000000               0.2402 ns/op
// BenchmarkPopCountLoop-16                244007541                4.855 ns/op
// BenchmarkBitCount-16                    1000000000               0.2408 ns/op
// BenchmarkPopCountByClearing-16          121914792                9.892 ns/op
// BenchmarkPopCountByShifting-16          76343859                16.07 ns/op
