package main

import (
	"encoding/hex"
	"testing"

	"github.com/radekg/kafka-protocol-go/common"
	"github.com/stretchr/testify/assert"
)

func TestEncodeDecodeResponseHeaderV1(t *testing.T) {
	a := assert.New(t)

	reqHex := "000000390000000200"
	reqBytes, err := hex.DecodeString(reqHex)
	a.Nil(err)

	rh := &ResponseHeaderV1{}
	err = common.Decode(reqBytes, rh)
	a.Nil(err)

	a.EqualValues(57, rh.Length)
	a.EqualValues(02, rh.CorrelationID)
	a.Len(rh.RawTaggedFields, 0)

	encoded, err := common.Encode(rh)
	a.Nil(err)
	a.EqualValues(reqHex, hex.EncodeToString(encoded))
}
