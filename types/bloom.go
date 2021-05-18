package types

import (
	"encoding/binary"
	"sync"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gogo/protobuf/jsonpb"
	"golang.org/x/crypto/sha3"
)

const (
	BloomByteLength = 256
)

// hasherPool holds LegacyKeccak256 hashers for rlpHash.
var hasherPool = sync.Pool{
	New: func() interface{} { return sha3.NewLegacyKeccak256() },
}

type Bloom [BloomByteLength]byte

func (b *Bloom) ProtoMessage() {}

func (b *Bloom) Size() int {
	return BloomByteLength
}

func (b *Bloom) MarshalTo(data []byte) (int, error) {
	data = data[:b.Size()]
	copy(data, b[:])

	return b.Size(), nil
}

func (b *Bloom) Unmarshal(data []byte) error {
	if len(data) > BloomByteLength {
		data = data[len(data)-BloomByteLength:]
	}

	copy(b[:], data)

	return nil
}

func (b *Bloom) MarshalJSONPB(m *jsonpb.Marshaler) ([]byte, error) {
	return b.MarshalJSON()
}

func (b *Bloom) UnmarshalJSONPB(m *jsonpb.Unmarshaler, data []byte) error {
	return b.UnmarshalJSON(data)
}

// Serialize given bloom to JSON
func (b *Bloom) MarshalJSON() ([]byte, error) {
	return b[:], nil
}

// UnmarshalJSON parses a bloom
func (b *Bloom) UnmarshalJSON(data []byte) error {
	copy(b[:], data)

	return nil
}

// Add adds d to the filter. Future calls of Test(d) will return true.
func (b *Bloom) Add(d []byte) {
	b.add(d, make([]byte, 6))
}

// add is internal version of Add, which takes a scratch buffer for reuse (needs to be at least 6 bytes)
func (b *Bloom) add(d []byte, buf []byte) {
	i1, v1, i2, v2, i3, v3 := bloomValues(d, buf)
	b[i1] |= v1
	b[i2] |= v2
	b[i3] |= v3
}

// MarshalText encodes b as a hex string with 0x prefix.
func (b Bloom) MarshalText() ([]byte, error) {
	return hexutil.Bytes(b[:]).MarshalText()
}

// UnmarshalText b as a hex string with 0x prefix.
func (b *Bloom) UnmarshalText(input []byte) error {
	return hexutil.UnmarshalFixedText("Bloom", input, b[:])
}

// Test checks if the given topic is present in the bloom filter
func (b Bloom) Test(topic []byte) bool {
	i1, v1, i2, v2, i3, v3 := bloomValues(topic, make([]byte, 6))
	return v1 == v1&b[i1] &&
		v2 == v2&b[i2] &&
		v3 == v3&b[i3]
}

// bloomValues returns the bytes (index-value pairs) to set for the given data
func bloomValues(data []byte, hashbuf []byte) (uint, byte, uint, byte, uint, byte) {
	sha := hasherPool.Get().(crypto.KeccakState)
	sha.Reset()
	sha.Write(data)
	sha.Read(hashbuf)
	hasherPool.Put(sha)
	// The actual bits to flip
	v1 := byte(1 << (hashbuf[1] & 0x7))
	v2 := byte(1 << (hashbuf[3] & 0x7))
	v3 := byte(1 << (hashbuf[5] & 0x7))
	// The indices for the bytes to OR in
	i1 := BloomByteLength - uint((binary.BigEndian.Uint16(hashbuf)&0x7ff)>>3) - 1
	i2 := BloomByteLength - uint((binary.BigEndian.Uint16(hashbuf[2:])&0x7ff)>>3) - 1
	i3 := BloomByteLength - uint((binary.BigEndian.Uint16(hashbuf[4:])&0x7ff)>>3) - 1

	return i1, v1, i2, v2, i3, v3
}
