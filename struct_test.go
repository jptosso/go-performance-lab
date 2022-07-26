package main

import "testing"

type test2 struct {
	data  string
	data1 [100]int
	data2 [100]interface{}
	data3 []bool
}

func BenchmarkNewStructPointer(b *testing.B) {
	b.Run("New Struct Simple", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			t := &test2{}
			t.data = "test"
			t.data1[0] = 1
			t.data2[0] = 1
			t.data3 = []bool{true}
		}
	})
	b.Run("New Struct Single obj", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			t := &test2{
				data:  "test",
				data1: [100]int{1},
				data2: [100]interface{}{1},
				data3: []bool{true},
			}
			t = t
		}
	})
	b.Run("New Struct new()", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			t := new(test2)
			t.data = "test"
			t.data1[0] = 1
			t.data2[0] = 1
			t.data3 = []bool{true}
		}
	})
}

func BenchmarkNewStructNoPointer(b *testing.B) {
	b.Run("New Struct Simple", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			t := test2{}
			t.data = "test"
			t.data1[0] = 1
			t.data2[0] = 1
			t.data3 = []bool{true}
		}
	})
	b.Run("New Struct Single obj", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			t := test2{
				data:  "test",
				data1: [100]int{1},
				data2: [100]interface{}{1},
				data3: []bool{true},
			}
			t = t
		}
	})
	b.Run("New Struct new()", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			t := new(test2)
			t.data = "test"
			t.data1[0] = 1
			t.data2[0] = 1
			t.data3 = []bool{true}
		}
	})
}
