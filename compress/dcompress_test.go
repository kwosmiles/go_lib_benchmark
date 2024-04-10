package compresstest

import (
	"fmt"
	"golibbenchmark/benchparam"
	"runtime"
	"testing"
)

var (
	scompress []byte
	mcompress []byte
	bcompress []byte
)

func benchmarkCompresssmall(b *testing.B, m compressLib) {
	runtime.GC()
	b.ResetTimer()

	b.Run("Compress-small", func(tb *testing.B) {
		for i := 0; i < tb.N; i++ {
			scompress, _ = m.Compress([]byte(benchparam.Smallbytes))
		}
	})
}
func benchmarkCompressmedium(b *testing.B, m compressLib) {
	runtime.GC()
	b.ResetTimer()

	b.Run("Compress-medium", func(tb *testing.B) {
		for i := 0; i < tb.N; i++ {
			mcompress, _ = m.Compress([]byte(benchparam.Mediumbytes))
		}
	})
}
func benchmarkCompressbig(b *testing.B, m compressLib) {
	runtime.GC()
	b.ResetTimer()

	b.Run("Compress-big", func(tb *testing.B) {
		for i := 0; i < tb.N; i++ {
			bcompress, _ = m.Compress([]byte(benchparam.Bigbytes))
		}
	})
}

func benchmarkDecompresssmall(b *testing.B, m compressLib) {
	runtime.GC()
	b.ResetTimer()

	b.Run("Decompress-small", func(tb *testing.B) {
		for i := 0; i < tb.N; i++ {
			m.Decompress(scompress)
		}
	})
}
func benchmarkDecompressmedium(b *testing.B, m compressLib) {
	runtime.GC()
	b.ResetTimer()

	b.Run("Decompress-medium", func(tb *testing.B) {
		for i := 0; i < tb.N; i++ {
			m.Decompress(mcompress)
		}
	})
}
func benchmarkDecompressbig(b *testing.B, m compressLib) {
	runtime.GC()
	b.ResetTimer()

	b.Run("Decompress-big", func(tb *testing.B) {
		for i := 0; i < tb.N; i++ {
			m.Decompress(bcompress)
		}
	})
}
func benchmark(b *testing.B, m compressLib) {
	b.ResetTimer()
	sbytes, _ := m.Compress([]byte(benchparam.Smallbytes))
	mbytes, _ := m.Compress([]byte(benchparam.Mediumbytes))
	bbytes, _ := m.Compress([]byte(benchparam.Bigbytes))
	fmt.Println("[", m.Name(), "] compress rate small:", float64(len(sbytes))/float64(len([]byte(benchparam.Smallbytes))), ";medium:", float64(len(mbytes))/float64(len([]byte(benchparam.Mediumbytes))), ";big:", float64(len(bbytes))/float64(len([]byte(benchparam.Bigbytes))))
	benchmarkCompresssmall(b, m)
	benchmarkCompressmedium(b, m)
	benchmarkCompressbig(b, m)
	benchmarkDecompresssmall(b, m)
	benchmarkDecompressmedium(b, m)
	benchmarkDecompressbig(b, m)
}

func BenchmarkOrigin_zlib(b *testing.B) {
	benchmark(b, NewZlibLibrary())
}

func BenchmarkKlauspost_zstd(b *testing.B) {
	benchmark(b, NewKlauspostLibrary())
}

func BenchmarkValyala_zstd(b *testing.B) {
	benchmark(b, NewValyalaLibrary())
}

func BenchmarkDataDog_zstd(b *testing.B) {
	benchmark(b, NewDataDogLibrary())
}
