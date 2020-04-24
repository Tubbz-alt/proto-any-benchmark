package main

import (
	"testing"
)

func BenchmarkAnyCodec(t *testing.B) {
	DoBenchCodec(t, AnyCodec{})
}

func BenchmarkOneofCodec(t *testing.B) {
	DoBenchCodec(t, OneofCodec{})
}

func DoBenchCodec(t *testing.B, cdc Codec) {
	testData := MakeTestData(t.N)
	for _, x := range testData {
		bz, err := cdc.Marshal(x)
		if err != nil {
			panic(err)
		}
		_, err = cdc.Unmarshal(bz)
		if err != nil {
			panic(err)
		}
	}
}
