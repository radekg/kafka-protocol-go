package sasl

import (
	"encoding/hex"
	"testing"

	"github.com/radekg/kafka-protocol-go/common"
	"github.com/radekg/kafka-protocol-go/request"
	"github.com/stretchr/testify/assert"
)

func TestEncodeDecodeSaslAuthenticateRequestV0(t *testing.T) {
	a := assert.New(t)

	reqHex := "002400000000000100144b61666b614578616d706c6550726f64756365720000001e006d792d746573742d75736572006d792d746573742d70617373776f7264"
	reqBytes, err := hex.DecodeString(reqHex)
	a.Nil(err)

	body := &SaslAuthenticateRequestV0{}
	request := &request.Request{Body: body}
	err = common.Decode(reqBytes, request)
	a.Nil(err)

	a.EqualValues(1, request.CorrelationID)
	a.EqualValues("KafkaExampleProducer", request.ClientID)
	a.EqualValues("006d792d746573742d75736572006d792d746573742d70617373776f7264", hex.EncodeToString(body.SaslAuthBytes))

	encoded, err := common.Encode(request)
	a.Nil(err)
	a.EqualValues(reqHex, hex.EncodeToString(encoded))

}
