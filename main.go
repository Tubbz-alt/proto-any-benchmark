package main

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v5"
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/types"
	"github.com/regen-network/proto-any-benchmark/a"
	"github.com/regen-network/proto-any-benchmark/b"
	"github.com/regen-network/proto-any-benchmark/c"
	"github.com/regen-network/proto-any-benchmark/d"
	"github.com/regen-network/proto-any-benchmark/e"
	"github.com/tecbot/gorocksdb"
	db "github.com/tendermint/tm-db"
	"math/rand"
	"runtime"
)

type Msg interface {
	proto.Message
}

func MakeTestData(n int) []Msg {
	res := make([]Msg, n)
	for i := 0; i < n; i++ {
		var x Msg
		r := rand.Intn(15)
		switch r {
		case 0:
			x = &a.MsgLpctqrWokmmwtlm{}
		case 1:
			x = &a.MsgMrtsdxfVexxblyqrzg{}
		case 2:
			x = &a.MsgPousljlwlybJloivfeuc{}
		case 3:
			x = &b.MsgMliexjcuyIcbatudgl{}
		case 4:
			x = &b.MsgVedycrEaaidqrielx{}
		case 5:
			x = &b.MsgYvclzswKwoiqfnvf{}
		case 6:
			x = &c.MsgCrbjxhtnahzEkoyklxdc{}
		case 7:
			x = &c.MsgIggmhogotlAdolnphpfi{}
		case 8:
			x = &c.MsgUzewjmoiKursyvlk{}
		case 9:
			x = &d.MsgBtolfeintfGglqwlxnq{}
		case 10:
			x = &d.MsgMuqmoycgmDomajqaep{}
		case 11:
			x = &d.MsgTqviyexlSsdnfgf{}
		case 12:
			x = &e.MsgJirzatkEjglwxlp{}
		case 13:
			x = &e.MsgUzaztuwkehwWuhanr{}
		case 14:
			x = &e.MsgXoxmigzwnxqIniehuizyb{}
		}
		gofakeit.Struct(x)
		res[i] = x
	}
	return res
}

type Codec interface {
	Marshal(x proto.Message) ([]byte, error)
	Unmarshal(bz []byte) (interface{}, error)
}

type OneofCodec struct{}

func (o OneofCodec) Marshal(x proto.Message) ([]byte, error) {
	oneof := OneOfTest{}
	err := oneof.SetMsg(x)
	if err != nil {
		return nil, err
	}
	return oneof.Marshal()
}

func (o OneofCodec) Unmarshal(bz []byte) (interface{}, error) {
	var oneof OneOfTest
	err := oneof.Unmarshal(bz)
	if err != nil {
		return nil, err
	}
	return oneof.GetMsg(), nil
}

type AnyCodec struct{}

func (a AnyCodec) Marshal(x proto.Message) ([]byte, error) {
	value, err := proto.Marshal(x)
	if err != nil {
		return nil, err
	}
	any := &types.Any{TypeUrl: "/" + proto.MessageName(x), Value: value}
	return any.Marshal()
}

func (a AnyCodec) Unmarshal(bz []byte) (interface{}, error) {
	var any types.Any
	err := any.Unmarshal(bz)
	if err != nil {
		return nil, err
	}
	var dynAny types.DynamicAny
	err = types.UnmarshalAny(&any, &dynAny)
	if err != nil {
		return nil, err
	}
	return dynAny.Message, nil
}

func LoadDB(db db.DB, cdc Codec, data []Msg) {
	for i, x := range data {
		bz, err := cdc.Marshal(x)
		if err != nil {
			panic(err)
		}
		err = db.SetSync([]byte(fmt.Sprintf("%d", i)), bz)
		if err != nil {
			panic(err)
		}
	}
}

func TestDB(db db.DB, cdc Codec, data []Msg) {
	LoadDB(db, cdc, data)
	err := db.Close()
	if err != nil {
		panic(err)
	}
}

func CommonRocksOpts() *gorocksdb.Options {
	bbto := gorocksdb.NewDefaultBlockBasedTableOptions()
	bbto.SetBlockCache(gorocksdb.NewLRUCache(1 << 30))
	bbto.SetFilterPolicy(gorocksdb.NewBloomFilter(10))

	opts := gorocksdb.NewDefaultOptions()
	opts.SetBlockBasedTableFactory(bbto)
	opts.SetCreateIfMissing(true)
	opts.IncreaseParallelism(runtime.NumCPU())
	// 1.5GB maximum memory use for writebuffer.
	opts.OptimizeLevelStyleCompaction(512 * 1024 * 1024)
	opts.SetCompression(gorocksdb.LZ4Compression)
	return opts
}

func LZ4RocksOpts() *gorocksdb.Options {
	opts := CommonRocksOpts()
	opts.SetCompression(gorocksdb.LZ4Compression)
	copts := gorocksdb.NewDefaultCompressionOptions()
	copts.MaxDictBytes = 1000000
	opts.SetCompressionOptions(copts)
	return opts
}

func ZSTDRocksOpts() *gorocksdb.Options {
	opts := CommonRocksOpts()
	opts.SetCompression(gorocksdb.ZSTDCompression)
	copts := gorocksdb.NewDefaultCompressionOptions()
	copts.MaxDictBytes = 1000000
	opts.SetCompressionOptions(copts)
	return opts
}

func NewRocksDBWithopts(name string, opts *gorocksdb.Options) *db.RocksDB {
	db, err := db.NewRocksDBWithOptions(name, "db", opts)
	if err != nil {
		panic(err)
	}
	return db
}

func main() {
	testData := MakeTestData(10000)
	TestDB(db.NewDB("goleveldb-oneof", db.GoLevelDBBackend, "db"), OneofCodec{}, testData)
	TestDB(db.NewDB("goleveldb-any", db.GoLevelDBBackend, "db"), AnyCodec{}, testData)
	TestDB(db.NewDB("rocksdb-oneof", db.RocksDBBackend, "db"), OneofCodec{}, testData)
	TestDB(db.NewDB("rocksdb-any", db.RocksDBBackend, "db"), AnyCodec{}, testData)
	TestDB(NewRocksDBWithopts("rocksdb-lz4-oneof", LZ4RocksOpts()), OneofCodec{}, testData)
	TestDB(NewRocksDBWithopts("rocksdb-lz4-any", LZ4RocksOpts()), AnyCodec{}, testData)
	TestDB(NewRocksDBWithopts("rocksdb-zstd-oneof", ZSTDRocksOpts()), OneofCodec{}, testData)
	TestDB(NewRocksDBWithopts("rocksdb-zstd-any", ZSTDRocksOpts()), AnyCodec{}, testData)
}
