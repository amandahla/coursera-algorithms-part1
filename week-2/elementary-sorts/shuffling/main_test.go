package main

import (
	"bytes"
	"log"
	"strings"
	"testing"
)

func BenchmarkMain(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		main()
	}
}

func Test_Main(t *testing.T) {
	for i := 0; i < 100; i++ {
		var s bytes.Buffer
		log.SetOutput(&s)
		log.SetFlags(0)
		main()
		output := s.String()
		if !strings.Contains(output, "1 2 3 4 5 6 7 8 9 10") {
			t.Error(s.String())
		}
	}
}
