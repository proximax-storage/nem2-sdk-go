package crypto

import (
	"encoding/binary"
	"errors"
	"math/big"
)

//Signature
type Signature struct {
	R []byte
	S []byte
}

var (
	errBadParamNewSignature          = errors.New("binary signature representation of r and s must both have 32 bytes length")
	errBadParamNewSignatureBigInt    = errors.New("bad parameters NewSignatureFromBigInt")
	errBadParamNewSignatureFromBytes = errors.New("binary signature representation must be 64 bytes")
)

// NewSignature R and S must fit into 32 bytes
func NewSignature(r []byte, s []byte) (*Signature, error) {
	if (len(r) != 32) || (len(s) != 32) {
		return nil, errBadParamNewSignature
	}
	ref := &Signature{r, s}
	return ref, nil
}
func NewSignatureFromBigInt(rInt, sInt *big.Int) (*Signature, error) {
	if (rInt == nil) || (sInt == nil) {
		return nil, errBadParamNewSignatureBigInt
	}

	var r, s [32]byte
	copy(r[:], rInt.Bytes())
	copy(s[:], sInt.Bytes())

	return NewSignature(r[:], s[:])
}

//NewSignatureFromBytes Creates a new signature from bytes array 64
func NewSignatureFromBytes(b []byte) (*Signature, error) {
	if len(b) < 64 {
		return nil, errBadParamNewSignatureFromBytes
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

//Bytes Gets a little-endian 64-byte representation of the signature.
func (ref *Signature) Bytes() []byte {

	return append(ref.R, ref.S...)
}

func (ref *Signature) String() string {

	return string(ref.Bytes())
}
