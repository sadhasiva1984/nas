package nasMessage_test

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sadhasiva1984/nas"
	"github.com/sadhasiva1984/nas/logger"
	"github.com/sadhasiva1984/nas/nasMessage"
	"github.com/sadhasiva1984/nas/nasType"
)

type nasMessageRegistrationCompleteData struct {
	inExtendedProtocolDiscriminator       uint8
	inSecurityHeader                      uint8
	inSpareHalfOctet                      uint8
	inRegistrationCompleteMessageIdentity uint8
	inSORTransparentContainer             nasType.SORTransparentContainer
}

var nasMessageRegistrationCompleteTable = []nasMessageRegistrationCompleteData{
	{
		inExtendedProtocolDiscriminator:       nasMessage.Epd5GSMobilityManagementMessage,
		inSecurityHeader:                      0x01,
		inSpareHalfOctet:                      0x01,
		inRegistrationCompleteMessageIdentity: nas.MsgTypeRegistrationComplete,
		inSORTransparentContainer: nasType.SORTransparentContainer{
			Iei:    nasMessage.RegistrationCompleteSORTransparentContainerType,
			Len:    17,
			Buffer: []uint8{0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01},
		},
	},
}

func TestNasTypeNewRegistrationComplete(t *testing.T) {
	a := nasMessage.NewRegistrationComplete(0)
	assert.NotNil(t, a)
}

func TestNasTypeNewRegistrationCompleteMessage(t *testing.T) {
	for i, table := range nasMessageRegistrationCompleteTable {
		t.Logf("Test Cnt:%d", i)
		a := nasMessage.NewRegistrationComplete(0)
		b := nasMessage.NewRegistrationComplete(0)
		assert.NotNil(t, a)
		assert.NotNil(t, b)

		a.ExtendedProtocolDiscriminator.SetExtendedProtocolDiscriminator(table.inExtendedProtocolDiscriminator)
		a.SpareHalfOctetAndSecurityHeaderType.SetSecurityHeaderType(table.inSecurityHeader)
		a.SpareHalfOctetAndSecurityHeaderType.SetSpareHalfOctet(table.inSpareHalfOctet)
		a.RegistrationCompleteMessageIdentity.SetMessageType(table.inRegistrationCompleteMessageIdentity)

		a.SORTransparentContainer = nasType.NewSORTransparentContainer(nasMessage.RegistrationCompleteSORTransparentContainerType)
		a.SORTransparentContainer = &table.inSORTransparentContainer

		buff := new(bytes.Buffer)
		a.EncodeRegistrationComplete(buff)
		logger.NasMsgLog.Debugln("Encode: ", a)

		data := make([]byte, buff.Len())
		buff.Read(data)
		logger.NasMsgLog.Debugln(data)
		b.DecodeRegistrationComplete(&data)
		logger.NasMsgLog.Debugln("Decode: ", b)

		if reflect.DeepEqual(a, b) != true {
			t.Errorf("Not correct")
		}
	}
}
