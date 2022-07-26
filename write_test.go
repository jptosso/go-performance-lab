package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func BenchmarkWriteTo(b *testing.B) {
	b.Run("Write to null", func(b *testing.B) {
		out, _ := os.OpenFile(os.DevNull, os.O_WRONLY, 0)
		defer out.Close()
		for i := 0; i < b.N; i++ {
			out.Write([]byte("Hello World"))
		}
	})
	b.Run("Write to file", func(b *testing.B) {
		out, err := ioutil.TempFile(os.TempDir(), "test")
		defer out.Close()
		if err != nil {
			b.Fatal(err)
		}
		defer out.Close()
		for i := 0; i < b.N; i++ {
			out.Write([]byte("Hello World"))
		}
	})
}
