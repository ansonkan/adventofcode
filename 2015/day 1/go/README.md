## Benchmarks

```console
goos: darwin
goarch: arm64
pkg: aoc/2015day1
Benchmark_ReadFile-8               64345             16614 ns/op           16712 B/op          6 allocs/op
Benchmark_ReadFile_large-8           268           3906637 ns/op         5292404 B/op          6 allocs/op
Benchmark_Seek-8                   18406             64105 ns/op             120 B/op          3 allocs/op
Benchmark_Seek_large-8                54          21681018 ns/op             128 B/op          3 allocs/op
PASS
ok      aoc/2015day1    5.890s
```