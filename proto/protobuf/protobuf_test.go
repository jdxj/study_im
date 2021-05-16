package protobuf

import (
	"fmt"
	"testing"

	"github.com/valyala/bytebufferpool"
)

func BenchmarkCopy(b *testing.B) {
	pt1 := make([]byte, 2)
	pt2 := make([]byte, 4096)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf := make([]byte, 2+4096)
		copy(buf, pt1)
		copy(buf[2:], pt2)
	}
}

func BenchmarkCopyWithBuffer(b *testing.B) {
	pt1 := make([]byte, 2)
	pt2 := make([]byte, 4096)
	buf := make([]byte, 2+4096)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		copy(buf, pt1)
		copy(buf[2:], pt2)
	}
}

func BenchmarkAppend(b *testing.B) {
	pt1 := make([]byte, 2)
	pt2 := make([]byte, 4096)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var buf []byte
		buf = append(buf, pt1...)
		buf = append(buf, pt2...)
	}
}

func BenchmarkAppendWithBuf(b *testing.B) {
	pt1 := make([]byte, 2)
	pt2 := make([]byte, 4096)
	buf := make([]byte, 0, 2+4096)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf = buf[:0]
		buf = append(buf, pt1...)
		buf = append(buf, pt2...)
	}
}

func BenchmarkBufferPool(b *testing.B) {
	pt1 := make([]byte, 2)
	pt2 := make([]byte, 4096)
	pool := bytebufferpool.Pool{}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf := pool.Get()
		buf.Write(pt1)
		buf.Write(pt2)
		buf.Bytes()
		pool.Put(buf)
	}
}

func TestProcessor_Marshal(t *testing.T) {
	p := NewProcessor()
	p.Register(0, &Test{})
	data, err := p.Marshal(&Test{
		Name: "abc",
		Age:  123,
	})
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("data: %v\n", data)

	_, msg, err := p.Unmarshal(data)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%s\n", msg)
}
