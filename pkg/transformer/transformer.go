package transformer

import (
	"github.com/go-playground/mold/v4"
	"github.com/go-playground/mold/v4/modifiers"
)

var globalTransformer *mold.Transformer

func init() {
	globalTransformer = modifiers.New()
}

func GetInstance() *mold.Transformer {
	return globalTransformer
}
