package hashing

//go:generate mockgen -destination=./mock/mock_$GOFILE -source=$GOFILE -package=mock

type Hashing interface {
	Hash(input string) (output []byte)
}
