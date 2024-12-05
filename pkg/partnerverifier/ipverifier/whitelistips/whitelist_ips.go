package whitelistips

import (
	"errors"
	"slices"

	"github.com/education-english-web/BE-english-web/pkg/partnerverifier/interfaces"
)

type whitelistIPs struct {
	ips []string
}

func New(ips []string) interfaces.IPVerifier {
	return &whitelistIPs{ips: ips}
}

func (v *whitelistIPs) Verify(ip string) error {
	if len(v.ips) > 0 && v.ips[0] != "*" &&
		!slices.Contains(v.ips, ip) {
		return errors.New("invalid request ip source")
	}

	return nil
}
