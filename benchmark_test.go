package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAnyCodec(t *testing.T) {
	DoTestCodec(t, AnyCodec{})
}

func TestOneofCodec(t *testing.T) {
	DoTestCodec(t, OneofCodec{})
}

func DoTestCodec(t *testing.T, cdc Codec) {
	testData := MakeTestData(10)
	for _, x := range testData {
		bz, err := cdc.Marshal(x)
		require.NoError(t, err)
		y, err := cdc.Unmarshal(bz)
		require.NoError(t, err)
		require.Equal(t, x, y)
	}
}

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
