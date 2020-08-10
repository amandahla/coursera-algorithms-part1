package main

import (
        "bytes"
        "log"
        "testing"
)

func Test_Main(t *testing.T) {
        for i:=0; i < 100; i++ {
                var s bytes.Buffer
                log.SetOutput(&s)
                log.SetFlags(0)
                main()
                if s.String() != "false\n" {
                        t.Error(s.String())
                }
        }
}

