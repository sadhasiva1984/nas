package nasMessage_test

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sadhasiva1984/nas/logger"

	//"fmt"
	"github.com/sadhasiva1984/nas/nasMessage"
	"github.com/sadhasiva1984/nas/nasType"
)

type nasMessageAuthenticationResponseData struct {
	inExtendedProtocolDiscriminator         uint8
	inSecurityHeader                        uint8
	inSpareHalfOctet                        uint8
	inAuthenticationResponseMessageIdentity uint8
	inAuthenticationResponseParameter       nasType.AuthenticationResponseParameter
	inEAPMessage                            nasType.EAPMessage
}

var nasMessageAuthenticationResponseTable = []nasMessageAuthenticationResponseData{
	{
		inExtendedProtocolDiscriminator:         0x01,
		inSecurityHeader:                        0x08,
		inSpareHalfOctet:                        0x01,
		inAuthenticationResponseMessageIdentity: 0x01,
		inAuthenticationResponseParameter:       nasType.AuthenticationResponseParameter{nasMessage.AuthenticationResponseAuthenticationResponseParameterType, 16, [16]uint8{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}},
		inEAPMessage:                            nasType.EAPMessage{nasMessage.AuthenticationResponseEAPMessageType, 4, []uint8{0x01, 0x01, 0x01, 0x01}},
	},
}

func TestNasTypeNewAuthenticationResponse(t *testing.T) {
	a := nasMessage.NewAuthenticationResponse(0)
	assert.NotNil(t, a)
}

func TestNasTypeNewAuthenticationResponseMessage(t *testing.T) {
	logger.NasMsgLog.Infoln("---Test NAS Message: AuthenticationResponseMessage---")
	for i, table := range nasMessageAuthenticationResponseTable {
		logger.NasMsgLog.Infoln("Test Cnt:", i)
		a := nasMessage.NewAuthenticationResponse(0)
		b := nasMessage.NewAuthenticationResponse(0)
		assert.NotNil(t, a)
		assert.NotNil(t, b)

		a.ExtendedProtocolDiscriminator.SetExtendedProtocolDiscriminator(table.inExtendedProtocolDiscriminator)
		a.SpareHalfOctetAndSecurityHeaderType.SetSecurityHeaderType(table.inSecurityHeader)
		a.SpareHalfOctetAndSecurityHeaderType.SetSpareHalfOctet(table.inSpareHalfOctet)
		a.AuthenticationResponseMessageIdentity.SetMessageType(table.inAuthenticationResponseMessageIdentity)

		a.AuthenticationResponseParameter = nasType.NewAuthenticationResponseParameter(nasMessage.AuthenticationResponseAuthenticationResponseParameterType)
		a.AuthenticationResponseParameter = &table.inAuthenticationResponseParameter

		a.EAPMessage = nasType.NewEAPMessage(nasMessage.AuthenticationResponseEAPMessageType)
		a.EAPMessage = &table.inEAPMessage

		buff := new(bytes.Buffer)
		a.EncodeAuthenticationResponse(buff)
		logger.NasMsgLog.Debugln("Encode: ", a)

		data := make([]byte, buff.Len())
		buff.Read(data)
		b.DecodeAuthenticationResponse(&data)
		logger.NasMsgLog.Debugln(data)
		logger.NasMsgLog.Debugln("Decode: ", b)

		if reflect.DeepEqual(a, b) != true {
			t.Errorf("Not correct")
		}

	}
}
