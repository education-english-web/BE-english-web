package ulid

import (
	"crypto/rand"
	"time"

	ulid_v2 "github.com/oklog/ulid/v2"

	"github.com/education-english-web/BE-english-web/pkg/uuidstring"
)

type ulid struct{}

// New return provider for ulid uuid
func New() uuidstring.UUIDString {
	return &ulid{}
}

func (h ulid) GetUUID() string {
	entropy := ulid_v2.Monotonic(rand.Reader, 0)
	id := ulid_v2.MustNew(ulid_v2.Timestamp(time.Now()), entropy)

	return id.String()
}

func (h ulid) NewUUID() string {
	entropy := ulid_v2.Monotonic(rand.Reader, 0)
	id := ulid_v2.MustNew(ulid_v2.Timestamp(time.Now()), entropy)

	return id.String()
}
