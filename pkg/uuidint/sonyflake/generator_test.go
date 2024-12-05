package sonyflake

import (
	"testing"

	"github.com/sony/sonyflake"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	t.Run("uuidflake.Get success", func(t *testing.T) {
		_ = Setup()

		var st sonyflake.Settings
		sf := sonyflake.NewSonyflake(st)

		want := &generator{
			sf: sf,
		}
		got := Get()

		assert.Equal(t, want, got)
	})
}

func Test_generator_nextid(t *testing.T) {
	_ = Setup()

	got := Get()
	_, _ = got.NextID()
}
