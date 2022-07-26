package main

import (
	"testing"
)

func BenchmarkCGO(b *testing.B) {
	b.Run("CGO Assign", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			obj := malloc()
			free(obj)
		}
	})
	b.Run("CGO Free Assign", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			obj := "test"
			obj = ""
			obj = obj
		}
	})

}
