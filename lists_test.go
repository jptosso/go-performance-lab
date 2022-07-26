package main

import "testing"

func BenchmarkListCreation(b *testing.B) {
	b.Run("Create Array", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var a [10000]int
			a[9999] = 1
		}
	})
	b.Run("Create Slice", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a := make([]int, 10000)
			a[9999] = 1
		}
	})
	b.Run("Create Slice from nothing", func(b *testing.B) {
		c := make([]int, 1)
		for i := 0; i < b.N; i++ {
			a := c[0:0]
			a = a
		}
	})
	b.Run("Create Slice from something", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a := []int{}
			a = a
		}
	})
}

func BenchmarkGrowSlice(b *testing.B) {
	b.Run("Grow 1 by 1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a := []int{}
			for j := 0; j < 10000; j++ {
				a = append(a, j)
			}
			a[9999] = 1
		}
	})
	b.Run("Grow in chunks using append", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a := []int{}
			for j := 0; j < 10000; j++ {
				if j%100 == 0 {
					a = append(a, make([]int, 100)...)
				}
				a[j] = j
			}
			a[9999] = 1
		}
	})
	b.Run("Preallocate", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a := make([]int, 10000)
			for j := 0; j < 10000; j++ {
				a[j] = j
			}
			a[9999] = 1
		}
	})
}

func BenchmarkSliceAppend(b *testing.B) {
	b.Run("Array to slice", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var a = [10000]int{}
			for j := 0; j < 10000; j++ {
				c := make([]int, 10000)
				for k := 0; k < 10000; k++ {
					c[k] = a[k]
				}
			}
		}
	})
	b.Run("Slice from air", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for j := 0; j < 10000; j++ {
				a := []int{}
				for k := 0; k < 10000; k++ {
					a = append(a, k)
				}
			}
		}
	})
}

func BenchmarkDestroySlice(b *testing.B) {
	// spoiler alert, results are the same
	b.Run("Destroy slice after use", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a := make([]int, 10000)
			a = a[0:0]
		}
	})
	b.Run("Use assigned variables", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a := make([]int, 10000)
			a = a
		}
	})
}
