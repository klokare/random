package random

import (
	"math/rand"
	"time"
)

var (
	// RandomSource function provides a new source for new random number generators
	RandomSource func() rand.Source
)

func init() {
	RandomSource = func() rand.Source { return rand.NewSource(time.Now().UnixNano()) }
}

// NewRandom returns a new random number generator using the source provided by the RandomSource function
func New() *rand.Rand {
	return rand.New(RandomSource())
}
