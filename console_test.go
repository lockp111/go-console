package console

import (
	"testing"
	"time"
)

var tc = &consoleSpace{}

func TestPrintln(t *testing.T) {
	tc.Println("==")
	tc.Println("===")
	tc.Println("=")
}

func TestProgress(t *testing.T) {
	n1 := tc.Println("===")
	n2 := tc.Println("===")
	n3 := tc.Println("===")
	time.Sleep(time.Second)
	tc.Progress(n1, "========")
	tc.Println(n1)
	tc.Progress(n2, "=")
	tc.Println(n2)
	tc.Progress(n3, "======================")
}
