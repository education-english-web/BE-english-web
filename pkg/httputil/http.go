package httputil

import (
	"errors"
	"net"
	"net/http"
	"strings"
)

type httpFactory struct{}

// NewHTTPFactory initiates the http factory
func NewHTTPFactory() HTTPFactory {
	return &httpFactory{}
}

func (f *httpFactory) GetIP(r *http.Request) (string, error) {
	// Get IP from the X-REAL-IP header
	ip := r.Header.Get("X-REAL-IP")

	netIP := net.ParseIP(ip)
	if netIP != nil {
		return ip, nil
	}

	// Get IP from X-FORWARDED-FOR header
	ips := r.Header.Get("X-FORWARDED-FOR")
	splitIps := strings.Split(ips, ",")

	for _, ip := range splitIps {
		netIP = net.ParseIP(ip)
		if netIP != nil {
			return ip, nil
		}
	}

	// Get IP from RemoteAddr
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "", err
	}

	netIP = net.ParseIP(ip)
	if netIP != nil {
		return ip, nil
	}

	return "", errors.New("no valid ip found")
}

func (f *httpFactory) GetUserAgent(r *http.Request) string {
	return r.UserAgent()
}
