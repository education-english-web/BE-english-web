package sonyflake

import (
	"errors"
	"sync"

	"github.com/sony/sonyflake"

	"github.com/education-english-web/BE-english-web/pkg/uuidint"
)

type generator struct {
	sf *sonyflake.Sonyflake
}

var (
	once     sync.Once
	instance *generator
	errInit  error
)

func Setup() error {
	once.Do(func() {
		var st sonyflake.Settings

		sf := sonyflake.NewSonyflake(st)

		if sf == nil {
			errInit = errors.New("sonyflake are not created")
		}

		instance = &generator{
			sf: sf,
		}
	})

	return errInit
}

func Get() uuidint.UUIDInt {
	return instance
}

func (g *generator) NextID() (uint64, error) {
	return g.sf.NextID()
}
