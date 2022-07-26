package main

import (
	"strings"
	"testing"
)

func BenchmarkStringCaseChanging(b *testing.B) {
	// for ASCII only
	a := "Hello World"
	b.Run("ASCII strings.To...", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			c := strings.ToLower(a)
			c = c
		}
	})
	b.Run("ASCII rune by rune", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			c := []byte(a)
			for i, r := range a {
				if r >= 65 && r <= 90 {
					c[i] = byte(r + 32)
				}
			}
		}
	})
}

func BenchmarkStringToByte(b *testing.B) {
	b.Run("String to byte", func(b *testing.B) {
		a := "Hello World"
		for i := 0; i < b.N; i++ {
			a := []byte(a)
			a = a
		}
	})
	b.Run("Bytes to string", func(b *testing.B) {
		a := []byte("Hello World")
		for i := 0; i < b.N; i++ {
			b := string(a)
			b = b
		}
	})
}
