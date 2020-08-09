package namedpool_test

import (
	"awesomeProject1/app-web/namedpool"
	"sync"
	"testing"
)

type Small struct {
	a int
}
func (p *Small)Reset(){}
func (p *Small)Clear(){}

type Big struct {
	b int
}
func (p *Big)Reset(){}
func (p *Big)Clear(){}


var name string = "pool"
var name1 string = "pool1"

func init()  {
	namedpool.New(name,func() interface{} { return new(Small) },)
	namedpool.New(name1,func() interface{} { return new(Big) },)
}

//go:noinline
func inc(s *Small) { s.a++;_ = s.a }
func inc1(s *Big) { s.b++;_ = s.b }

//func BenchmarkWithoutPool(b *testing.B) {
//	var s *Small
//	var s1 *Big
//	for i := 0; i < b.N; i++ {
//		for j := 0; j < 10000; j++ {
//			s = &Small{ a: 1, }
//			s1 = &Big{ b: 100, }
//			b.StopTimer()
//			inc(s)
//			inc1(s1)
//			b.StartTimer()
//		}
//	}
//}

func BenchmarkWithPool(b *testing.B) {
	var s *Small
	//var s1 *Big
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			s = namedpool.Get(name).(*Small)
			s.a = 1
			//ss1 = namedpool.Get(name).(*Small)
			//ss1.a = 11
			//ss2 = namedpool.Get(name).(*Small)
			//ss2.a = 1111
			//s1 = namedpool.Get(name1).(*Big)
			//s1.b = 100
			//b.StopTimer()
			inc(s)
			//inc1(s1)
			//b.StartTimer()
			namedpool.Put(name,s)
			//namedpool.Put(name,ss1)
			//namedpool.Put(name,ss2)
			//namedpool.Put(name1,s1)
		}
	}
}

var pool = sync.Pool{
	New: func() interface{} { return new(Small) },
}
func BenchmarkWithPool2(b *testing.B) {
	var s *Small
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			s = pool.Get().(*Small)
			s.a = 1
			//b.StopTimer();
			inc(s);
			//b.StartTimer()
			pool.Put(s)
		}
	}
}

//func BenchmarkBuilder(b *testing.B) {
////func main(){
//	for y:=1;y<100;y++{
//		named := "data1"+string(y)
//		New(named,NewData)
//
//		for i:=0;i<1000;i++{
//			var t [100]interface{}
//			for z:=0;z<100;z++{
//				t[z] = Get(named)
//			}
//			for z:=0;z<100;z++{
//				 Put(named,t[z])
//			}
//		}
//	}
//}
////go test -benchmem -run=^$  -bench=. -v -count=1 -memprofile=mem.out

/*
C:\proj\go\src\awesomeProject1\app-web\namedpool>go test -benchmem -run=^$  -bench=. -v -count=1
goos: windows
goarch: amd64
pkg: awesomeProject1/app-web/namedpool
BenchmarkWithoutPool
BenchmarkWithoutPool-12              240           5747085 ns/op          160034 B/op      20000 allocs/op
BenchmarkWithPool
BenchmarkWithPool-12                  99          11632069 ns/op              31 B/op          0 allocs/op
PASS
ok      awesomeProject1/app-web/namedpool       187.894s
 */
// 相同时间内：a * b 是相等
// a. 関数の実行回数、有用な結果が得られるまで実行される
//240

// b. １回の実行にかかった時間
// 少ないほど良い
//5747085 ns/op

// 実行ごとに割り当てられたメモリのサイズ
// 少ないほど良い
//336 B/op

// １回の実行でメモリアロケーションが行われた回数
// 少ないほど良い
//9 allocs/op