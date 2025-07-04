// Code generated by generate.sh, DO NOT EDIT.

package nasMessage

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/sadhasiva1984/nas/nasType"
)

type SecurityProtected5GSNASMessage struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.SpareHalfOctetAndSecurityHeaderType
	nasType.MessageAuthenticationCode
	nasType.SequenceNumber
	nasType.Plain5GSNASMessage
}

func NewSecurityProtected5GSNASMessage(iei uint8) (securityProtected5GSNASMessage *SecurityProtected5GSNASMessage) {
	securityProtected5GSNASMessage = &SecurityProtected5GSNASMessage{}
	return securityProtected5GSNASMessage
}

func (a *SecurityProtected5GSNASMessage) EncodeSecurityProtected5GSNASMessage(buffer *bytes.Buffer) error {
	if err := binary.Write(buffer, binary.BigEndian, a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return fmt.Errorf("NAS encode error (SecurityProtected5GSNASMessage/ExtendedProtocolDiscriminator): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return fmt.Errorf("NAS encode error (SecurityProtected5GSNASMessage/SpareHalfOctetAndSecurityHeaderType): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.MessageAuthenticationCode.Octet[:]); err != nil {
		return fmt.Errorf("NAS encode error (SecurityProtected5GSNASMessage/MessageAuthenticationCode): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.SequenceNumber.Octet); err != nil {
		return fmt.Errorf("NAS encode error (SecurityProtected5GSNASMessage/SequenceNumber): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, &a.Plain5GSNASMessage); err != nil {
		return fmt.Errorf("NAS encode error (SecurityProtected5GSNASMessage/Plain5GSNASMessage): %w", err)
	}
	return nil
}

func (a *SecurityProtected5GSNASMessage) DecodeSecurityProtected5GSNASMessage(byteArray *[]byte) error {
	buffer := bytes.NewBuffer(*byteArray)
	if err := binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return fmt.Errorf("NAS decode error (SecurityProtected5GSNASMessage/ExtendedProtocolDiscriminator): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return fmt.Errorf("NAS decode error (SecurityProtected5GSNASMessage/SpareHalfOctetAndSecurityHeaderType): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, a.MessageAuthenticationCode.Octet[:]); err != nil {
		return fmt.Errorf("NAS decode error (SecurityProtected5GSNASMessage/MessageAuthenticationCode): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.SequenceNumber.Octet); err != nil {
		return fmt.Errorf("NAS decode error (SecurityProtected5GSNASMessage/SequenceNumber): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.Plain5GSNASMessage); err != nil {
		return fmt.Errorf("NAS decode error (SecurityProtected5GSNASMessage/Plain5GSNASMessage): %w", err)
	}
	for buffer.Len() > 0 {
		var ieiN uint8
		var tmpIeiN uint8
		if err := binary.Read(buffer, binary.BigEndian, &ieiN); err != nil {
			return fmt.Errorf("NAS decode error (SecurityProtected5GSNASMessage/iei): %w", err)
		}
		// fmt.Println(ieiN)
		if ieiN >= 0x80 {
			tmpIeiN = (ieiN & 0xf0) >> 4
		} else {
			tmpIeiN = ieiN
		}
		// fmt.Println("type", tmpIeiN)
		switch tmpIeiN {
		default:
		}
	}
	return nil
}
