package benchmark_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"unsafe"
)

// 在不进行内存分配的情况下，将 []byte 转换为 string
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// 在不进行内存分配的情况下，将 string 转换为 []byte
func StringToBytes(str string) []byte {
	s := (*[2]uintptr)(unsafe.Pointer(&str))
	h := [3]uintptr{s[0], s[1], s[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

type UtilStrBuf struct {
	bytes.Buffer
}
func (this *UtilStrBuf)Str() string {
	buf := this.Bytes()
	return *(*string)(unsafe.Pointer(&buf))
}

func BenchmarkPlus(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		s += "hello world"
		_ = s
	}
}

func BenchmarkFormat(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		s = fmt.Sprintf("%s%s", s, "hello world")
		_ = s
	}
}
func BenchmarkUtilStrBuf(b *testing.B) {
	var buf UtilStrBuf
	for i := 0; i < b.N; i++ {
		buf.WriteString("hello world")
		//_ = buf.String()
		_ = buf.Str()
	}
}

//buf.Bytes()仅作为对比buf.String()
func BenchmarkBufferBytes(b *testing.B) {
	var buf bytes.Buffer
	for i := 0; i < b.N; i++ {
		buf.WriteString("hello world")
		_ = buf.Bytes()
	}
}

func BenchmarkBuilder(b *testing.B) {
	var builder strings.Builder
	for i := 0; i < b.N; i++ {
		builder.WriteString("hello world")
		_ = builder.String()
	}
}

//go test -benchmem -run=^$  -bench=. -v -count=1 -memprofile=mem.out
//go test -bench . -benchmem


/*
go test -run=. -bench=. -benchtime=5s -count 5 -benchmem -cpuprofile=cpu.out -memprofile=mem.out -trace=trace.out ./package | tee bench.txt
go tool pprof -http :8080 cpu.out
go tool pprof -http :8081 mem.out
go tool trace trace.out

go tool pprof $FILENAME.test cpu.out
# (pprof) list <func name>

# go get -u golang.org/x/perf/cmd/benchstat
benchstat bench.txt
rm cpu.out mem.out trace.out *.test
 */