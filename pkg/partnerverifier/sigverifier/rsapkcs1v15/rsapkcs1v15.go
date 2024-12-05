package rsapkcs1v15

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"

	"github.com/education-english-web/BE-english-web/pkg/partnerverifier/interfaces"
)

type rsapkcs1v15 struct {
	publicKey *rsa.PublicKey
}

func New(publicKeyBase64Str string) (interfaces.SignatureVerifier, error) {
	publicKeyStr, err := base64.StdEncoding.DecodeString(publicKeyBase64Str)
	if err != nil {
		return nil, fmt.Errorf("b64 decode publicKey: %w", err)
	}

	pubBlock, _ := pem.Decode(publicKeyStr)
	if pubBlock == nil {
		return nil, fmt.Errorf("pem decode public key string")
	}

	pubKByte, err := x509.ParsePKIXPublicKey(pubBlock.Bytes)
	if err != nil {
		return nil, fmt.Errorf("parse public key: %w", err)
	}

	pubKey, ok := pubKByte.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("failed public key type assertion")
	}

	return &rsapkcs1v15{
		publicKey: pubKey,
	}, nil
}

func (v *rsapkcs1v15) Verify(message, signature string) error {
	rawSignature, _ := base64.StdEncoding.DecodeString(signature)
	hashed := sha256.Sum256([]byte(message))

	return rsa.VerifyPKCS1v15(v.publicKey, crypto.SHA256, hashed[:], rawSignature)
}
