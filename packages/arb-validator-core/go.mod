module github.com/offchainlabs/arbitrum/packages/arb-validator-core

go 1.13

require (
	github.com/ethereum/go-ethereum v1.9.16
	github.com/golang/protobuf v1.5.1
	github.com/gorilla/handlers v1.4.2
	github.com/gorilla/mux v1.7.4
	github.com/kr/pretty v0.2.0 // indirect
	github.com/offchainlabs/arbitrum/packages/arb-util v0.6.5
	github.com/pkg/errors v0.9.1
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
	google.golang.org/protobuf v1.26.0
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
)

replace github.com/offchainlabs/arbitrum/packages/arb-util => ../arb-util
