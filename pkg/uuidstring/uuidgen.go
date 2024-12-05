package uuidstring

//go:generate mockgen -destination=./mock/mock_$GOFILE -source=$GOFILE -package=mock

type UUIDString interface {
	GetUUID() string
	NewUUID() string
}
