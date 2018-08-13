package crypto

import (
	"encoding/binary"
	"errors"
)

type Signature struct {
	R []byte // private
	S []byte // private
	/**
	 * Creates a new signature.
	 *
	 * @param R The R-part of the signature.
	 * @param S The S-part of the signature.
	 */
} /* Signature */
// NewSignature R and S must fit into 32 bytes
func NewSignature(r []byte, s []byte) (*Signature, error) {
	if (len(r) != 32) || (len(s) != 32) {
		return nil, errors.New("binary signature representation of r and s must both have 32 bytes length")
	}
	ref := &Signature{r, s}
	return ref, nil
}

/**
 * Creates a new signature.
 *
 * @param bytes The binary representation of the signature.
 */
func NewSignatureFromBytes(b []byte) (*Signature, error) {
	if len(b) < 64 {
		return nil, errors.New("binary signature representation must be 64 bytes")
	}
	ref := &Signature{b[:32], b[32:]}
	return ref, nil
}

/**
 * Creates a new signature.
 *
 * @param R The binary representation of R.
 * @param S The binary representation of S.
 */
//func NewSignature ( []byte R,  []byte S) *Signature {
//    ref := &Signature{
//        if (32 != R.length || 32 != S.length) {
//            panic(IllegalArgumentException{"binary signature representation of R and S must both have 32 bytes length"})
//}
//    return ref
//}
//
//        R,
//        S,
//}

/**
 * Gets the R-part of the signature.
 *
 * @return The R-part of the signature.
 */
func (ref *Signature) GetR() uint32 {

	return binary.BigEndian.Uint32(ref.R)
}

/**
 * Gets the S-part of the signature.
 *
 * @return The S-part of the signature.
 */
func (ref *Signature) GetS() uint32 {

	return binary.BigEndian.Uint32(ref.S)
}

/**
 * Gets a little-endian 64-byte representation of the signature.
 *
 * @return a little-endian 64-byte representation of the signature
 */
func (ref *Signature) getBytes() []byte {

	return append(ref.R, ref.S...)
}

// @Override
func (ref *Signature) String() string {

	return string(ref.getBytes())
}
