package utils

import (
	"math/rand"
	"time"
)

func RandomUint32() uint32 {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return rand.Uint32()
}

// rand.Seed is deprecated: As of Go 1.20 there is no reason to call Seed with a random value. Programs that call Seed with a known value to get a specific sequence of results should use New(NewSource(seed)) to obtain a local random generator.deprecateddefault
