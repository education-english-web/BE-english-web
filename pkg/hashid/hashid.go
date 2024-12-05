package hashid

//go:generate mockgen -destination=./mock/mock_$GOFILE -source=$GOFILE -package=mock

import (
	"fmt"

	hashids "github.com/speps/go-hashids"

	appLog "github.com/education-english-web/BE-english-web/pkg/log"
)

var singleton IDHasher

type IDHasher interface {
	Encode(id uint32) string
	EncodeUint64(id uint64) string
	Decode(hashed string) (uint32, error)
	DecodeUint64(hashed string) (uint64, error)
}

type idHasher struct {
	hashID *hashids.HashID
}

// InitIDHasher initials an ID hasher
func InitIDHasher(minLength int, salt string) error {
	if singleton != nil {
		return nil
	}

	hashID, err := hashids.NewWithData(&hashids.HashIDData{
		Alphabet:  hashids.DefaultAlphabet,
		MinLength: minLength,
		Salt:      salt,
	})
	if err != nil {
		return fmt.Errorf("error while init hash ID: %w", err)
	}

	singleton = &idHasher{
		hashID: hashID,
	}

	return nil
}

// GetIDHasher Get returns the singleton instance of ID hasher
func GetIDHasher() IDHasher {
	return singleton
}

// Encode encodes a non-negative integer based ID to a string
func (h *idHasher) Encode(id uint32) string {
	hashed, err := h.hashID.Encode([]int{int(id)})
	if err != nil {
		appLog.Errorf("error while encoding id: %s", err.Error())
	}

	return hashed
}

func (h *idHasher) EncodeUint64(id uint64) string {
	hashed, err := h.hashID.EncodeInt64([]int64{int64(id)})
	if err != nil {
		appLog.Errorf("error while encoding id: %s", err.Error())
	}

	return hashed
}

// Decode decodes the hashed value to a non-negative integer
func (h *idHasher) Decode(hashed string) (uint32, error) {
	id, err := h.hashID.DecodeWithError(hashed)
	if err != nil {
		return 0, fmt.Errorf("error while decoding hashed id: %w", err)
	}

	if len(id) != 1 {
		return 0, fmt.Errorf("incorrect hash id format")
	}

	return uint32(id[0]), nil
}

func (h *idHasher) DecodeUint64(hashed string) (uint64, error) {
	id, err := h.hashID.DecodeInt64WithError(hashed)
	if err != nil {
		return 0, fmt.Errorf("error while decoding hashed id: %w", err)
	}

	if len(id) != 1 {
		return 0, fmt.Errorf("incorrect hash id format")
	}

	return uint64(id[0]), nil
}
