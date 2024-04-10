# benchmark for json libs
init []byte with length:350bytes,1.5k,30k
then test the marshal and unmarshal of json libs

## benchmark libs:
```github.com/mailru/easyjson
github.com/bytedance/sonic
github.com/json-iterator/go
github.com/pquerna/ffjson/ffjson
```
## compare result:
1st: Mailru_easyjson
2st: Pquerna_ffjson

## running
```
go test --bench . --benchmem
```
## result
```goos: linux
goarch: amd64
pkg: golibbenchmark/json
cpu: Intel(R) Core(TM) i5-4460  CPU @ 3.20GHz
BenchmarkOrigin_json/UnMarshal-small-4            344979              3362 ns/op             568 B/op          6 allocs/op
BenchmarkOrigin_json/UnMarshal-medium-4            89912             13163 ns/op            1816 B/op          7 allocs/op
BenchmarkOrigin_json/UnMarshal-big-4                4561            255778 ns/op           33048 B/op          7 allocs/op
BenchmarkOrigin_json/Marshal-small-4              805483              1422 ns/op             296 B/op          5 allocs/op
BenchmarkOrigin_json/Marshal-medium-4             716940              1558 ns/op             872 B/op          5 allocs/op
BenchmarkOrigin_json/Marshal-big-4                698468              1620 ns/op            1256 B/op          5 allocs/op
BenchmarkBytedance_sonic/UnMarshal-small-4       1000000              1024 ns/op             761 B/op          4 allocs/op
BenchmarkBytedance_sonic/UnMarshal-medium-4       346796              3343 ns/op            3148 B/op          4 allocs/op
BenchmarkBytedance_sonic/UnMarshal-big-4           19766             60791 ns/op           66304 B/op          4 allocs/op
BenchmarkBytedance_sonic/Marshal-small-4          251118              4663 ns/op             346 B/op          8 allocs/op
BenchmarkBytedance_sonic/Marshal-medium-4         240768              4922 ns/op             944 B/op          8 allocs/op
BenchmarkBytedance_sonic/Marshal-big-4            235418              5060 ns/op            1345 B/op          8 allocs/op
BenchmarkJson_iter/UnMarshal-small-4              748641              1477 ns/op             872 B/op         21 allocs/op
BenchmarkJson_iter/UnMarshal-medium-4             226572              5238 ns/op            3576 B/op         69 allocs/op
BenchmarkJson_iter/UnMarshal-big-4                 10000            104300 ns/op           78855 B/op       1236 allocs/op
BenchmarkJson_iter/Marshal-small-4               6252986               191.1 ns/op           192 B/op          1 allocs/op
BenchmarkJson_iter/Marshal-medium-4              3607518               333.4 ns/op           768 B/op          1 allocs/op
BenchmarkJson_iter/Marshal-big-4                 2900989               412.7 ns/op          1152 B/op          1 allocs/op
BenchmarkPquerna_ffjson/UnMarshal-small-4        9915133               119.6 ns/op           352 B/op          1 allocs/op
BenchmarkPquerna_ffjson/UnMarshal-medium-4       2548074               463.3 ns/op          1536 B/op          1 allocs/op
BenchmarkPquerna_ffjson/UnMarshal-big-4           148297              7967 ns/op           32768 B/op          1 allocs/op
BenchmarkPquerna_ffjson/Marshal-small-4         11102060               107.8 ns/op           192 B/op          1 allocs/op
BenchmarkPquerna_ffjson/Marshal-medium-4         4492784               260.9 ns/op           768 B/op          1 allocs/op
BenchmarkPquerna_ffjson/Marshal-big-4            3576588               338.6 ns/op          1152 B/op          1 allocs/op
BenchmarkMailru_easyjson/UnMarshal-small-4      72472515                15.98 ns/op            0 B/op          0 allocs/op
BenchmarkMailru_easyjson/UnMarshal-medium-4     25019618                47.89 ns/op            0 B/op          0 allocs/op
BenchmarkMailru_easyjson/UnMarshal-big-4          946107              1242 ns/op               0 B/op          0 allocs/op
BenchmarkMailru_easyjson/Marshal-small-4         7955756               149.5 ns/op           304 B/op          2 allocs/op
BenchmarkMailru_easyjson/Marshal-medium-4        4057591               296.9 ns/op           880 B/op          2 allocs/op
BenchmarkMailru_easyjson/Marshal-big-4           3113920               382.2 ns/op          1264 B/op          2 allocs/op
PASS
ok      golibbenchmark/json     39.502s
````