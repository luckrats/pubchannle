package pubchannle_test

import (
	"testing"

	. "test/testmain/linkedchannle/pubchannle"
)

func TestXXX(t *testing.T) {

}

func BenchmarkLinkElement___Add(b *testing.B) {
	e := NewLinkElement()
	for i := 0; i < b.N; i++ {
		e = e.Add(i)
	}
	b.ReportAllocs()
}

func BenchmarkNativeChannle_Add(b *testing.B) {
	ch := make(chan interface{}, b.N)
	for i := 0; i < b.N; i++ {
		ch <- i
	}
	b.ReportAllocs()
}
