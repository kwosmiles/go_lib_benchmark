package maptest

import (
	"sync"
	"unsafe"

	ooocmap "github.com/OneOfOne/cmap"
	"github.com/alphadose/haxmap"
	"github.com/antoniomo/shardedmap"
	"github.com/cornelk/hashmap"
	"github.com/dustinxie/lockfree"
	ecmap "github.com/easierway/concurrent_map"
	fcmap "github.com/fanliao/go-concurrentMap"
	hcmap "github.com/hfdpx/ConcurrentHashMap"
	"github.com/kazu/skiplistmap"
	lcmap "github.com/lrita/cmap"
	csmap "github.com/mhmtszr/concurrent-swiss-map"
	mcmap "github.com/mojinfu/cmap"
	ocmap "github.com/orcaman/concurrent-map"
	rcmap "github.com/rfyiamcool/ccmap"
	"github.com/tidwall/shardmap"
	xcmap "github.com/xiao7737/concurrentMap"
	"github.com/zhangyunhao116/skipmap"
)

// ConcurrentMap is the interface all of our
// concurrently accessible maps should conform to
type ConcurrentMap interface {
	Get(string) (interface{}, bool)
	Set(string, interface{})
}

type OriginMutexMap struct {
	sync.Mutex
	m map[string]interface{}
}

func (m *OriginMutexMap) Get(key string) (interface{}, bool) {
	m.Lock()
	val, exists := m.m[key]
	m.Unlock()
	return val, exists
}

func (m *OriginMutexMap) Set(key string, value interface{}) {
	m.Lock()
	m.m[key] = value
	m.Unlock()
}

func NewOriginMutexMap() *OriginMutexMap {
	return &OriginMutexMap{
		m: make(map[string]interface{}),
	}
}

type OriginRWMutexMap struct {
	sync.RWMutex
	m map[string]interface{}
}

func (m *OriginRWMutexMap) Get(key string) (interface{}, bool) {
	m.RLock()
	val, exists := m.m[key]
	m.RUnlock()
	return val, exists
}

func (m *OriginRWMutexMap) Set(key string, value interface{}) {
	m.Lock()
	m.m[key] = value
	m.Unlock()
}

func NewOriginRWMutexMap() *OriginRWMutexMap {
	return &OriginRWMutexMap{
		m: make(map[string]interface{}),
	}
}

type SyncMap struct {
	sync.Map
}

func (m *SyncMap) Get(key string) (interface{}, bool) {
	return m.Load(key)
}

func (m *SyncMap) Set(key string, value interface{}) {
	m.Store(key, value)
}

func NewSyncMap() *SyncMap {
	return &SyncMap{}
}

type OrcamanLibrary struct {
	internal ocmap.ConcurrentMap
}

func (m *OrcamanLibrary) Get(key string) (interface{}, bool) {
	return m.internal.Get(key)
}

func (m *OrcamanLibrary) Set(key string, value interface{}) {
	m.internal.Set(key, value)
}

func NewOrcamanLibrary() *OrcamanLibrary {
	return &OrcamanLibrary{
		internal: ocmap.New(),
	}
}

type FanliaoLibrary struct {
	internal *fcmap.ConcurrentMap
}

func (m *FanliaoLibrary) Get(key string) (interface{}, bool) {
	val, err := m.internal.Get(key)
	if err != nil || val == nil {
		return "", false
	}
	return val, true
}

func (m *FanliaoLibrary) Set(key string, value interface{}) {
	m.internal.Put(key, value)
}

func NewFanliaoLibrary() *FanliaoLibrary {
	return &FanliaoLibrary{
		internal: fcmap.NewConcurrentMap(),
	}
}

type TidwallLibrary struct {
	internal *shardmap.Map
}

func (m *TidwallLibrary) Get(key string) (interface{}, bool) {
	return m.internal.Get(key)
}

func (m *TidwallLibrary) Set(key string, value interface{}) {
	m.internal.Set(key, value)
}

func NewTidwallLibrary() *TidwallLibrary {
	return &TidwallLibrary{
		internal: &shardmap.Map{},
	}
}

type CornelkLibrary struct {
	internal *hashmap.Map[string, interface{}]
}

func (m *CornelkLibrary) Get(key string) (interface{}, bool) {
	return m.internal.Get(key)
}

func (m *CornelkLibrary) Set(key string, value interface{}) {
	m.internal.Set(key, value)
}

func NewCornelkLibrary() *CornelkLibrary {
	return &CornelkLibrary{
		internal: hashmap.NewSized[string, interface{}](2_000_000),
	}
}

type DustinxieLibrary struct {
	internal lockfree.HashMap
}

func (m *DustinxieLibrary) Get(key string) (interface{}, bool) {
	return m.internal.Get(key)
}

func (m *DustinxieLibrary) Set(key string, value interface{}) {
	m.internal.Set(key, value)
}

func NewDustinxieLibrary() *DustinxieLibrary {
	return &DustinxieLibrary{
		internal: lockfree.NewHashMap(),
	}
}

type AlphadoseLibrary struct {
	internal *haxmap.Map[string, interface{}]
}

func (m *AlphadoseLibrary) Get(key string) (interface{}, bool) {
	return m.internal.Get(key)
}

func (m *AlphadoseLibrary) Set(key string, value interface{}) {
	m.internal.Set(key, value)
}

func NewAlphadoseLibrary() *AlphadoseLibrary {
	return &AlphadoseLibrary{
		internal: haxmap.New[string, interface{}](),
	}
}

// ///////////////////////for EasierwayLibrary using /////////////////
// StringKey is for the string type key
type StringKey struct {
	value string
}

const (
	c1_32 uint32 = 0xcc9e2d51
	c2_32 uint32 = 0x1b873593
)

func hash(str string) uint32 {
	data := ([]byte)(str)
	var h1 uint32 = 37

	nblocks := len(data) / 4
	var p uintptr
	if len(data) > 0 {
		p = uintptr(unsafe.Pointer(&data[0]))
	}

	p1 := p + uintptr(4*nblocks)
	for ; p < p1; p += 4 {
		k1 := *(*uint32)(unsafe.Pointer(p))
		k1 *= c1_32
		k1 = (k1 << 15) | (k1 >> 17) // rotl32(k1, 15)
		k1 *= c2_32

		h1 ^= k1
		h1 = (h1 << 13) | (h1 >> 19) // rotl32(h1, 13)
		h1 = h1*5 + 0xe6546b64
	}

	tail := data[nblocks*4:]

	var k1 uint32
	switch len(tail) & 3 {
	case 3:
		k1 ^= uint32(tail[2]) << 16
		fallthrough
	case 2:
		k1 ^= uint32(tail[1]) << 8
		fallthrough
	case 1:
		k1 ^= uint32(tail[0])
		k1 *= c1_32
		k1 = (k1 << 15) | (k1 >> 17) // rotl32(k1, 15)
		k1 *= c2_32
		h1 ^= k1
	}

	h1 ^= uint32(len(data))

	h1 ^= h1 >> 16
	h1 *= 0x85ebca6b
	h1 ^= h1 >> 13
	h1 *= 0xc2b2ae35
	h1 ^= h1 >> 16

	return (h1 << 24) | (((h1 >> 8) << 16) & 0xFF0000) | (((h1 >> 16) << 8) & 0xFF00) | (h1 >> 24)
}

// PartitionID is created by string's hash
func (s *StringKey) PartitionKey() int64 {
	return int64(hash(s.value))
}

// Value is the raw string
func (s *StringKey) Value() interface{} {
	return s.value
}

// StrKey is to convert a string to StringKey
func StrKey(key string) *StringKey {
	return &StringKey{key}
}

type EasierwayLibrary struct {
	internal *ecmap.ConcurrentMap
}

func (m *EasierwayLibrary) Get(key string) (interface{}, bool) {
	return m.internal.Get(StrKey(key))
}

func (m *EasierwayLibrary) Set(key string, value interface{}) {
	m.internal.Set(StrKey(key), value)
}

func NewEasierwayLibrary() *EasierwayLibrary {
	return &EasierwayLibrary{
		internal: ecmap.CreateConcurrentMap(199),
	}
}

type MhmtszrLibrary struct {
	internal *csmap.CsMap[string, interface{}]
}

func (m *MhmtszrLibrary) Get(key string) (interface{}, bool) {
	return m.internal.Load(key)
}

func (m *MhmtszrLibrary) Set(key string, value interface{}) {
	m.internal.Store(key, value)
}

func NewMhmtszrLibrary() *MhmtszrLibrary {
	return &MhmtszrLibrary{
		internal: csmap.Create[string, interface{}](),
	}
}

type Zhangyunhao116Library struct {
	internal *skipmap.StringMap[interface{}]
}

func (m *Zhangyunhao116Library) Get(key string) (interface{}, bool) {
	return m.internal.Load(key)
}

func (m *Zhangyunhao116Library) Set(key string, value interface{}) {
	m.internal.Store(key, value)
}

func NewZhangyunhao116Library() *Zhangyunhao116Library {
	return &Zhangyunhao116Library{
		internal: skipmap.NewString[interface{}](),
	}
}

type LritaLibrary struct {
	internal lcmap.Cmap
}

func (m *LritaLibrary) Get(key string) (interface{}, bool) {
	return m.internal.Load(key)
}

func (m *LritaLibrary) Set(key string, value interface{}) {
	m.internal.Store(key, value)
}

func NewLritaLibrary() *LritaLibrary {
	return &LritaLibrary{}
}

type OneOfOneLibrary struct {
	internal *ooocmap.CMap
}

func (m *OneOfOneLibrary) Get(key string) (interface{}, bool) {
	return m.internal.Get(key), true
}

func (m *OneOfOneLibrary) Set(key string, value interface{}) {
	m.internal.Set(key, value)
}

func NewOneOfOneLibrary() *OneOfOneLibrary {
	return &OneOfOneLibrary{
		internal: ooocmap.New(),
	}
}

type AntoniomoLibrary struct {
	internal *shardedmap.StrMap
}

func (m *AntoniomoLibrary) Get(key string) (interface{}, bool) {
	return m.internal.Load(key)
}

func (m *AntoniomoLibrary) Set(key string, value interface{}) {
	m.internal.Store(key, value)
}

func NewAntoniomoLibrary() *AntoniomoLibrary {
	return &AntoniomoLibrary{
		internal: shardedmap.NewStrMap(64),
	}
}

type HfdpxLibrary struct {
	internal hcmap.ConcurrentHashMap
}

func (m *HfdpxLibrary) Get(key string) (interface{}, bool) {
	return m.internal.Get(key)
}

func (m *HfdpxLibrary) Set(key string, value interface{}) {
	m.internal.Set(key, value)
}

func NewHfdpxLibrary() *HfdpxLibrary {
	return &HfdpxLibrary{
		internal: hcmap.New(),
	}
}

type MojinfuLibrary struct {
	internal mcmap.Map
}

func (m *MojinfuLibrary) Get(key string) (interface{}, bool) {
	return m.internal.Load(key)
}

func (m *MojinfuLibrary) Set(key string, value interface{}) {
	m.internal.Store(key, value)
}

func NewMojinfuLibrary() *MojinfuLibrary {
	return &MojinfuLibrary{}
}

type Xiao7737Library struct {
	internal *xcmap.ConcurrentMap
}

func (m *Xiao7737Library) Get(key string) (interface{}, bool) {
	return m.internal.Get(StrKey(key))
}

func (m *Xiao7737Library) Set(key string, value interface{}) {
	m.internal.Set(StrKey(key), value)
}

func NewXiao7737Library() *Xiao7737Library {
	return &Xiao7737Library{
		internal: xcmap.CreateConcurrentMap(128),
	}
}

type RfyiamcoolLibrary struct {
	internal *rcmap.SyncMap
}

func (m *RfyiamcoolLibrary) Get(key string) (interface{}, bool) {
	return m.internal.Get(key)
}

func (m *RfyiamcoolLibrary) Set(key string, value interface{}) {
	m.internal.Set(key, value)
}

func NewRfyiamcoolLibrary() *RfyiamcoolLibrary {
	return &RfyiamcoolLibrary{
		internal: rcmap.New(),
	}
}

type KazuLibrary struct {
	internal *skiplistmap.Map
}

func (m *KazuLibrary) Get(key string) (interface{}, bool) {
	return m.internal.Get(key)
}

func (m *KazuLibrary) Set(key string, value interface{}) {
	m.internal.Set(key, value)
}

func NewKazuLibrary() *KazuLibrary {
	return &KazuLibrary{
		internal: skiplistmap.New(),
	}
}
