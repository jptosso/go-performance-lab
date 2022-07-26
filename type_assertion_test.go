package main

import (
	"reflect"
	"testing"
)

func BenchmarkTypeAssertion(b *testing.B) {
	var a interface{}
	a = "sample data"
	b.Run("Type assertion switch", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			switch data := a.(type) {
			case string:
				data = data
			}
		}
	})
	b.Run("Type assertion", func(b *testing.B) {
		// a.(string) creates a copy of a
		for i := 0; i < b.N; i++ {
			c := a.(string)
			c = c
		}
	})
	b.Run("Reflection", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			c := reflect.ValueOf(a).String()
			c = c
		}
	})
	b.Run("No assertion", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			c := a
			c = c
		}
	})

}
