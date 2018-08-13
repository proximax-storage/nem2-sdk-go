package crypto

import (
	"encoding/binary"
	"errors"
)

//Signature
type Signature struct {
	R []byte
	S []byte
}

// NewSignature R and S must fit into 32 bytes
func NewSignature(r []byte, s []byte) (*Signature, error) {
	if (len(r) != 32) || (len(s) != 32) {
		return nil, errors.New("binary signature representation of r and s must both have 32 bytes length")
	}
	ref := &Signature{r, s}
	return ref, nil
}

//NewSignatureFromBytes Creates a new signature from bytes array 64
func NewSignatureFromBytes(b []byte) (*Signature, error) {
	if len(b) < 64 {
		return nil, errors.New("binary signature representation must be 64 bytes")
	}
	ref := &Signature{b[:32], b[32:]}
	return ref, nil
}

/**
 * Gets the R-part of the signature.
 *
 * @return The R-part of the signature.
 */
func (ref *Signature) GetR() uint32 {

	return binary.BigEndian.Uint32(ref.R)
}

//GetS Gets the S-part of the signature.
func (ref *Signature) GetS() uint32 {

	return binary.BigEndian.Uint32(ref.S)
}

//getBytes Gets a little-endian 64-byte representation of the signature.
func (ref *Signature) getBytes() []byte {

	return append(ref.R, ref.S...)
}

func (ref *Signature) String() string {

	return string(ref.getBytes())
}
