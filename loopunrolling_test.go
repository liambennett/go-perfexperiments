package main

import "testing"

func BenchmarkLoop1(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		Loop1(10000)
	}
}

func BenchmarkLoop2(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		Loop2(10000)
	}
}

func BenchmarkLoop5(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		Loop5(10000)
	}
}
