package derangement

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsDerangedHappyPath(t *testing.T) {
	inStrings := []string{"a", "b", "c", "d"}
	outStrings := []string{"d", "c", "b", "a"}
	assert.True(t, isDeranged(inStrings, outStrings))
}

func TestIsDerangedSadPath(t *testing.T) {
	inStrings := []string{"a", "b", "c", "d"}
	outStrings := []string{"a", "d", "b", "c"}
	assert.False(t, isDeranged(inStrings, outStrings))
}
