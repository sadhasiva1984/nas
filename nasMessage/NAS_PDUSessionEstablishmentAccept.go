// Code generated by generate.sh, DO NOT EDIT.

package nasMessage

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/sadhasiva1984/nas/nasType"
)

type PDUSessionEstablishmentAccept struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.PDUSessionID
	nasType.PTI
	nasType.PDUSESSIONESTABLISHMENTACCEPTMessageIdentity
	nasType.SelectedSSCModeAndSelectedPDUSessionType
	nasType.AuthorizedQosRules
	nasType.SessionAMBR
	*nasType.Cause5GSM
	*nasType.PDUAddress
	*nasType.RQTimerValue
	*nasType.SNSSAI
	*nasType.AlwaysonPDUSessionIndication
	*nasType.MappedEPSBearerContexts
	*nasType.EAPMessage
	*nasType.AuthorizedQosFlowDescriptions
	*nasType.ExtendedProtocolConfigurationOptions
	*nasType.DNN
}

func NewPDUSessionEstablishmentAccept(iei uint8) (pDUSessionEstablishmentAccept *PDUSessionEstablishmentAccept) {
	pDUSessionEstablishmentAccept = &PDUSessionEstablishmentAccept{}
	return pDUSessionEstablishmentAccept
}

const (
	PDUSessionEstablishmentAcceptCause5GSMType                            uint8 = 0x59
	PDUSessionEstablishmentAcceptPDUAddressType                           uint8 = 0x29
	PDUSessionEstablishmentAcceptRQTimerValueType                         uint8 = 0x56
	PDUSessionEstablishmentAcceptSNSSAIType                               uint8 = 0x22
	PDUSessionEstablishmentAcceptAlwaysonPDUSessionIndicationType         uint8 = 0x08
	PDUSessionEstablishmentAcceptMappedEPSBearerContextsType              uint8 = 0x75
	PDUSessionEstablishmentAcceptEAPMessageType                           uint8 = 0x78
	PDUSessionEstablishmentAcceptAuthorizedQosFlowDescriptionsType        uint8 = 0x79
	PDUSessionEstablishmentAcceptExtendedProtocolConfigurationOptionsType uint8 = 0x7B
	PDUSessionEstablishmentAcceptDNNType                                  uint8 = 0x25
)

func (a *PDUSessionEstablishmentAccept) EncodePDUSessionEstablishmentAccept(buffer *bytes.Buffer) error {
	if err := binary.Write(buffer, binary.BigEndian, a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return fmt.Errorf("NAS encode error (PDUSessionEstablishmentAccept/ExtendedProtocolDiscriminator): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.PDUSessionID.Octet); err != nil {
		return fmt.Errorf("NAS encode error (PDUSessionEstablishmentAccept/PDUSessionID): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.PTI.Octet); err != nil {
		return fmt.Errorf("NAS encode error (PDUSessionEstablishmentAccept/PTI): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.PDUSESSIONESTABLISHMENTACCEPTMessageIdentity.Octet); err != nil {
		return fmt.Errorf("NAS encode error (PDUSessionEstablishmentAccept/PDUSESSIONESTABLISHMENTACCEPTMessageIdentity): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.SelectedSSCModeAndSelectedPDUSessionType.Octet); err != nil {
		return fmt.Errorf("NAS encode error (PDUSessionEstablishmentAccept/SelectedSSCModeAndSelectedPDUSessionType): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.AuthorizedQosRules.GetLen()); err != nil {
		return fmt.Errorf("NAS encode error (PDUSessionEstablishmentAccept/AuthorizedQosRules): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.AuthorizedQosRules.Buffer); err != nil {
		return fmt.Errorf("NAS encode error (PDUSessionEstablishmentAccept/AuthorizedQosRules): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.SessionAMBR.GetLen()); err != nil {
		return fmt.Errorf("NAS encode error (PDUSessionEstablishmentAccept/SessionAMBR): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.SessionAMBR.Octet[:]); err != nil {
		return fmt.Errorf("NAS encode error (PDUSessionEstablishmentAccept/SessionAMBR): %w", err)
	}
	if a.Cause5GSM != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.Cause5GSM.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionEstablishmentAccept/Cause5GSM): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.Cause5GSM.Octet); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionEstablishmentAccept/Cause5GSM): %w", err)
		}
	}
	if a.PDUAddress != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.PDUAddress.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionEstablishmentAccept/PDUAddress): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.PDUAddress.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionEstablishmentAccept/PDUAddress): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.PDUAddress.Octet[:a.PDUAddress.GetLen()]); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionEstablishmentAccept/PDUAddress): %w", err)
		}
	}
	if a.RQTimerValue != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.RQTimerValue.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionEstablishmentAccept/RQTimerValue): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.RQTimerValue.Octet); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionEstablishmentAccept/RQTimerValue): %w", err)
		}
	}
	if a.SNSSAI != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.SNSSAI.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionEstablishmentAccept/SNSSAI): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.SNSSAI.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionEstablishmentAccept/SNSSAI): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.SNSSAI.Octet[:a.SNSSAI.GetLen()]); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionEstablishmentAccept/SNSSAI): %w", err)
		}
	}
	if a.AlwaysonPDUSessionIndication != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.AlwaysonPDUSessionIndication.Octet); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionEstablishmentAccept/AlwaysonPDUSessionIndication): %w", err)
		}
	}
	if a.MappedEPSBearerContexts != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.MappedEPSBearerContexts.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionEstablishmentAccept/MappedEPSBearerContexts): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.MappedEPSBearerContexts.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionEstablishmentAccept/MappedEPSBearerContexts): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.MappedEPSBearerContexts.Buffer); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionEstablishmentAccept/MappedEPSBearerContexts): %w", err)
		}
	}
	if a.EAPMessage != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.EAPMessage.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionEstablishmentAccept/EAPMessage): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.EAPMessage.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionEstablishmentAccept/EAPMessage): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.EAPMessage.Buffer); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionEstablishmentAccept/EAPMessage): %w", err)
		}
	}
	if a.AuthorizedQosFlowDescriptions != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.AuthorizedQosFlowDescriptions.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionEstablishmentAccept/AuthorizedQosFlowDescriptions): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.AuthorizedQosFlowDescriptions.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionEstablishmentAccept/AuthorizedQosFlowDescriptions): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.AuthorizedQosFlowDescriptions.Buffer); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionEstablishmentAccept/AuthorizedQosFlowDescriptions): %w", err)
		}
	}
	if a.ExtendedProtocolConfigurationOptions != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.ExtendedProtocolConfigurationOptions.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionEstablishmentAccept/ExtendedProtocolConfigurationOptions): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.ExtendedProtocolConfigurationOptions.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionEstablishmentAccept/ExtendedProtocolConfigurationOptions): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.ExtendedProtocolConfigurationOptions.Buffer); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionEstablishmentAccept/ExtendedProtocolConfigurationOptions): %w", err)
		}
	}
	if a.DNN != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.DNN.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionEstablishmentAccept/DNN): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.DNN.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionEstablishmentAccept/DNN): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.DNN.Buffer); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionEstablishmentAccept/DNN): %w", err)
		}
	}
	return nil
}

func (a *PDUSessionEstablishmentAccept) DecodePDUSessionEstablishmentAccept(byteArray *[]byte) error {
	buffer := bytes.NewBuffer(*byteArray)
	if err := binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return fmt.Errorf("NAS decode error (PDUSessionEstablishmentAccept/ExtendedProtocolDiscriminator): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.PDUSessionID.Octet); err != nil {
		return fmt.Errorf("NAS decode error (PDUSessionEstablishmentAccept/PDUSessionID): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.PTI.Octet); err != nil {
		return fmt.Errorf("NAS decode error (PDUSessionEstablishmentAccept/PTI): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.PDUSESSIONESTABLISHMENTACCEPTMessageIdentity.Octet); err != nil {
		return fmt.Errorf("NAS decode error (PDUSessionEstablishmentAccept/PDUSESSIONESTABLISHMENTACCEPTMessageIdentity): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.SelectedSSCModeAndSelectedPDUSessionType.Octet); err != nil {
		return fmt.Errorf("NAS decode error (PDUSessionEstablishmentAccept/SelectedSSCModeAndSelectedPDUSessionType): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.AuthorizedQosRules.Len); err != nil {
		return fmt.Errorf("NAS decode error (PDUSessionEstablishmentAccept/AuthorizedQosRules): %w", err)
	}
	if a.AuthorizedQosRules.Len < 4 {
		return fmt.Errorf("invalid ie length (PDUSessionEstablishmentAccept/AuthorizedQosRules): %d", a.AuthorizedQosRules.Len)
	}
	a.AuthorizedQosRules.SetLen(a.AuthorizedQosRules.GetLen())
	if err := binary.Read(buffer, binary.BigEndian, a.AuthorizedQosRules.Buffer); err != nil {
		return fmt.Errorf("NAS decode error (PDUSessionEstablishmentAccept/AuthorizedQosRules): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.SessionAMBR.Len); err != nil {
		return fmt.Errorf("NAS decode error (PDUSessionEstablishmentAccept/SessionAMBR): %w", err)
	}
	if a.SessionAMBR.Len != 6 {
		return fmt.Errorf("invalid ie length (PDUSessionEstablishmentAccept/SessionAMBR): %d", a.SessionAMBR.Len)
	}
	a.SessionAMBR.SetLen(a.SessionAMBR.GetLen())
	if err := binary.Read(buffer, binary.BigEndian, a.SessionAMBR.Octet[:]); err != nil {
		return fmt.Errorf("NAS decode error (PDUSessionEstablishmentAccept/SessionAMBR): %w", err)
	}
	for buffer.Len() > 0 {
		var ieiN uint8
		var tmpIeiN uint8
		if err := binary.Read(buffer, binary.BigEndian, &ieiN); err != nil {
			return fmt.Errorf("NAS decode error (PDUSessionEstablishmentAccept/iei): %w", err)
		}
		// fmt.Println(ieiN)
		if ieiN >= 0x80 {
			tmpIeiN = (ieiN & 0xf0) >> 4
		} else {
			tmpIeiN = ieiN
		}
		// fmt.Println("type", tmpIeiN)
		switch tmpIeiN {
		case PDUSessionEstablishmentAcceptCause5GSMType:
			a.Cause5GSM = nasType.NewCause5GSM(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.Cause5GSM.Octet); err != nil {
				return fmt.Errorf("NAS decode error (PDUSessionEstablishmentAccept/Cause5GSM): %w", err)
			}
		case PDUSessionEstablishmentAcceptPDUAddressType:
			a.PDUAddress = nasType.NewPDUAddress(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.PDUAddress.Len); err != nil {
				return fmt.Errorf("NAS decode error (PDUSessionEstablishmentAccept/PDUAddress): %w", err)
			}
			if a.PDUAddress.Len != 5 && a.PDUAddress.Len != 9 && a.PDUAddress.Len != 13 {
				return fmt.Errorf("invalid ie length (PDUSessionEstablishmentAccept/PDUAddress): %d", a.PDUAddress.Len)
			}
			a.PDUAddress.SetLen(a.PDUAddress.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.PDUAddress.Octet[:a.PDUAddress.GetLen()]); err != nil {
				return fmt.Errorf("NAS decode error (PDUSessionEstablishmentAccept/PDUAddress): %w", err)
			}
		case PDUSessionEstablishmentAcceptRQTimerValueType:
			a.RQTimerValue = nasType.NewRQTimerValue(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.RQTimerValue.Octet); err != nil {
				return fmt.Errorf("NAS decode error (PDUSessionEstablishmentAccept/RQTimerValue): %w", err)
			}
		case PDUSessionEstablishmentAcceptSNSSAIType:
			a.SNSSAI = nasType.NewSNSSAI(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.SNSSAI.Len); err != nil {
				return fmt.Errorf("NAS decode error (PDUSessionEstablishmentAccept/SNSSAI): %w", err)
			}
			if a.SNSSAI.Len < 1 || a.SNSSAI.Len > 8 {
				return fmt.Errorf("invalid ie length (PDUSessionEstablishmentAccept/SNSSAI): %d", a.SNSSAI.Len)
			}
			a.SNSSAI.SetLen(a.SNSSAI.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.SNSSAI.Octet[:a.SNSSAI.GetLen()]); err != nil {
				return fmt.Errorf("NAS decode error (PDUSessionEstablishmentAccept/SNSSAI): %w", err)
			}
		case PDUSessionEstablishmentAcceptAlwaysonPDUSessionIndicationType:
			a.AlwaysonPDUSessionIndication = nasType.NewAlwaysonPDUSessionIndication(ieiN)
			a.AlwaysonPDUSessionIndication.Octet = ieiN
		case PDUSessionEstablishmentAcceptMappedEPSBearerContextsType:
			a.MappedEPSBearerContexts = nasType.NewMappedEPSBearerContexts(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.MappedEPSBearerContexts.Len); err != nil {
				return fmt.Errorf("NAS decode error (PDUSessionEstablishmentAccept/MappedEPSBearerContexts): %w", err)
			}
			if a.MappedEPSBearerContexts.Len < 4 {
				return fmt.Errorf("invalid ie length (PDUSessionEstablishmentAccept/MappedEPSBearerContexts): %d", a.MappedEPSBearerContexts.Len)
			}
			a.MappedEPSBearerContexts.SetLen(a.MappedEPSBearerContexts.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.MappedEPSBearerContexts.Buffer); err != nil {
				return fmt.Errorf("NAS decode error (PDUSessionEstablishmentAccept/MappedEPSBearerContexts): %w", err)
			}
		case PDUSessionEstablishmentAcceptEAPMessageType:
			a.EAPMessage = nasType.NewEAPMessage(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.EAPMessage.Len); err != nil {
				return fmt.Errorf("NAS decode error (PDUSessionEstablishmentAccept/EAPMessage): %w", err)
			}
			if a.EAPMessage.Len < 4 || a.EAPMessage.Len > 1500 {
				return fmt.Errorf("invalid ie length (PDUSessionEstablishmentAccept/EAPMessage): %d", a.EAPMessage.Len)
			}
			a.EAPMessage.SetLen(a.EAPMessage.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.EAPMessage.Buffer); err != nil {
				return fmt.Errorf("NAS decode error (PDUSessionEstablishmentAccept/EAPMessage): %w", err)
			}
		case PDUSessionEstablishmentAcceptAuthorizedQosFlowDescriptionsType:
			a.AuthorizedQosFlowDescriptions = nasType.NewAuthorizedQosFlowDescriptions(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.AuthorizedQosFlowDescriptions.Len); err != nil {
				return fmt.Errorf("NAS decode error (PDUSessionEstablishmentAccept/AuthorizedQosFlowDescriptions): %w", err)
			}
			if a.AuthorizedQosFlowDescriptions.Len < 3 {
				return fmt.Errorf("invalid ie length (PDUSessionEstablishmentAccept/AuthorizedQosFlowDescriptions): %d", a.AuthorizedQosFlowDescriptions.Len)
			}
			a.AuthorizedQosFlowDescriptions.SetLen(a.AuthorizedQosFlowDescriptions.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.AuthorizedQosFlowDescriptions.Buffer); err != nil {
				return fmt.Errorf("NAS decode error (PDUSessionEstablishmentAccept/AuthorizedQosFlowDescriptions): %w", err)
			}
		case PDUSessionEstablishmentAcceptExtendedProtocolConfigurationOptionsType:
			a.ExtendedProtocolConfigurationOptions = nasType.NewExtendedProtocolConfigurationOptions(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolConfigurationOptions.Len); err != nil {
				return fmt.Errorf("NAS decode error (PDUSessionEstablishmentAccept/ExtendedProtocolConfigurationOptions): %w", err)
			}
			if a.ExtendedProtocolConfigurationOptions.Len < 1 {
				return fmt.Errorf("invalid ie length (PDUSessionEstablishmentAccept/ExtendedProtocolConfigurationOptions): %d", a.ExtendedProtocolConfigurationOptions.Len)
			}
			a.ExtendedProtocolConfigurationOptions.SetLen(a.ExtendedProtocolConfigurationOptions.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.ExtendedProtocolConfigurationOptions.Buffer); err != nil {
				return fmt.Errorf("NAS decode error (PDUSessionEstablishmentAccept/ExtendedProtocolConfigurationOptions): %w", err)
			}
		case PDUSessionEstablishmentAcceptDNNType:
			a.DNN = nasType.NewDNN(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.DNN.Len); err != nil {
				return fmt.Errorf("NAS decode error (PDUSessionEstablishmentAccept/DNN): %w", err)
			}
			if a.DNN.Len < 1 || a.DNN.Len > 100 {
				return fmt.Errorf("invalid ie length (PDUSessionEstablishmentAccept/DNN): %d", a.DNN.Len)
			}
			a.DNN.SetLen(a.DNN.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.DNN.Buffer); err != nil {
				return fmt.Errorf("NAS decode error (PDUSessionEstablishmentAccept/DNN): %w", err)
			}
		default:
		}
	}
	return nil
}
