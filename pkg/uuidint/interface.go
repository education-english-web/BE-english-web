package uuidint

//go:generate mockgen -destination=./mock/mock_$GOFILE -source=$GOFILE -package=mock
type UUIDInt interface {
	NextID() (uint64, error)
}
