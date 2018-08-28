package crypto

import (
	"bytes"
	"testing"
)

//region basic construction
func TestNewRandomKeyPair_HasPrivateKey(t *testing.T) {

	kp, err := NewRandomKeyPair()
	if err != nil {
		t.Fatal(err)
	}
	if !kp.HasPrivateKey() {
		t.Error("kp.hasPrivateKey() must be true!")
	}
	if kp.privateKey == nil {
		t.Error("kp.privatKey must by not nil!")
	}
	if kp.publicKey == nil {
		t.Error("kp.publicKey must by not nil!")
	}
}

func TestKeyPair_HasPrivateKey(t *testing.T) {
	// Arrange:
	kp1, err := NewRandomKeyPair()
	if err != nil {
		t.Fatal(err)
	}
	kp2, err := NewKeyPair(kp1.privateKey, nil, nil)
	if err != nil {
		t.Error(err)
	} else {
		if !kp2.HasPrivateKey() {
			t.Error("kp2.hasPrivateKey() must be true!")
		}
		if !bytes.Equal(kp2.PrivateKey(), kp1.PrivateKey()) {
			t.Errorf("kp2.privatKey ('%v')\n must by equal \nkp1.privatKey ('%v') !", kp2.privateKey, kp1.privateKey)
		}
		if !bytes.Equal(kp2.PublicKey(), kp1.PublicKey()) {
			t.Errorf("kp2.publicKey ('%v')\n  must by equal \nkp1.publicKey ('%v') !", kp2.publicKey, kp1.publicKey)
		}
	}
}

func TestNewKeyPair(t *testing.T) {

	// Arrange:
	kp1, err := NewRandomKeyPair()
	if err != nil {
		t.Fatal(err)
	}
	kp2, err := NewKeyPair(nil, kp1.publicKey, nil)
	if err != nil {
		t.Error(err)
	} else {
		if kp2.HasPrivateKey() {
			t.Error("kp2.hasPrivateKey() must by equal false!")
		}
		if kp2.privateKey != nil {
			t.Error("kp2.privatKey must by nil!")
		}
		if kp2.publicKey != kp1.publicKey {
			t.Error("kp2.publicKey must by equal kp1.publicKey!")
		}
	}
}

//endregion
func TestNewRandomKeyPair(t *testing.T) {

	kp1, err := NewRandomKeyPair()
	if err != nil {
		t.Fatal(err)
	}
	kp2, err := NewRandomKeyPair()
	if err != nil {
		t.Fatal(err)
	}
	if kp2.privateKey == kp1.privateKey {
		t.Error("kp2.getPrivateKey() and kp1.getPrivateKey() must by not equal !")
	}
	if kp2.publicKey == kp1.publicKey {
		t.Error("kp2.getPublicKey() and kp1.getPublicKey() must by not equal !")
	}
}
