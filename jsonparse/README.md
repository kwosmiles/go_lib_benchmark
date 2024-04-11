# benchmark for json get quickly libs
init []byte with length:350bytes,1.5k,30k
then test the get and set of json get quickly libs

it's slow then ffjson to unmarshal json and modifies and marshal again. it is better to use ffjson. and it is easy to replace for the golang's encoding/json.

## benchmark libs:
```
github.com/buger/jsonparser
github.com/tidwall/gjson
```

## running
```
go test --bench . --benchmem
```
## result
```
goos: linux
goarch: amd64
pkg: golibbenchmark/jsonparse
cpu: Intel(R) Core(TM) i5-4460  CPU @ 3.20GHz
BenchmarkTidwall_gjson/Get-small-4               5054497               236.7 ns/op             0 B/op          0 allocs/op
BenchmarkTidwall_gjson/Get-medium-4              6884654               173.9 ns/op             0 B/op          0 allocs/op
BenchmarkTidwall_gjson/Get-big-4                   37614             31867 ns/op               0 B/op          0 allocs/op
BenchmarkBuger_jsonparser/Get-small-4            3332146               358.1 ns/op           352 B/op          1 allocs/op
BenchmarkBuger_jsonparser/Get-medium-4           2430700               495.1 ns/op          1536 B/op          1 allocs/op
BenchmarkBuger_jsonparser/Get-big-4                32210             37434 ns/op           32768 B/op          1 allocs/op
BenchmarkBuger_jsonparser/Set-small-4            2566134               466.3 ns/op           704 B/op          2 allocs/op
BenchmarkBuger_jsonparser/Set-medium-4           1530916               777.5 ns/op          3072 B/op          2 allocs/op
BenchmarkBuger_jsonparser/Set-big-4                27890             43252 ns/op           65536 B/op          2 allocs/op
PASS
ok      golibbenchmark/jsonparse        14.486s
````