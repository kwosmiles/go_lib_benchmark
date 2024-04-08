package main

import (
	"golibbenchmark/benchparam"
	"math/rand"
	"runtime"
	"sync/atomic"
	"testing"
)

const (
	NUM_KEYS      = benchparam.BENCH_COUNT
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

// https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
func getRand(n int) string {
	b := make([]byte, n)
	// A rand.Int63() generates 63 random bits, enough for letterIdxMax letters!
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

// benchmarkMapSets performs sets concurrently
func benchmarkMapSet(b *testing.B, m ConcurrentMap) {

	// Prefill the sets we will perform on the
	// map
	setKeys := make([]string, NUM_KEYS)
	for i := 0; i < NUM_KEYS; i++ {
		setKeys[i] = getRand(15)
	}

	setValues := make([]string, NUM_KEYS)
	for i := 0; i < NUM_KEYS; i++ {
		setValues[i] = getRand(15)
	}

	ptr := uint32(0)

	runtime.GC()
	b.ResetTimer()

	b.Run("Set", func(sb *testing.B) {
		sb.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				atomic.AddUint32(&ptr, 1)
				m.Set(setKeys[ptr%NUM_KEYS], setValues[ptr%NUM_KEYS])
			}
		})
	})
}

// benchmarkMapGet performs gets concurrently
func benchmarkMapGet(b *testing.B, m ConcurrentMap) {

	// Prefill the gets we will perform on the
	// map
	getKeys := make([]string, NUM_KEYS)
	for i := 0; i < NUM_KEYS; i++ {
		getKeys[i] = getRand(15)
	}

	ptr := uint32(0)

	runtime.GC()
	b.ResetTimer()

	b.Run("Get", func(sb *testing.B) {
		sb.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				atomic.AddUint32(&ptr, 1)
				_, _ = m.Get(getKeys[ptr%NUM_KEYS])
			}
		})
	})
}

// benchmarkMapMix tests alternating sets and gets in parallel.
// This should simulate a balanced set/get load
func benchmarkMapMix(b *testing.B, m ConcurrentMap) {

	// Prefill the sets we will perform on the
	// map
	setKeys := make([]string, NUM_KEYS)
	for i := 0; i < NUM_KEYS; i++ {
		setKeys[i] = getRand(15)
	}

	setValues := make([]string, NUM_KEYS)
	for i := 0; i < NUM_KEYS; i++ {
		setValues[i] = getRand(15)
	}

	// Prefill the gets we will perform on the
	// map
	getKeys := make([]string, NUM_KEYS)
	for i := 0; i < NUM_KEYS; i++ {
		getKeys[i] = getRand(15)
	}

	setPtr := uint32(0)
	getPtr := uint32(0)

	runtime.GC()
	b.ResetTimer()

	b.Run("Mix", func(sb *testing.B) {
		sb.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set := false
				if set {
					atomic.AddUint32(&setPtr, 1)
					m.Set(setKeys[setPtr%NUM_KEYS], setValues[setPtr%NUM_KEYS])
				} else {
					atomic.AddUint32(&getPtr, 1)
					_, _ = m.Get(getKeys[getPtr%NUM_KEYS])
				}
				set = !set
			}
		})
	})
}

func benchmark(b *testing.B, m ConcurrentMap) {
	b.ResetTimer()
	benchmarkMapGet(b, m)
	benchmarkMapSet(b, m)
	benchmarkMapMix(b, m)
}

func BenchmarkOrignMap_WithRWMutex(b *testing.B) {
	benchmark(b, NewOriginRWMutexMap())
}

func BenchmarkOrignMap_WithMutex(b *testing.B) {
	benchmark(b, NewOriginMutexMap())
}

func BenchmarkSyncMap(b *testing.B) {
	benchmark(b, NewSyncMap())
}

func BenchmarkOrcaman_con_map(b *testing.B) {
	benchmark(b, NewOrcamanLibrary())
}

func BenchmarkFanLiao_con_map(b *testing.B) {
	benchmark(b, NewFanliaoLibrary())
}

func BenchmarkTidwall_shardmap(b *testing.B) {
	benchmark(b, NewTidwallLibrary())
}

/* init cost for more than 10 minutes
func BenchmarkCornelk_hashmap(b *testing.B) {
	benchmark(b, NewCornelkLibrary())
}
*/

func BenchmarkDustinxie_lockfree(b *testing.B) {
	benchmark(b, NewDustinxieLibrary())
}

func BenchmarkAlphadose_haxmap(b *testing.B) {
	benchmark(b, NewAlphadoseLibrary())
}

func BenchmarkEasierway_con_map(b *testing.B) {
	benchmark(b, NewEasierwayLibrary())
}

func BenchmarkMhmtszr_con_s_map(b *testing.B) {
	benchmark(b, NewMhmtszrLibrary())
}

func BenchmarkZhangyunhao_skipmap(b *testing.B) {
	benchmark(b, NewMhmtszrLibrary())
}

func BenchmarkLrita_cmap(b *testing.B) {
	benchmark(b, NewLritaLibrary())
}

func BenchmarkOneOfOne_cmap(b *testing.B) {
	benchmark(b, NewOneOfOneLibrary())
}

func BenchmarkAntoniomo_shardedmap(b *testing.B) {
	benchmark(b, NewAntoniomoLibrary())
}

func BenchmarkHfdpx_chashmap(b *testing.B) {
	benchmark(b, NewHfdpxLibrary())
}

func BenchmarkMojinfu_cmap(b *testing.B) {
	benchmark(b, NewMojinfuLibrary())
}

func BenchmarkXiao7737_cmap(b *testing.B) {
	benchmark(b, NewXiao7737Library())
}

func BenchmarkRfyiamcool_cmap(b *testing.B) {
	benchmark(b, NewRfyiamcoolLibrary())
}

/* can not set interface{}
func BenchmarkKazu_skiplistmap(b *testing.B) {
	benchmark(b, NewKazuLibrary())
}
*/
