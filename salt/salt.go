package salt

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
)

var saltBasis = []uint32{0x89866356, 0x04011986, 0x09171989, 0x10071956,
	0x07151954, 0x05221963}

// Generate (un string) []byte a "salt" for a given username. The salt is only dependent on the
// username
func Generate(un string) []byte {
	unenc := make([]byte, 2*len(un))
	hex.Encode(unenc, []byte(un))

	buf := new(bytes.Buffer)
	for _, basis := range saltBasis {
		binary.Write(buf, binary.LittleEndian, basis)
	}
	binary.Write(buf, binary.LittleEndian, []byte(un))

	saltenc := make([]byte, 2*buf.Len())
	hex.Encode(saltenc, buf.Bytes())
	return saltenc
}
