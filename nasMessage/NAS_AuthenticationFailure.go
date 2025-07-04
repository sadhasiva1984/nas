// Code generated by generate.sh, DO NOT EDIT.

package nasMessage

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/sadhasiva1984/nas/nasType"
)

type AuthenticationFailure struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.SpareHalfOctetAndSecurityHeaderType
	nasType.AuthenticationFailureMessageIdentity
	nasType.Cause5GMM
	*nasType.AuthenticationFailureParameter
}

func NewAuthenticationFailure(iei uint8) (authenticationFailure *AuthenticationFailure) {
	authenticationFailure = &AuthenticationFailure{}
	return authenticationFailure
}

const (
	AuthenticationFailureAuthenticationFailureParameterType uint8 = 0x30
)

func (a *AuthenticationFailure) EncodeAuthenticationFailure(buffer *bytes.Buffer) error {
	if err := binary.Write(buffer, binary.BigEndian, a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return fmt.Errorf("NAS encode error (AuthenticationFailure/ExtendedProtocolDiscriminator): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return fmt.Errorf("NAS encode error (AuthenticationFailure/SpareHalfOctetAndSecurityHeaderType): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.AuthenticationFailureMessageIdentity.Octet); err != nil {
		return fmt.Errorf("NAS encode error (AuthenticationFailure/AuthenticationFailureMessageIdentity): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.Cause5GMM.Octet); err != nil {
		return fmt.Errorf("NAS encode error (AuthenticationFailure/Cause5GMM): %w", err)
	}
	if a.AuthenticationFailureParameter != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.AuthenticationFailureParameter.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (AuthenticationFailure/AuthenticationFailureParameter): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.AuthenticationFailureParameter.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (AuthenticationFailure/AuthenticationFailureParameter): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.AuthenticationFailureParameter.Octet[:]); err != nil {
			return fmt.Errorf("NAS encode error (AuthenticationFailure/AuthenticationFailureParameter): %w", err)
		}
	}
	return nil
}

func (a *AuthenticationFailure) DecodeAuthenticationFailure(byteArray *[]byte) error {
	buffer := bytes.NewBuffer(*byteArray)
	if err := binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return fmt.Errorf("NAS decode error (AuthenticationFailure/ExtendedProtocolDiscriminator): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return fmt.Errorf("NAS decode error (AuthenticationFailure/SpareHalfOctetAndSecurityHeaderType): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.AuthenticationFailureMessageIdentity.Octet); err != nil {
		return fmt.Errorf("NAS decode error (AuthenticationFailure/AuthenticationFailureMessageIdentity): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.Cause5GMM.Octet); err != nil {
		return fmt.Errorf("NAS decode error (AuthenticationFailure/Cause5GMM): %w", err)
	}
	for buffer.Len() > 0 {
		var ieiN uint8
		var tmpIeiN uint8
		if err := binary.Read(buffer, binary.BigEndian, &ieiN); err != nil {
			return fmt.Errorf("NAS decode error (AuthenticationFailure/iei): %w", err)
		}
		// fmt.Println(ieiN)
		if ieiN >= 0x80 {
			tmpIeiN = (ieiN & 0xf0) >> 4
		} else {
			tmpIeiN = ieiN
		}
		// fmt.Println("type", tmpIeiN)
		switch tmpIeiN {
		case AuthenticationFailureAuthenticationFailureParameterType:
			a.AuthenticationFailureParameter = nasType.NewAuthenticationFailureParameter(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.AuthenticationFailureParameter.Len); err != nil {
				return fmt.Errorf("NAS decode error (AuthenticationFailure/AuthenticationFailureParameter): %w", err)
			}
			if a.AuthenticationFailureParameter.Len != 14 {
				return fmt.Errorf("invalid ie length (AuthenticationFailure/AuthenticationFailureParameter): %d", a.AuthenticationFailureParameter.Len)
			}
			a.AuthenticationFailureParameter.SetLen(a.AuthenticationFailureParameter.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.AuthenticationFailureParameter.Octet[:]); err != nil {
				return fmt.Errorf("NAS decode error (AuthenticationFailure/AuthenticationFailureParameter): %w", err)
			}
		default:
		}
	}
	return nil
}
