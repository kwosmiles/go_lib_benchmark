package jsontest

import (
	"golibbenchmark/benchparam"
	"testing"

	"github.com/mailru/easyjson"
)

var (
	smallStruct  SmallStruct
	mediumStruct MediumStruct
	bigStruct    BigStruct
)

// benchmarkMarshal small struct
func benchmarkMarshalsmall(b *testing.B, m JsonLib) {
	b.ResetTimer()
	b.Run("Marshal-small", func(tb *testing.B) {
		for i := 0; i < tb.N; i++ {
			m.Marshal(smallStruct)
		}
	})
}

// benchmarkMarshal small struct
func benchmarkMarshalmedium(b *testing.B, m JsonLib) {
	b.ResetTimer()
	b.Run("Marshal-medium", func(tb *testing.B) {
		for i := 0; i < tb.N; i++ {
			m.Marshal(mediumStruct)
		}
	})
}

// benchmarkMarshal small struct
func benchmarkMarshalbig(b *testing.B, m JsonLib) {
	b.ResetTimer()
	b.Run("Marshal-big", func(tb *testing.B) {
		for i := 0; i < tb.N; i++ {
			m.Marshal(bigStruct)
		}
	})
}

// benchmarkUnmarshal
func benchmarkUnMarshalsmall(b *testing.B, m JsonLib) {
	b.ResetTimer()

	b.Run("UnMarshal-small", func(tb *testing.B) {
		for i := 0; i < tb.N; i++ {
			m.UnMarshal([]byte(benchparam.Smallbytes), &smallStruct)
		}
	})
}

func benchmarkUnMarshalmedium(b *testing.B, m JsonLib) {
	b.ResetTimer()

	b.Run("UnMarshal-medium", func(tb *testing.B) {
		for i := 0; i < tb.N; i++ {
			m.UnMarshal([]byte(benchparam.Mediumbytes), &mediumStruct)
		}
	})
}

func benchmarkUnMarshalbig(b *testing.B, m JsonLib) {
	b.ResetTimer()

	b.Run("UnMarshal-big", func(tb *testing.B) {
		for i := 0; i < tb.N; i++ {
			m.UnMarshal([]byte(benchparam.Bigbytes), &bigStruct)
		}
	})
}

func benchmark(b *testing.B, m JsonLib) {
	b.ResetTimer()
	benchmarkUnMarshalsmall(b, m)
	benchmarkUnMarshalmedium(b, m)
	benchmarkUnMarshalbig(b, m)
	benchmarkMarshalsmall(b, m)
	benchmarkMarshalmedium(b, m)
	benchmarkMarshalbig(b, m)
}

func BenchmarkOrigin_json(b *testing.B) {
	benchmark(b, NewOriginLibrary())
}

func BenchmarkBytedance_sonic(b *testing.B) {
	benchmark(b, NewSonicLibrary())
}

func BenchmarkJson_iter(b *testing.B) {
	benchmark(b, NewJsoniterLibrary())
}

func BenchmarkPquerna_ffjson(b *testing.B) {
	benchmark(b, NewPquernaLibrary())
}

// func BenchmarkCosmWasm_tinyjson(b *testing.B) {
// 	benchmark(b, NewCosmWasmLibrary())
// }

//for easyjson

func BenchmarkMailru_easyjson(b *testing.B) {
	// benchmark(b, NewMailruLibrary())
	b.ResetTimer()
	benchmarkUnMarshal_M_e_small(b)
	benchmarkUnMarshal_M_e_medium(b)
	benchmarkUnMarshal_M_e_big(b)
	benchmarkMarshal_M_e_small(b)
	benchmarkMarshal_M_e_medium(b)
	benchmarkMarshal_M_e_big(b)
}

// benchmarkMarshal small struct
func benchmarkMarshal_M_e_small(b *testing.B) {
	b.ResetTimer()
	b.Run("Marshal-small", func(tb *testing.B) {
		for i := 0; i < tb.N; i++ {
			easyjson.Marshal(smallStruct)
		}
	})
}

// benchmarkMarshal small struct
func benchmarkMarshal_M_e_medium(b *testing.B) {
	b.ResetTimer()
	b.Run("Marshal-medium", func(tb *testing.B) {
		for i := 0; i < tb.N; i++ {
			easyjson.Marshal(mediumStruct)
		}
	})
}

// benchmarkMarshal small struct
func benchmarkMarshal_M_e_big(b *testing.B) {
	b.ResetTimer()
	b.Run("Marshal-big", func(tb *testing.B) {
		for i := 0; i < tb.N; i++ {
			easyjson.Marshal(bigStruct)
		}
	})
}

// benchmarkUnmarshal
func benchmarkUnMarshal_M_e_small(b *testing.B) {
	b.ResetTimer()

	b.Run("UnMarshal-small", func(tb *testing.B) {
		for i := 0; i < tb.N; i++ {
			easyjson.Unmarshal([]byte(benchparam.Smallbytes), &smallStruct)
		}
	})
}

func benchmarkUnMarshal_M_e_medium(b *testing.B) {
	b.ResetTimer()

	b.Run("UnMarshal-medium", func(tb *testing.B) {
		for i := 0; i < tb.N; i++ {
			easyjson.Unmarshal([]byte(benchparam.Mediumbytes), &mediumStruct)
		}
	})
}

func benchmarkUnMarshal_M_e_big(b *testing.B) {
	b.ResetTimer()

	b.Run("UnMarshal-big", func(tb *testing.B) {
		for i := 0; i < tb.N; i++ {
			easyjson.Unmarshal([]byte(benchparam.Bigbytes), &bigStruct)
		}
	})
}
