package main

import "testing"

type test3 struct {
	data int
}

func t1(t test3) {
	t.data = 1
}

func t2(t *test3) {
	t.data = 1
}
func BenchmarkShareStruct(b *testing.B) {
	t := test3{}
	b.Run("Send Struct", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			t1(t)
		}
	})
	b.Run("Send Struct Pointer", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			t2(&t)
		}
	})
}
