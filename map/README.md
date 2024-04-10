# benchmark for map

collect from https://github.com/s0lesurviv0r/go-concurrent-map-bench
add more concurrent lib with 10 more stars

## benchmark libs:
go origin map with RWMutex
go origin map with Mutex
github.com/OneOfOne/cmap
github.com/alphadose/haxmap
github.com/cornelk/hashmap --- init cost for a long time
github.com/dustinxie/lockfree
github.com/easierway/concurrent_map
github.com/fanliao/go-concurrentMap
github.com/lrita/cmap"
github.com/mhmtszr/concurrent-swiss-map
github.com/orcaman/concurrent-map
github.com/tidwall/shardmap
github.com/zhangyunhao116/skipmap
github.com/antoniomo/shardedmap
github.com/hfdpx/ConcurrentHashMap
github.com/mojinfu/cmap
github.com/xiao7737/concurrentMap
github.com/rfyiamcool/ccmap
github.com/kazu/skiplistmap -- can not set interface

## running
```
go test --bench .
```
## result
```
goos: linux
goarch: amd64
pkg: golibbenchmark/map
cpu: Intel(R) Core(TM) i5-4460  CPU @ 3.20GHz
BenchmarkOrignMap_WithRWMutex/Get-4             20059665                59.84 ns/op
BenchmarkOrignMap_WithRWMutex/Set-4              3785324               321.5 ns/op
BenchmarkOrignMap_WithRWMutex/Mix-4             15720046                76.06 ns/op
BenchmarkOrignMap_WithMutex/Get-4                8837232               134.4 ns/op
BenchmarkOrignMap_WithMutex/Set-4                4177994               287.8 ns/op
BenchmarkOrignMap_WithMutex/Mix-4                5083863               235.4 ns/op
BenchmarkSyncMap/Get-4                          36386488                32.80 ns/op
BenchmarkSyncMap/Set-4                           1000000              1107 ns/op
BenchmarkSyncMap/Mix-4                          19418042                61.58 ns/op
BenchmarkOrcaman_con_map/Get-4                  19204095                62.48 ns/op
BenchmarkOrcaman_con_map/Set-4                   7327550               156.9 ns/op
BenchmarkOrcaman_con_map/Mix-4                  15217290                78.94 ns/op
BenchmarkFanLiao_con_map/Get-4                  25592751                50.88 ns/op
BenchmarkFanLiao_con_map/Set-4                   4195335               266.8 ns/op
BenchmarkFanLiao_con_map/Mix-4                  13834550               103.9 ns/op
BenchmarkTidwall_shardmap/Get-4                 23897686                50.00 ns/op
BenchmarkTidwall_shardmap/Set-4                  7759197               133.4 ns/op
BenchmarkTidwall_shardmap/Mix-4                 17573176                67.90 ns/op
BenchmarkDustinxie_lockfree/Get-4               11798685               104.3 ns/op
BenchmarkDustinxie_lockfree/Set-4                2615641               429.3 ns/op
BenchmarkDustinxie_lockfree/Mix-4                2798797               428.2 ns/op
BenchmarkAlphadose_haxmap/Get-4                 32679894                36.49 ns/op
BenchmarkAlphadose_haxmap/Set-4                  7161895               163.7 ns/op
BenchmarkAlphadose_haxmap/Mix-4                 11134938               107.8 ns/op
BenchmarkEasierway_con_map/Get-4                14996077                78.92 ns/op
BenchmarkEasierway_con_map/Set-4                 4855020               233.3 ns/op
BenchmarkEasierway_con_map/Mix-4                12445772               107.1 ns/op
BenchmarkMhmtszr_con_s_map/Get-4                22738508                52.95 ns/op
BenchmarkMhmtszr_con_s_map/Set-4                 6642570               183.1 ns/op
BenchmarkMhmtszr_con_s_map/Mix-4                11469250               104.4 ns/op
BenchmarkZhangyunhao_skipmap/Get-4              22841899                52.41 ns/op
BenchmarkZhangyunhao_skipmap/Set-4               6789756               182.9 ns/op
BenchmarkZhangyunhao_skipmap/Mix-4              11440526               104.6 ns/op
BenchmarkLrita_cmap/Get-4                       19259799                61.99 ns/op
BenchmarkLrita_cmap/Set-4                        5249988               228.6 ns/op
BenchmarkLrita_cmap/Mix-4                       12370696                96.63 ns/op
BenchmarkOneOfOne_cmap/Get-4                    21610464                58.22 ns/op
BenchmarkOneOfOne_cmap/Set-4                     5253439               198.8 ns/op
BenchmarkOneOfOne_cmap/Mix-4                    14568510                96.09 ns/op
BenchmarkAntoniomo_shardedmap/Get-4             23687881                50.48 ns/op
BenchmarkAntoniomo_shardedmap/Set-4              8199264               141.0 ns/op
BenchmarkAntoniomo_shardedmap/Mix-4             15001922                79.82 ns/op
BenchmarkHfdpx_chashmap/Get-4                   18994599                63.19 ns/op
BenchmarkHfdpx_chashmap/Set-4                    7468166               158.0 ns/op
BenchmarkHfdpx_chashmap/Mix-4                   15352245                78.45 ns/op
BenchmarkMojinfu_cmap/Get-4                     34554555                34.57 ns/op
BenchmarkMojinfu_cmap/Set-4                      2727090               447.9 ns/op
BenchmarkMojinfu_cmap/Mix-4                     19522741                61.56 ns/op
BenchmarkXiao7737_cmap/Get-4                    14984178                78.83 ns/op
BenchmarkXiao7737_cmap/Set-4                     5478997               228.8 ns/op
BenchmarkXiao7737_cmap/Mix-4                    12698864               105.6 ns/op
BenchmarkRfyiamcool_cmap/Get-4                  19595985                61.13 ns/op
BenchmarkRfyiamcool_cmap/Set-4                   8074651               147.2 ns/op
BenchmarkRfyiamcool_cmap/Mix-4                  15659810                76.48 ns/op
PASS
ok      golibbenchmark/map      141.247s
```