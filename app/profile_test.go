package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckProfile(t *testing.T) {
	tests := []struct {
		name   string
		prefix string
	}{
		{"CheckProfile", "./test/profile"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CheckProfile(tt.prefix)
			assert.NoError(t, err)
		})
	}
}
