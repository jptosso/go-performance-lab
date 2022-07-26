package main

import (
	"sync"
	"testing"
)

type test struct {
	data string
}

var pool = sync.Pool{
	New: func() interface{} {
		return &test{}
	},
}

func BenchmarkAllocateTestPool(b *testing.B) {
	b.Run("Allocate from pool", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			t := pool.Get().(*test)
			t.data = "test"
			pool.Put(t)
		}
	})
	b.Run("Allocate without pool", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			t := &test{}
			t.data = "test"
			t = t
		}
	})
}

func BenchmarkAllocateWithMux(b *testing.B) {
	b.Run("Allocate with mux", func(b *testing.B) {
		mutex := sync.Mutex{}
		for i := 0; i < b.N; i++ {
			mutex.Lock()
			t := &test{}
			t.data = "test"
			t = t
			mutex.Unlock()
		}
	})
	b.Run("Allocate without mux", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			t := &test{}
			t.data = "test"
			t = t
		}
	})
}
