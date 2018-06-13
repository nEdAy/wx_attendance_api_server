package main

import (
	"testing"
)

func Test_Main(t *testing.T) {
	main()
}

func BenchmarkMain(b *testing.B) {
	main()
}
