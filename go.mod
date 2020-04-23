module github.com/regen-network/proto-any-benchmark

require (
	github.com/brianvoe/gofakeit/v5 v5.4.2
	github.com/gogo/protobuf v1.3.1
	github.com/regen-network/cosmos-proto v0.2.2
	github.com/stretchr/testify v1.5.1
	github.com/tendermint/tm-db v0.5.1
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.1

go 1.14
