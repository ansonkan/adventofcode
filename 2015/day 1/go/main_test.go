package main

import (
	"testing"
)

func Benchmark_Par1_ReadFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part1_ReadFile("../input.txt")
	}
}

func Benchmark_Par1_ReadFile_large(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part1_ReadFile("../input_large.txt")
	}
}

func Benchmark_Par1_Seek(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part1_Seek("../input.txt")
	}
}

func Benchmark_Par1_Seek_large(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part1_Seek("../input_large.txt")
	}
}
