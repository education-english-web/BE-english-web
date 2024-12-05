package payload

import (
	"context"
	"strings"

	"github.com/education-english-web/BE-english-web/pkg/transformer"
)

type RequestPayload interface {
	StructName() string
}

func Validate(p RequestPayload) error {

	return nil
}

// normalizeFieldName remove suffix [0] for a field type is slice error. Ex: Users[0] -> Users
func normalizeFieldName(name string) string {
	i := strings.Index(name, "[")

	if i > 0 {
		return name[:i]
	}

	return name
}

func transform(p RequestPayload) error {
	return transformer.GetInstance().Struct(context.Background(), p)
}
