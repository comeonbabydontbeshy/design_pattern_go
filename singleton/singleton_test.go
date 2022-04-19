package singleton

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLazyInstance(t *testing.T) {
	assert.Equal(t, GetLazyInstance(), GetLazyInstance())
}

func BenchmarkGetLazyInstance(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			if GetLazyInstance() != GetLazyInstance() {
				b.Error("test fail")
			}
		}
	})
}
