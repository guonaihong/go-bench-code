golang 打印相关测试
#### 打印
```
go test -test.bench="Benchmark_printf" -benchmem .

goos: linux
goarch: amd64
Benchmark_printf-4   	hello world
 1000000	      1272 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	_/home/guo/src/github.com/guonaihong/go-bench-code/fmt	1.287s
```

```
go test -test.bench="Benchmark_println" -benchmem .

goos: linux
goarch: amd64
Benchmark_println-4   	hello, world
 1000000	      1287 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	_/home/guo/src/github.com/guonaihong/go-bench-code/fmt	1.302s
```
