golang字符串压测速度
#### 字符串拼接
* 写入速度
```
env GOPATH=`pwd` go test -test.bench="Benchmark_Add" -v .
goos: linux
goarch: amd64
pkg: github.com/guonaihong/test/string
Benchmark_Add-4   	  500000	    157961 ns/op
PASS
ok  	github.com/guonaihong/test/string	79.027s



env GOPATH=`pwd` go test -bench "Benchmark_BytesWrite" -v .
goos: linux
goarch: amd64
pkg: github.com/guonaihong/test/string
Benchmark_BytesWrite-4   	100000000	        10.7 ns/op
PASS
ok  	github.com/guonaihong/test/string	1.120s



env GOPATH=`pwd` go test -bench "Benchmark_builder" -v .
goos: linux
goarch: amd64
pkg: github.com/guonaihong/test/string
Benchmark_builder-4   	200000000	         6.48 ns/op
PASS
ok  	github.com/guonaihong/test/string	2.060s
```

* 内存占用
```
env GOPATH=`pwd` go test -bench "Benchmark_builderMem" -benchmem  -v .
goos: linux
goarch: amd64
pkg: github.com/guonaihong/test/string
Benchmark_builderMem-4   	200000000	         9.54 ns/op	      30 B/op	       0 allocs/op
PASS
ok  	github.com/guonaihong/test/string	2.831s


env GOPATH=`pwd` go test -bench "Benchmark_BytesWriteMem" -benchmem  -v .
goos: linux
goarch: amd64
pkg: github.com/guonaihong/test/string
Benchmark_BytesWriteMem-4   	  500000	    104201 ns/op	 1254075 B/op	       1 allocs/op
PASS
ok  	github.com/guonaihong/test/string	52.153s

```

#### 数字转字符串
```
env GOPATH=`pwd` go test -bench "Benchmark_Sprintf" -benchmem  -v .
goos: linux
goarch: amd64
Benchmark_Sprintf-4   	20000000	        86.5 ns/op	      16 B/op	       2 allocs/op
PASS
ok  	_/home/guo/src/github.com/guonaihong/go-test-code/string	1.821s


env GOPATH=`pwd` go test -bench "Benchmark_Itoa" -benchmem  -v .
goos: linux
goarch: amd64
Benchmark_Itoa-4   	50000000	        29.5 ns/op	       7 B/op	       0 allocs/op
PASS
ok  	_/home/guo/src/github.com/guonaihong/go-test-code/string	1.507s

```
