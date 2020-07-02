package cmrand

import (
	crand "crypto/rand"
	"encoding/binary"
	"fmt"
	mrand "math/rand"
)

// CrSource is random source using crypto/rand default generator
type CrSource struct{}

// New returns *math.Rand initializes with crypto/rand source
func New() *mrand.Rand {
	return mrand.New(NewSource(1))
}

// NewSource returns math.rand compatible source using crypto/rand's Read() to get the numbers
func NewSource(int64) CrSource {
	return CrSource{}
}

//Seed is dummy implementation to satisfy interface math/rand.Source interafce
func (s CrSource) Seed(int64) {}

// Int63 returns a non-negative 63 bit integer sourced from crypto/rand default source, or panics.
func (s CrSource) Int63() int64 {
	var uv uint64
	err := binary.Read(crand.Reader, binary.LittleEndian, &uv)
	if err != nil {
		panic(fmt.Sprintf("could not read random data: %s", err))
	}
	return int64(uv & ^uint64(1<<63))
}
