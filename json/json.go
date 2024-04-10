package jsontest

import (
	"encoding/json"

	"github.com/bytedance/sonic"
	jsoniter "github.com/json-iterator/go"
	"github.com/pquerna/ffjson/ffjson"
)

var jsonx = jsoniter.ConfigCompatibleWithStandardLibrary

type JsonLib interface {
	Marshal(interface{}) ([]byte, error)
	UnMarshal([]byte, interface{}) error
}

type OriginLibrary struct {
}

func (m *OriginLibrary) Marshal(data interface{}) ([]byte, error) {
	return json.Marshal(&data)
}

func (m *OriginLibrary) UnMarshal(output []byte, data interface{}) error {
	return json.Unmarshal(output, &data)
}
func NewOriginLibrary() *OriginLibrary {
	return &OriginLibrary{}
}

type SonicLibrary struct {
}

func (m *SonicLibrary) Marshal(data interface{}) ([]byte, error) {
	return sonic.Marshal(&data)
}

func (m *SonicLibrary) UnMarshal(output []byte, data interface{}) error {
	return sonic.Unmarshal(output, &data)
}

func NewSonicLibrary() *SonicLibrary {
	return &SonicLibrary{}
}

type JsoniterLibrary struct {
	internal interface{}
}

func (m *JsoniterLibrary) Marshal(data interface{}) ([]byte, error) {
	return jsonx.Marshal(data)
}

func (m *JsoniterLibrary) UnMarshal(output []byte, data interface{}) error {
	return jsonx.Unmarshal(output, data)
}
func NewJsoniterLibrary() *JsoniterLibrary {
	return &JsoniterLibrary{}
}

// github.com/mailru/easyjson need generate easyjson code first
// type MailruLibrary struct {
// 	internal interface{}
// }

// func (m *MailruLibrary) Marshal(data interface{}) ([]byte, error) {
// 	return easyjson.Marshal(data)
// }

// func (m *MailruLibrary) UnMarshal(output []byte, data interface{}) error {
// 	return easyjson.Unmarshal(output, data)
// }
// func NewMailruLibrary() *MailruLibrary {
// 	return &MailruLibrary{}
// }

type PquernaLibrary struct {
	internal interface{}
}

func (m *PquernaLibrary) Marshal(data interface{}) ([]byte, error) {
	return ffjson.Marshal(data)
}

func (m *PquernaLibrary) UnMarshal(output []byte, data interface{}) error {
	return ffjson.Unmarshal(output, data)
}

func NewPquernaLibrary() *PquernaLibrary {
	return &PquernaLibrary{}
}

// github.com/CosmWasm/tinyjson need generate tinyjson code first

// type CosmWasmLibrary struct {
// 	internal interface{}
// }

// func (m *CosmWasmLibrary) Marshal(data interface{}) ([]byte, error) {
// 	return tinyjson.Marshal(data)
// }

// func (m *CosmWasmLibrary) UnMarshal(output []byte, data interface{}) error {
// 	return tinyjson.Unmarshal(output, data)
// }

// func NewCosmWasmLibrary() *CosmWasmLibrary {
// 	return &CosmWasmLibrary{}
// }
