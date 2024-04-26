//go:generate protoc -I=proto -I=$GOPATH/src -I=$GOPATH/src/github.com/multiversx/protobuf/protobuf  --gogoslick_out=. dct.proto
package dct

import "math/big"

// New returns a new batch from given buffers
func New() *ESDigitalToken {
	return &ESDigitalToken{
		Value: big.NewInt(0),
	}
}
