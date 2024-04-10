# benchmark for compress/decompress libs
init []byte with length:350bytes,1.5k,30k
then test compress and decompress of zlib/zstd libs

## benchmark libs:
```orggin zlib
github.com/DataDog/zstd
github.com/klauspost/compress/zstd
github.com/valyala/gozstd
```
## running 
```
go test --bench . --benchmem
```

## result
```
[ origin zlib ] compress rate small: 0.78 ;medium: 0.444 ;big: 0.1719
goos: linux
goarch: amd64
pkg: golibbenchmark/compress
cpu: Intel(R) Core(TM) i5-4460  CPU @ 3.20GHz
BenchmarkOrigin_zlib/Compress-small-4               6913            146784 ns/op          815046 B/op         24 allocs/op
BenchmarkOrigin_zlib/Compress-medium-4              6234            175498 ns/op          817253 B/op         25 allocs/op
BenchmarkOrigin_zlib/Compress-big-4                 1948            609578 ns/op          862821 B/op         28 allocs/op
BenchmarkOrigin_zlib/Decompress-small-4            98342             12245 ns/op           41124 B/op          8 allocs/op
BenchmarkOrigin_zlib/Decompress-medium-4           56954             21005 ns/op           45648 B/op         17 allocs/op
BenchmarkOrigin_zlib/Decompress-big-4               7638            146724 ns/op          195461 B/op         33 allocs/op
[ klauspost zstd ] compress rate small: 0.7914285714285715 ;medium: 0.472 ;big: 0.17996666666666666
BenchmarkKlauspost_zstd/Compress-small-4          130446              8806 ns/op             704 B/op          2 allocs/op
BenchmarkKlauspost_zstd/Compress-medium-4          57070             20787 ns/op            3072 B/op          2 allocs/op
BenchmarkKlauspost_zstd/Compress-big-4              6622            156350 ns/op           65536 B/op          2 allocs/op
BenchmarkKlauspost_zstd/Decompress-small-4        372393              3097 ns/op             384 B/op          1 allocs/op
BenchmarkKlauspost_zstd/Decompress-medium-4       197110              5827 ns/op            1538 B/op          1 allocs/op
BenchmarkKlauspost_zstd/Decompress-big-4           32840             35947 ns/op           32792 B/op          1 allocs/op
[ valyala zstd ] compress rate small: 0.8028571428571428 ;medium: 0.472 ;big: 0.17693333333333333
BenchmarkValyala_zstd/Compress-small-4            191348              5881 ns/op             768 B/op          2 allocs/op
BenchmarkValyala_zstd/Compress-medium-4           104661             11152 ns/op            3328 B/op          2 allocs/op
BenchmarkValyala_zstd/Compress-big-4               12985             92527 ns/op           70917 B/op          3 allocs/op
BenchmarkValyala_zstd/Decompress-small-4          463382              2492 ns/op             352 B/op          1 allocs/op
BenchmarkValyala_zstd/Decompress-medium-4         230900              4958 ns/op            1536 B/op          1 allocs/op
BenchmarkValyala_zstd/Decompress-big-4             51158             23365 ns/op           32770 B/op          1 allocs/op
[ DataDog zstd ] compress rate small: 0.8057142857142857 ;medium: 0.4593333333333333 ;big: 0.1729
BenchmarkDataDog_zstd/Compress-small-4            111409             10515 ns/op             816 B/op          4 allocs/op
BenchmarkDataDog_zstd/Compress-medium-4            53156             22107 ns/op            3376 B/op          4 allocs/op
BenchmarkDataDog_zstd/Compress-big-4                4632            238942 ns/op           65584 B/op          4 allocs/op
BenchmarkDataDog_zstd/Decompress-small-4          396765              2874 ns/op             424 B/op          4 allocs/op
BenchmarkDataDog_zstd/Decompress-medium-4         256328              4441 ns/op            1608 B/op          4 allocs/op
BenchmarkDataDog_zstd/Decompress-big-4             47806             24963 ns/op           32840 B/op          4 allocs/op
PASS
ok      golibbenchmark/compress 31.005s
```
