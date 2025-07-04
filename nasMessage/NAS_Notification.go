// Code generated by generate.sh, DO NOT EDIT.

package nasMessage

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/sadhasiva1984/nas/nasType"
)

type Notification struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.SpareHalfOctetAndSecurityHeaderType
	nasType.NotificationMessageIdentity
	nasType.SpareHalfOctetAndAccessType
}

func NewNotification(iei uint8) (notification *Notification) {
	notification = &Notification{}
	return notification
}

func (a *Notification) EncodeNotification(buffer *bytes.Buffer) error {
	if err := binary.Write(buffer, binary.BigEndian, a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return fmt.Errorf("NAS encode error (Notification/ExtendedProtocolDiscriminator): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return fmt.Errorf("NAS encode error (Notification/SpareHalfOctetAndSecurityHeaderType): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.NotificationMessageIdentity.Octet); err != nil {
		return fmt.Errorf("NAS encode error (Notification/NotificationMessageIdentity): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.SpareHalfOctetAndAccessType.Octet); err != nil {
		return fmt.Errorf("NAS encode error (Notification/SpareHalfOctetAndAccessType): %w", err)
	}
	return nil
}

func (a *Notification) DecodeNotification(byteArray *[]byte) error {
	buffer := bytes.NewBuffer(*byteArray)
	if err := binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return fmt.Errorf("NAS decode error (Notification/ExtendedProtocolDiscriminator): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return fmt.Errorf("NAS decode error (Notification/SpareHalfOctetAndSecurityHeaderType): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.NotificationMessageIdentity.Octet); err != nil {
		return fmt.Errorf("NAS decode error (Notification/NotificationMessageIdentity): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndAccessType.Octet); err != nil {
		return fmt.Errorf("NAS decode error (Notification/SpareHalfOctetAndAccessType): %w", err)
	}
	for buffer.Len() > 0 {
		var ieiN uint8
		var tmpIeiN uint8
		if err := binary.Read(buffer, binary.BigEndian, &ieiN); err != nil {
			return fmt.Errorf("NAS decode error (Notification/iei): %w", err)
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
