package partnerverifier

import (
	"sync"

	"github.com/education-english-web/BE-english-web/pkg/partnerverifier/interfaces"
)

//go:generate mockgen -destination=./mock/mock_$GOFILE -source=$GOFILE -package=mock

type PartnerName string

const (
	PartnerNameKitAlive PartnerName = "KITALIVE"
	PartnerNamePayable  PartnerName = "PAYABLE"
)

var (
	singleton *partnerVerifier
	once      sync.Once
)

type PartnerVerifier interface {
	GetVerifier(partnerName string) Verifier
}

type partnerVerifier struct {
	m map[PartnerName]Verifier
}

type Verifier struct {
	SigVerifier interfaces.SignatureVerifier
	IPVerifier  interfaces.IPVerifier
}

func InitPartnerVerifiers(mVerifers map[PartnerName]Verifier) PartnerVerifier {
	once.Do(func() {
		singleton = &partnerVerifier{
			m: mVerifers,
		}
	})

	return singleton
}

func (v *partnerVerifier) GetVerifier(partner string) Verifier {
	var partnerName PartnerName

	switch partner {
	case "kitalive":
		partnerName = PartnerNameKitAlive
	case "payable":
		partnerName = PartnerNamePayable
	default:
		partnerName = ""
	}

	return v.m[partnerName]
}
