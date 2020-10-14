#go version go1.13.4 darwin/amd64
goos: darwin
goarch: amd64
BenchmarkLoop1-8   	 1654904	       739 ns/op	      72 B/op	       9 allocs/op
BenchmarkLoop2-8   	 1692212	       709 ns/op	      64 B/op	       8 allocs/op
BenchmarkLoop5-8   	 1793840	       691 ns/op	      40 B/op	       5 allocs/op
PASS
ok  	_/Users/liambennett/code/go-perfexperiments	5.857s

#go version go1.15.2 darwin/amd64
goos: darwin
goarch: amd64
BenchmarkLoop1-8   	 2117413	       571 ns/op	       0 B/op	       0 allocs/op
BenchmarkLoop2-8   	 2235693	       567 ns/op	       0 B/op	       0 allocs/op
BenchmarkLoop5-8   	 2107467	       569 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	_/Users/liambennett/code/go-perfexperiments	5.532s