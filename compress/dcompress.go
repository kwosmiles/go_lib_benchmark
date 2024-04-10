package compresstest

import (
	"bytes"
	"compress/zlib"
	"io/ioutil"

	dzstd "github.com/DataDog/zstd"
	kzstd "github.com/klauspost/compress/zstd"
	vzstd "github.com/valyala/gozstd"
)

var (
	decoder, _ = kzstd.NewReader(nil, kzstd.WithDecoderConcurrency(0))
	encoder, _ = kzstd.NewWriter(nil)
)

type compressLib interface {
	Name() string
	Compress([]byte) ([]byte, error)
	Decompress([]byte) ([]byte, error)
}

func zlibCompress(data []byte) ([]byte, error) {
	var b bytes.Buffer
	// w, err := zlib.NewWriterLevel(&b, zlib.BestCompression)
	w := zlib.NewWriter(&b)
	_, err := w.Write(data)
	if err != nil {
		return nil, err
	}
	if err := w.Close(); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func zlibDecompress(data []byte) ([]byte, error) {
	r, err := zlib.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	defer r.Close()
	return ioutil.ReadAll(r)
}

type ZlibLibrary struct {
	name string
}

func (m *ZlibLibrary) Name() string {
	return "origin zlib"
}
func (m *ZlibLibrary) Compress(data []byte) ([]byte, error) {
	return zlibCompress(data)
}

func (m *ZlibLibrary) Decompress(data []byte) ([]byte, error) {
	return zlibDecompress(data)
}

func NewZlibLibrary() *ZlibLibrary {
	return &ZlibLibrary{}
}

type KlauspostLibrary struct {
	name string
}

func (m *KlauspostLibrary) Name() string {
	return "klauspost zstd"
}

func (m *KlauspostLibrary) Compress(data []byte) ([]byte, error) {
	return encoder.EncodeAll(data, make([]byte, 0, len(data))), nil
}

func (m *KlauspostLibrary) Decompress(data []byte) ([]byte, error) {
	return decoder.DecodeAll(data, nil)
}

func NewKlauspostLibrary() *KlauspostLibrary {
	return &KlauspostLibrary{}
}

type ValyalaLibrary struct {
	name string
}

func (m *ValyalaLibrary) Name() string {
	return "valyala zstd"
}

func (m *ValyalaLibrary) Compress(data []byte) ([]byte, error) {
	return vzstd.Compress(nil, data), nil
}

func (m *ValyalaLibrary) Decompress(data []byte) ([]byte, error) {
	return vzstd.Decompress(nil, data)
}

func NewValyalaLibrary() *ValyalaLibrary {
	return &ValyalaLibrary{}
}

type DataDogLibrary struct {
	name string
}

func (m *DataDogLibrary) Name() string {
	return "DataDog zstd"
}

func (m *DataDogLibrary) Compress(data []byte) ([]byte, error) {
	return dzstd.Compress(nil, data)
}

func (m *DataDogLibrary) Decompress(data []byte) ([]byte, error) {
	return dzstd.Decompress(nil, data)
}

func NewDataDogLibrary() *DataDogLibrary {
	return &DataDogLibrary{}
}
