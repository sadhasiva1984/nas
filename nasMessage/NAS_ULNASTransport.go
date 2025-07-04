// Code generated by generate.sh, DO NOT EDIT.

package nasMessage

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/sadhasiva1984/nas/nasType"
)

type ULNASTransport struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.SpareHalfOctetAndSecurityHeaderType
	nasType.ULNASTRANSPORTMessageIdentity
	nasType.SpareHalfOctetAndPayloadContainerType
	nasType.PayloadContainer
	*nasType.PduSessionID2Value
	*nasType.OldPDUSessionID
	*nasType.RequestType
	*nasType.SNSSAI
	*nasType.DNN
	*nasType.AdditionalInformation
}

func NewULNASTransport(iei uint8) (uLNASTransport *ULNASTransport) {
	uLNASTransport = &ULNASTransport{}
	return uLNASTransport
}

const (
	ULNASTransportPduSessionID2ValueType    uint8 = 0x12
	ULNASTransportOldPDUSessionIDType       uint8 = 0x59
	ULNASTransportRequestTypeType           uint8 = 0x08
	ULNASTransportSNSSAIType                uint8 = 0x22
	ULNASTransportDNNType                   uint8 = 0x25
	ULNASTransportAdditionalInformationType uint8 = 0x24
)

func (a *ULNASTransport) EncodeULNASTransport(buffer *bytes.Buffer) error {
	if err := binary.Write(buffer, binary.BigEndian, a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return fmt.Errorf("NAS encode error (ULNASTransport/ExtendedProtocolDiscriminator): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return fmt.Errorf("NAS encode error (ULNASTransport/SpareHalfOctetAndSecurityHeaderType): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.ULNASTRANSPORTMessageIdentity.Octet); err != nil {
		return fmt.Errorf("NAS encode error (ULNASTransport/ULNASTRANSPORTMessageIdentity): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.SpareHalfOctetAndPayloadContainerType.Octet); err != nil {
		return fmt.Errorf("NAS encode error (ULNASTransport/SpareHalfOctetAndPayloadContainerType): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.PayloadContainer.GetLen()); err != nil {
		return fmt.Errorf("NAS encode error (ULNASTransport/PayloadContainer): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.PayloadContainer.Buffer); err != nil {
		return fmt.Errorf("NAS encode error (ULNASTransport/PayloadContainer): %w", err)
	}
	if a.PduSessionID2Value != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.PduSessionID2Value.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (ULNASTransport/PduSessionID2Value): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.PduSessionID2Value.Octet); err != nil {
			return fmt.Errorf("NAS encode error (ULNASTransport/PduSessionID2Value): %w", err)
		}
	}
	if a.OldPDUSessionID != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.OldPDUSessionID.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (ULNASTransport/OldPDUSessionID): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.OldPDUSessionID.Octet); err != nil {
			return fmt.Errorf("NAS encode error (ULNASTransport/OldPDUSessionID): %w", err)
		}
	}
	if a.RequestType != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.RequestType.Octet); err != nil {
			return fmt.Errorf("NAS encode error (ULNASTransport/RequestType): %w", err)
		}
	}
	if a.SNSSAI != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.SNSSAI.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (ULNASTransport/SNSSAI): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.SNSSAI.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (ULNASTransport/SNSSAI): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.SNSSAI.Octet[:a.SNSSAI.GetLen()]); err != nil {
			return fmt.Errorf("NAS encode error (ULNASTransport/SNSSAI): %w", err)
		}
	}
	if a.DNN != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.DNN.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (ULNASTransport/DNN): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.DNN.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (ULNASTransport/DNN): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.DNN.Buffer); err != nil {
			return fmt.Errorf("NAS encode error (ULNASTransport/DNN): %w", err)
		}
	}
	if a.AdditionalInformation != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.AdditionalInformation.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (ULNASTransport/AdditionalInformation): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.AdditionalInformation.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (ULNASTransport/AdditionalInformation): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.AdditionalInformation.Buffer); err != nil {
			return fmt.Errorf("NAS encode error (ULNASTransport/AdditionalInformation): %w", err)
		}
	}
	return nil
}

func (a *ULNASTransport) DecodeULNASTransport(byteArray *[]byte) error {
	buffer := bytes.NewBuffer(*byteArray)
	if err := binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return fmt.Errorf("NAS decode error (ULNASTransport/ExtendedProtocolDiscriminator): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return fmt.Errorf("NAS decode error (ULNASTransport/SpareHalfOctetAndSecurityHeaderType): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.ULNASTRANSPORTMessageIdentity.Octet); err != nil {
		return fmt.Errorf("NAS decode error (ULNASTransport/ULNASTRANSPORTMessageIdentity): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndPayloadContainerType.Octet); err != nil {
		return fmt.Errorf("NAS decode error (ULNASTransport/SpareHalfOctetAndPayloadContainerType): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.PayloadContainer.Len); err != nil {
		return fmt.Errorf("NAS decode error (ULNASTransport/PayloadContainer): %w", err)
	}
	if a.PayloadContainer.Len < 1 {
		return fmt.Errorf("invalid ie length (ULNASTransport/PayloadContainer): %d", a.PayloadContainer.Len)
	}
	a.PayloadContainer.SetLen(a.PayloadContainer.GetLen())
	if err := binary.Read(buffer, binary.BigEndian, a.PayloadContainer.Buffer); err != nil {
		return fmt.Errorf("NAS decode error (ULNASTransport/PayloadContainer): %w", err)
	}
	for buffer.Len() > 0 {
		var ieiN uint8
		var tmpIeiN uint8
		if err := binary.Read(buffer, binary.BigEndian, &ieiN); err != nil {
			return fmt.Errorf("NAS decode error (ULNASTransport/iei): %w", err)
		}
		// fmt.Println(ieiN)
		if ieiN >= 0x80 {
			tmpIeiN = (ieiN & 0xf0) >> 4
		} else {
			tmpIeiN = ieiN
		}
		// fmt.Println("type", tmpIeiN)
		switch tmpIeiN {
		case ULNASTransportPduSessionID2ValueType:
			a.PduSessionID2Value = nasType.NewPduSessionID2Value(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.PduSessionID2Value.Octet); err != nil {
				return fmt.Errorf("NAS decode error (ULNASTransport/PduSessionID2Value): %w", err)
			}
		case ULNASTransportOldPDUSessionIDType:
			a.OldPDUSessionID = nasType.NewOldPDUSessionID(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.OldPDUSessionID.Octet); err != nil {
				return fmt.Errorf("NAS decode error (ULNASTransport/OldPDUSessionID): %w", err)
			}
		case ULNASTransportRequestTypeType:
			a.RequestType = nasType.NewRequestType(ieiN)
			a.RequestType.Octet = ieiN
		case ULNASTransportSNSSAIType:
			a.SNSSAI = nasType.NewSNSSAI(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.SNSSAI.Len); err != nil {
				return fmt.Errorf("NAS decode error (ULNASTransport/SNSSAI): %w", err)
			}
			if a.SNSSAI.Len < 1 || a.SNSSAI.Len > 8 {
				return fmt.Errorf("invalid ie length (ULNASTransport/SNSSAI): %d", a.SNSSAI.Len)
			}
			a.SNSSAI.SetLen(a.SNSSAI.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.SNSSAI.Octet[:a.SNSSAI.GetLen()]); err != nil {
				return fmt.Errorf("NAS decode error (ULNASTransport/SNSSAI): %w", err)
			}
		case ULNASTransportDNNType:
			a.DNN = nasType.NewDNN(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.DNN.Len); err != nil {
				return fmt.Errorf("NAS decode error (ULNASTransport/DNN): %w", err)
			}
			if a.DNN.Len < 1 || a.DNN.Len > 100 {
				return fmt.Errorf("invalid ie length (ULNASTransport/DNN): %d", a.DNN.Len)
			}
			a.DNN.SetLen(a.DNN.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.DNN.Buffer); err != nil {
				return fmt.Errorf("NAS decode error (ULNASTransport/DNN): %w", err)
			}
		case ULNASTransportAdditionalInformationType:
			a.AdditionalInformation = nasType.NewAdditionalInformation(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.AdditionalInformation.Len); err != nil {
				return fmt.Errorf("NAS decode error (ULNASTransport/AdditionalInformation): %w", err)
			}
			if a.AdditionalInformation.Len < 1 {
				return fmt.Errorf("invalid ie length (ULNASTransport/AdditionalInformation): %d", a.AdditionalInformation.Len)
			}
			a.AdditionalInformation.SetLen(a.AdditionalInformation.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.AdditionalInformation.Buffer); err != nil {
				return fmt.Errorf("NAS decode error (ULNASTransport/AdditionalInformation): %w", err)
			}
		default:
		}
	}
	return nil
}
