package main

import (
	"encoding/hex"
	"testing"

	"github.com/radekg/kafka-protocol-go/common"
	"github.com/stretchr/testify/assert"
)

func TestEncodeDecodeResponseHeader(t *testing.T) {
	a := assert.New(t)

	reqHex := "0000003300000001"
	reqBytes, err := hex.DecodeString(reqHex)
	a.Nil(err)

	rh := &ResponseHeader{}
	err = common.Decode(reqBytes, rh)
	a.Nil(err)

	a.EqualValues(51, rh.Length)
	a.EqualValues(01, rh.CorrelationID)

	encoded, err := common.Encode(rh)
	a.Nil(err)
	a.EqualValues(reqHex, hex.EncodeToString(encoded))
}
