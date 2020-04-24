package main

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func BenchmarkCodecs(b *testing.B) {
	testData := MakeTestData(b.N * 100000)
	overhead := CalcAnyOverheads(testData)
	fmt.Printf("Any Overhead: %f%%\n", overhead * 100.0)
	b.Run("any roundtrip", func(b *testing.B) {
		DoBenchCodec(b, testData, AnyCodec{})
	})
	b.Run("oneof roundtrip", func(b *testing.B) {
		DoBenchCodec(b, testData, OneofCodec{})
	})
	var bz [][]byte
	b.Run("any marshal", func(b *testing.B) {
		bz = DoBenchMarshal(b, testData, AnyCodec{})
	})
	anyBytes := SumBytes(bz)
	b.Run("any unmarshal", func(b *testing.B) {
		DoBenchUnmarshal(b, bz, AnyCodec{})
	})
	b.Run("oneof marshal", func(b *testing.B) {
		bz = DoBenchMarshal(b, testData, OneofCodec{})
	})
	b.Run("oneof unmarshal", func(b *testing.B) {
		DoBenchUnmarshal(b, bz, OneofCodec{})
	})
	oneofBytes := SumBytes(bz)
	fmt.Printf("Encoded Any overhead: %f%%\n", ((float64(anyBytes) - float64(oneofBytes)) / float64(anyBytes)) * 100.0)
}

func DoBenchCodec(t *testing.B, testData []Msg, cdc Codec) {
	for _, x := range testData {
		bz, err := cdc.Marshal(x)
		require.NoError(t, err)
		y, err := cdc.Unmarshal(bz)
		require.NoError(t, err)
		require.Equal(t, x, y)
	}
}

func DoBenchMarshal(t *testing.B, testData []Msg, cdc Codec) [][]byte {
	res := make([][]byte, len(testData))
	for i, x := range testData {
		bz, err := cdc.Marshal(x)
		require.NoError(t, err)
		res[i] = bz
	}
	return res
}

func DoBenchUnmarshal(t *testing.B, testData [][]byte, cdc Codec) {
	for _, x := range testData {
		_, err := cdc.Unmarshal(x)
		require.NoError(t, err)
	}
}

func SumBytes(bzs [][]byte) int {
	sum := 0
	for _, bz := range bzs {
		sum += len(bz)
	}
	return sum
}
