// #nosec
//
//nolint:gosec
package md5

import (
	"crypto/md5"

	"github.com/education-english-web/BE-english-web/pkg/hashing"
)

type md5Hashing struct{}

func New() hashing.Hashing {
	return &md5Hashing{}
}

func (h *md5Hashing) Hash(input string) []byte {
	b := md5.Sum([]byte(input))

	return b[:]
}
