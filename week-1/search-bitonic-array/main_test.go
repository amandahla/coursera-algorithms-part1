package main

import (
	"bytes"
	"log"
	"strings"
	"testing"
)

func Test_Main(t *testing.T) {
	for i := 0; i < 100; i++ {
		var s bytes.Buffer
		log.SetOutput(&s)
		log.SetFlags(0)
		main()
		output := s.String()
		if !strings.Contains(output, "Found") {
			t.Error(s.String())
		}
	}
}
