内存拷贝
```
go test -test.bench="Benchmark_append" -benchmem .
goos: linux
goarch: amd64
Benchmark_append-4   	20000000	        67.5 ns/op	      48 B/op	       1 allocs/op
PASS
ok  	_/home/guo/src/github.com/guonaihong/go-bench-code/copy	1.424s

go test -test.bench="Benchmark_make_copy" -benchmem .
goos: linux
goarch: amd64
Benchmark_make_copy-4   	20000000	        59.0 ns/op	      48 B/op	       1 allocs/op
PASS
ok  	_/home/guo/src/github.com/guonaihong/go-bench-code/copy	1.245s

go test -test.bench="Benchmark_make_append" -benchmem .
goos: linux
goarch: amd64
Benchmark_make_append-4   	20000000	        59.4 ns/op	      48 B/op	       1 allocs/op
PASS
ok  	_/home/guo/src/github.com/guonaihong/go-bench-code/copy	1.254s
```
