package interfaces

//go:generate mockgen -destination=./mock/mock_$GOFILE -source=$GOFILE -package=mock

type SignatureVerifier interface {
	Verify(message, signature string) error
}

type IPVerifier interface {
	Verify(ip string) error
}
