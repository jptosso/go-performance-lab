package main

import "testing"

func BenchmarkUseVariables(b *testing.B) {
	// spoiler alert, results are the same
	b.Run("Use unassigned variables", func(b *testing.B) {
		a := 0
		for i := 0; i < b.N; i++ {
			a += a + i
		}
	})
	b.Run("Use assigned variables", func(b *testing.B) {
		a := 0
		for i := 0; i < b.N; i++ {
			c := a + i
			a += c
		}
	})
}

func BenchmarkCreateVSReuse(b *testing.B) {
	// spoiler alert, results are the same
	b.Run("Hardcoded string", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a := "some test"
			a = a
		}
	})
	b.Run("Copy string", func(b *testing.B) {
		c := "some test"
		for i := 0; i < b.N; i++ {
			a := c
			a = a
		}
	})
}
