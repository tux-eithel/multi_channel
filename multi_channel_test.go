package main

import (
	"runtime"
	"testing"
)

var data []string
var totalCPU int = 4

func init() {
	data = prepareData()
}

func BenchmarkAll(b *testing.B) {

	for i := 0; i < b.N; i++ {
		initAll(data, totalCPU)
	}

}

func BenchmarkSequential(b *testing.B) {

	for i := 0; i < b.N; i++ {
		initSequential(data, totalCPU)
	}

}

func BenchmarkMulti(b *testing.B) {

	for i := 0; i < b.N; i++ {
		initMultiChannel(data, totalCPU)
	}

}

func BenchmarkAllCPU(b *testing.B) {

	runtime.GOMAXPROCS(runtime.NumCPU())

	for i := 0; i < b.N; i++ {
		initAll(data, totalCPU)
	}

}

func BenchmarkSequentialCPU(b *testing.B) {

	runtime.GOMAXPROCS(runtime.NumCPU())

	for i := 0; i < b.N; i++ {
		initSequential(data, totalCPU)
	}

}

func BenchmarkMultiCPU(b *testing.B) {

	runtime.GOMAXPROCS(runtime.NumCPU())

	for i := 0; i < b.N; i++ {
		initMultiChannel(data, totalCPU)
	}

}
