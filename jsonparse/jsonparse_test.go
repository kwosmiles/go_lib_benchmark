package jsonparsetest

import (
	"golibbenchmark/benchparam"
	"testing"

	"github.com/buger/jsonparser"
	"github.com/tidwall/gjson"
)

func BenchmarkTidwall_gjson(b *testing.B) {
	b.ResetTimer()
	benchmarkGet_gjson_small(b)
	benchmarkGet_gjson_medium(b)
	benchmarkGet_gjson_big(b)
}
func benchmarkGet_gjson_small(b *testing.B) {
	b.ResetTimer()
	b.Run("Get-small", func(tb *testing.B) {
		for i := 0; i < tb.N; i++ {
			gjson.Get(benchparam.Smallbytes, "participant.identity").String()
		}
	})
}
func benchmarkGet_gjson_medium(b *testing.B) {
	b.ResetTimer()
	b.Run("Get-medium", func(tb *testing.B) {
		for i := 0; i < tb.N; i++ {
			gjson.Get(benchparam.Mediumbytes, "person.name.fullName").String()
		}
	})
}

func benchmarkGet_gjson_big(b *testing.B) {
	b.ResetTimer()
	b.Run("Get-big", func(tb *testing.B) {
		for i := 0; i < tb.N; i++ {
			gjson.Get(benchparam.Bigbytes, "person.person.name.givenName").String()
		}
	})
}
func BenchmarkBuger_jsonparser(b *testing.B) {
	b.ResetTimer()
	benchmarkGet_jsonparser_small(b)
	benchmarkGet_jsonparser_medium(b)
	benchmarkGet_jsonparser_big(b)
	benchmarkSet_jsonparser_small(b)
	benchmarkSet_jsonparser_medium(b)
	benchmarkSet_jsonparser_big(b)
}
func benchmarkGet_jsonparser_small(b *testing.B) {
	b.ResetTimer()
	b.Run("Get-small", func(tb *testing.B) {
		for i := 0; i < tb.N; i++ {
			jsonparser.Get([]byte(benchparam.Smallbytes), "participant", "identity")
		}
	})
}
func benchmarkGet_jsonparser_medium(b *testing.B) {
	b.ResetTimer()
	b.Run("Get-medium", func(tb *testing.B) {
		for i := 0; i < tb.N; i++ {
			jsonparser.Get([]byte(benchparam.Mediumbytes), "person", "name", "fullName")
		}
	})
}

func benchmarkGet_jsonparser_big(b *testing.B) {
	b.ResetTimer()
	b.Run("Get-big", func(tb *testing.B) {
		for i := 0; i < tb.N; i++ {
			jsonparser.Get([]byte(benchparam.Bigbytes), "person", "person", "name", "givenName")
		}
	})
}
func benchmarkSet_jsonparser_small(b *testing.B) {
	b.ResetTimer()
	b.Run("Set-small", func(tb *testing.B) {
		for i := 0; i < tb.N; i++ {
			jsonparser.Set([]byte(benchparam.Smallbytes), []byte("tttest"), "participant", "identity")
		}
	})
}
func benchmarkSet_jsonparser_medium(b *testing.B) {
	b.ResetTimer()
	b.Run("Set-medium", func(tb *testing.B) {
		for i := 0; i < tb.N; i++ {
			jsonparser.Set([]byte(benchparam.Mediumbytes), []byte("tttest"), "person", "name", "fullName")
		}
	})
}
func benchmarkSet_jsonparser_big(b *testing.B) {
	b.ResetTimer()
	b.Run("Set-big", func(tb *testing.B) {
		for i := 0; i < tb.N; i++ {
			jsonparser.Set([]byte(benchparam.Bigbytes), []byte("tttest"), "person", "person", "name", "givenName")
		}
	})
}
