package hub

import (
	"fmt"
	"testing"
)

func TestCmd_Status(t *testing.T) {
	hub := Hub(false)
	hub.Status()
}

func BenchmarkCmd_Status(b *testing.B) {
	hub := Hub(false)
	b.StartTimer()
	hub.Status()
	b.StopTimer()
	b.ReportAllocs()
}

func ExampleCmd_Status() {
	fmt.Println("Run hub status with stdout redirect")
	hub := Hub(true)
	hub.Status()
}
