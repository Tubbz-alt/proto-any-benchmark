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
	db "github.com/tendermint/tm-db"
	"math/rand"
)

type Msg interface {
	proto.Message
}

type Fixture struct {
	DB db.DB
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

type OneofCodec struct { }

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

type AnyCodec struct { }

func (a AnyCodec) Marshal(x proto.Message) ([]byte, error) {
	any, err := types.MarshalAny(x)
	if err != nil {
		return nil, err
	}
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

func main() {
	testData := MakeTestData(100)
	fmt.Printf("%+v", testData)
}
