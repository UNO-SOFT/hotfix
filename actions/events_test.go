package actions

import (
	"bytes"
	"encoding/base64"
	"testing"
)

const (
	sigMessage  = "proba"
	naclPubkey  = NaCLPublicPrefix + `rmWa1cGJb38fTh/JFAVMP1H5G2f2jIk1qKG0kxxryEU=`
	naclPrivkey = NaCLPrivatePrefix + `MxsTgOGK8JWfezUC0MTh/ZK9LpI5zSbkpBOfMsdgaVKuZZrVwYlvfx9OH8kUBUw/UfkbZ/aMiTWoobSTHGvIRQ==`
)

func TestVerifySignatureNaCL(t *testing.T) {
	pubKey, privKey, err := GenerateKey(nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("pub: %s", base64.StdEncoding.EncodeToString(pubKey))
	t.Logf("sec: %s", base64.StdEncoding.EncodeToString(privKey))

	recip, err := ParseKey(nil, []byte(naclPubkey))
	if err != nil {
		t.Fatal(err)
	}
	ident, err := ParseKey(nil, []byte(naclPrivkey))
	_, _ = recip, ident
	signedMessage := Sign(nil, []byte(sigMessage), ident)
	t.Logf("%x", signedMessage)
	msg, ok := Open(nil, signedMessage, recip)
	if !ok {
		t.Fatal("not ok")
	}
	if !bytes.Equal(msg, []byte(sigMessage)) {
		t.Errorf("signature mismatch: got %x wanted %x", msg, sigMessage)
	}
}
