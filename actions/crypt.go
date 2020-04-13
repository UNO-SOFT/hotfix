package actions

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"errors"

	"golang.org/x/crypto/nacl/sign"
)

func GenerateKey(pubOut, privOut []byte) (publicKey, privateKey []byte, err error) {
	pubKey, privKey, err := sign.GenerateKey(rand.Reader)
	if err != nil {
		return pubOut, privOut, err
	}
	pubOut = append(pubOut, (*pubKey)[:32]...)
	privOut = append(privOut, (*privKey)[:64]...)
	return pubOut, privOut, err
}

const NaCLPublicPrefix = "nacl"
const NaCLPrivatePrefix = "NACL-SECRET-KEY-"

func ParseKey(out, in []byte) ([]byte, error) {
	if bytes.HasPrefix(in, []byte(NaCLPublicPrefix)) {
		in = in[4:]
	} else if bytes.HasPrefix(in, []byte(NaCLPrivatePrefix)) {
		in = in[len(NaCLPrivatePrefix):]
	} else {
		return out, errors.New("not a NaCL key")
	}
	out = append(make([]byte, base64.StdEncoding.DecodedLen(len(in))))
	n, err := base64.StdEncoding.Decode(out, in)
	if err != nil {
		return out, err
	}
	return out[:n], nil
}
func Open(out, signedMessage, publicKey []byte) ([]byte, bool) {
	var pubKey [32]byte
	copy(pubKey[:], publicKey[:32])
	return sign.Open(out, signedMessage, &pubKey)

}
func Sign(out, message, privateKey []byte) []byte {
	var privKey [64]byte
	n := 64
	if len(privateKey) < 64 {
		n = len(privateKey)
		if n < 32 {
			n = 32
		}
	}
	copy(privKey[:], privateKey[:n])
	return sign.Sign(out, message, &privKey)
}
