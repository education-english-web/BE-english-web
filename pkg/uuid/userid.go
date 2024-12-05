package uuid

import (
	"crypto/md5"

	"github.com/google/uuid"
)

func CreatUserID(email string) string {
	hash := md5.New()
	hash.Write([]byte(email))
	return uuid.NewHash(hash, uuid.Nil, nil, 0).String()
}
