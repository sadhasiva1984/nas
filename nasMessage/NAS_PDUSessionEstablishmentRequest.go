// Code generated by generate.sh, DO NOT EDIT.

package nasMessage

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/sadhasiva1984/nas/nasType"
)

type PDUSessionEstablishmentRequest struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.PDUSessionID
	nasType.PTI
	nasType.PDUSESSIONESTABLISHMENTREQUESTMessageIdentity
	nasType.IntegrityProtectionMaximumDataRate
	*nasType.PDUSessionType
	*nasType.SSCMode
	*nasType.Capability5GSM
	*nasType.MaximumNumberOfSupportedPacketFilters
	*nasType.AlwaysonPDUSessionRequested
	*nasType.SMPDUDNRequestContainer
	*nasType.ExtendedProtocolConfigurationOptions
}

func NewPDUSessionEstablishmentRequest(iei uint8) (pDUSessionEstablishmentRequest *PDUSessionEstablishmentRequest) {
	pDUSessionEstablishmentRequest = &PDUSessionEstablishmentRequest{}
	return pDUSessionEstablishmentRequest
}

const (
	PDUSessionEstablishmentRequestPDUSessionTypeType                        uint8 = 0x09
	PDUSessionEstablishmentRequestSSCModeType                               uint8 = 0x0A
	PDUSessionEstablishmentRequestCapability5GSMType                        uint8 = 0x28
	PDUSessionEstablishmentRequestMaximumNumberOfSupportedPacketFiltersType uint8 = 0x55
	PDUSessionEstablishmentRequestAlwaysonPDUSessionRequestedType           uint8 = 0x0B
	PDUSessionEstablishmentRequestSMPDUDNRequestContainerType               uint8 = 0x39
	PDUSessionEstablishmentRequestExtendedProtocolConfigurationOptionsType  uint8 = 0x7B
)

func (a *PDUSessionEstablishmentRequest) EncodePDUSessionEstablishmentRequest(buffer *bytes.Buffer) error {
	if err := binary.Write(buffer, binary.BigEndian, a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return fmt.Errorf("NAS encode error (PDUSessionEstablishmentRequest/ExtendedProtocolDiscriminator): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.PDUSessionID.Octet); err != nil {
		return fmt.Errorf("NAS encode error (PDUSessionEstablishmentRequest/PDUSessionID): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.PTI.Octet); err != nil {
		return fmt.Errorf("NAS encode error (PDUSessionEstablishmentRequest/PTI): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.PDUSESSIONESTABLISHMENTREQUESTMessageIdentity.Octet); err != nil {
		return fmt.Errorf("NAS encode error (PDUSessionEstablishmentRequest/PDUSESSIONESTABLISHMENTREQUESTMessageIdentity): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.IntegrityProtectionMaximumDataRate.Octet[:]); err != nil {
		return fmt.Errorf("NAS encode error (PDUSessionEstablishmentRequest/IntegrityProtectionMaximumDataRate): %w", err)
	}
	if a.PDUSessionType != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.PDUSessionType.Octet); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionEstablishmentRequest/PDUSessionType): %w", err)
		}
	}
	if a.SSCMode != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.SSCMode.Octet); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionEstablishmentRequest/SSCMode): %w", err)
		}
	}
	if a.Capability5GSM != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.Capability5GSM.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionEstablishmentRequest/Capability5GSM): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.Capability5GSM.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionEstablishmentRequest/Capability5GSM): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.Capability5GSM.Octet[:a.Capability5GSM.GetLen()]); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionEstablishmentRequest/Capability5GSM): %w", err)
		}
	}
	if a.MaximumNumberOfSupportedPacketFilters != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.MaximumNumberOfSupportedPacketFilters.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionEstablishmentRequest/MaximumNumberOfSupportedPacketFilters): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.MaximumNumberOfSupportedPacketFilters.Octet[:]); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionEstablishmentRequest/MaximumNumberOfSupportedPacketFilters): %w", err)
		}
	}
	if a.AlwaysonPDUSessionRequested != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.AlwaysonPDUSessionRequested.Octet); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionEstablishmentRequest/AlwaysonPDUSessionRequested): %w", err)
		}
	}
	if a.SMPDUDNRequestContainer != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.SMPDUDNRequestContainer.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionEstablishmentRequest/SMPDUDNRequestContainer): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.SMPDUDNRequestContainer.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionEstablishmentRequest/SMPDUDNRequestContainer): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.SMPDUDNRequestContainer.Buffer); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionEstablishmentRequest/SMPDUDNRequestContainer): %w", err)
		}
	}
	if a.ExtendedProtocolConfigurationOptions != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.ExtendedProtocolConfigurationOptions.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionEstablishmentRequest/ExtendedProtocolConfigurationOptions): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.ExtendedProtocolConfigurationOptions.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionEstablishmentRequest/ExtendedProtocolConfigurationOptions): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.ExtendedProtocolConfigurationOptions.Buffer); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionEstablishmentRequest/ExtendedProtocolConfigurationOptions): %w", err)
		}
	}
	return nil
}

func (a *PDUSessionEstablishmentRequest) DecodePDUSessionEstablishmentRequest(byteArray *[]byte) error {
	buffer := bytes.NewBuffer(*byteArray)
	if err := binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return fmt.Errorf("NAS decode error (PDUSessionEstablishmentRequest/ExtendedProtocolDiscriminator): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.PDUSessionID.Octet); err != nil {
		return fmt.Errorf("NAS decode error (PDUSessionEstablishmentRequest/PDUSessionID): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.PTI.Octet); err != nil {
		return fmt.Errorf("NAS decode error (PDUSessionEstablishmentRequest/PTI): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.PDUSESSIONESTABLISHMENTREQUESTMessageIdentity.Octet); err != nil {
		return fmt.Errorf("NAS decode error (PDUSessionEstablishmentRequest/PDUSESSIONESTABLISHMENTREQUESTMessageIdentity): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, a.IntegrityProtectionMaximumDataRate.Octet[:]); err != nil {
		return fmt.Errorf("NAS decode error (PDUSessionEstablishmentRequest/IntegrityProtectionMaximumDataRate): %w", err)
	}
	for buffer.Len() > 0 {
		var ieiN uint8
		var tmpIeiN uint8
		if err := binary.Read(buffer, binary.BigEndian, &ieiN); err != nil {
			return fmt.Errorf("NAS decode error (PDUSessionEstablishmentRequest/iei): %w", err)
		}
		// fmt.Println(ieiN)
		if ieiN >= 0x80 {
			tmpIeiN = (ieiN & 0xf0) >> 4
		} else {
			tmpIeiN = ieiN
		}
		// fmt.Println("type", tmpIeiN)
		switch tmpIeiN {
		case PDUSessionEstablishmentRequestPDUSessionTypeType:
			a.PDUSessionType = nasType.NewPDUSessionType(ieiN)
			a.PDUSessionType.Octet = ieiN
		case PDUSessionEstablishmentRequestSSCModeType:
			a.SSCMode = nasType.NewSSCMode(ieiN)
			a.SSCMode.Octet = ieiN
		case PDUSessionEstablishmentRequestCapability5GSMType:
			a.Capability5GSM = nasType.NewCapability5GSM(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.Capability5GSM.Len); err != nil {
				return fmt.Errorf("NAS decode error (PDUSessionEstablishmentRequest/Capability5GSM): %w", err)
			}
			if a.Capability5GSM.Len < 1 || a.Capability5GSM.Len > 13 {
				return fmt.Errorf("invalid ie length (PDUSessionEstablishmentRequest/Capability5GSM): %d", a.Capability5GSM.Len)
			}
			a.Capability5GSM.SetLen(a.Capability5GSM.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.Capability5GSM.Octet[:a.Capability5GSM.GetLen()]); err != nil {
				return fmt.Errorf("NAS decode error (PDUSessionEstablishmentRequest/Capability5GSM): %w", err)
			}
		case PDUSessionEstablishmentRequestMaximumNumberOfSupportedPacketFiltersType:
			a.MaximumNumberOfSupportedPacketFilters = nasType.NewMaximumNumberOfSupportedPacketFilters(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, a.MaximumNumberOfSupportedPacketFilters.Octet[:]); err != nil {
				return fmt.Errorf("NAS decode error (PDUSessionEstablishmentRequest/MaximumNumberOfSupportedPacketFilters): %w", err)
			}
		case PDUSessionEstablishmentRequestAlwaysonPDUSessionRequestedType:
			a.AlwaysonPDUSessionRequested = nasType.NewAlwaysonPDUSessionRequested(ieiN)
			a.AlwaysonPDUSessionRequested.Octet = ieiN
		case PDUSessionEstablishmentRequestSMPDUDNRequestContainerType:
			a.SMPDUDNRequestContainer = nasType.NewSMPDUDNRequestContainer(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.SMPDUDNRequestContainer.Len); err != nil {
				return fmt.Errorf("NAS decode error (PDUSessionEstablishmentRequest/SMPDUDNRequestContainer): %w", err)
			}
			if a.SMPDUDNRequestContainer.Len < 1 {
				return fmt.Errorf("invalid ie length (PDUSessionEstablishmentRequest/SMPDUDNRequestContainer): %d", a.SMPDUDNRequestContainer.Len)
			}
			a.SMPDUDNRequestContainer.SetLen(a.SMPDUDNRequestContainer.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.SMPDUDNRequestContainer.Buffer); err != nil {
				return fmt.Errorf("NAS decode error (PDUSessionEstablishmentRequest/SMPDUDNRequestContainer): %w", err)
			}
		case PDUSessionEstablishmentRequestExtendedProtocolConfigurationOptionsType:
			a.ExtendedProtocolConfigurationOptions = nasType.NewExtendedProtocolConfigurationOptions(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolConfigurationOptions.Len); err != nil {
				return fmt.Errorf("NAS decode error (PDUSessionEstablishmentRequest/ExtendedProtocolConfigurationOptions): %w", err)
			}
			if a.ExtendedProtocolConfigurationOptions.Len < 1 {
				return fmt.Errorf("invalid ie length (PDUSessionEstablishmentRequest/ExtendedProtocolConfigurationOptions): %d", a.ExtendedProtocolConfigurationOptions.Len)
			}
			a.ExtendedProtocolConfigurationOptions.SetLen(a.ExtendedProtocolConfigurationOptions.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.ExtendedProtocolConfigurationOptions.Buffer); err != nil {
				return fmt.Errorf("NAS decode error (PDUSessionEstablishmentRequest/ExtendedProtocolConfigurationOptions): %w", err)
			}
		default:
		}
	}
	return nil
}
