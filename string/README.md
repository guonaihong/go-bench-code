golang字符串压测速度

* 写入速度
```
env GOPATH=`pwd` go test -test.bench="Benchmark_Add" -v github.com/guonaihong/test/string
goos: linux
goarch: amd64
pkg: github.com/guonaihong/test/string
Benchmark_Add-4   	  500000	    157961 ns/op
PASS
ok  	github.com/guonaihong/test/string	79.027s



env GOPATH=`pwd` go test -bench "Benchmark_BytesWrite" -v github.com/guonaihong/test/string
goos: linux
goarch: amd64
pkg: github.com/guonaihong/test/string
Benchmark_BytesWrite-4   	100000000	        10.7 ns/op
PASS
ok  	github.com/guonaihong/test/string	1.120s



env GOPATH=`pwd` go test -bench "Benchmark_builder" -v github.com/guonaihong/test/string
goos: linux
goarch: amd64
pkg: github.com/guonaihong/test/string
Benchmark_builder-4   	200000000	         6.48 ns/op
PASS
ok  	github.com/guonaihong/test/string	2.060s
```

* 内存占用
```
env GOPATH=`pwd` go test -bench "Benchmark_builderMem" -benchmem  -v github.com/guonaihong/test/string
goos: linux
goarch: amd64
pkg: github.com/guonaihong/test/string
Benchmark_builderMem-4   	200000000	         9.54 ns/op	      30 B/op	       0 allocs/op
PASS
ok  	github.com/guonaihong/test/string	2.831s


env GOPATH=`pwd` go test -bench "Benchmark_BytesWriteMem" -benchmem  -v github.com/guonaihong/test/string
goos: linux
goarch: amd64
pkg: github.com/guonaihong/test/string
Benchmark_BytesWriteMem-4   	  500000	    104201 ns/op	 1254075 B/op	       1 allocs/op
PASS
ok  	github.com/guonaihong/test/string	52.153s

```
