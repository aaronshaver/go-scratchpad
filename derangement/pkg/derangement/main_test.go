package derangement

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var inStrings = []string{"a", "b", "c", "d"}

func TestIsDerangedHappyPath(t *testing.T) {
	outStrings := []string{"d", "c", "b", "a"}
	assert.True(t, isDeranged(inStrings, outStrings))
}

func TestIsDerangedSadPath(t *testing.T) {
	outStrings := []string{"a", "d", "b", "c"}
	assert.False(t, isDeranged(inStrings, outStrings))
}

func TestIsDerangedHappyPath2(t *testing.T) {
	customInStrings := []string{"a", "b", "c", "d"}
	outStrings := []string{"b", "d", "a", "c"}
	assert.True(t, isDeranged(customInStrings, outStrings))
}

func TestDerangeHappyPath(t *testing.T) {
	originalStrings := append([]string(nil), inStrings...)
	for i := 0; i < 1000; i++ { // since derange can be correct by pure chance, test it many times
		tempStrings := append([]string(nil), originalStrings...)
		outStrings := derange(tempStrings)
		assert.True(t, isDeranged(originalStrings, outStrings))
	}
}
