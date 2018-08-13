package crypto

import "testing"

//region basic construction
// @Test
func TestNewRandomKeyPair_HasPrivateKey(t *testing.T) {

	kp, err := NewRandomKeyPair()
	if err != nil {
		t.Fatal(err)
	}
	// Assert:
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

// @Test

func TestKeyPair_HasPrivateKey(t *testing.T) {
	// Arrange:
	kp1, err := NewRandomKeyPair()
	if err != nil {
		t.Fatal(err)
	}
	// Act:
	kp2, err := NewKeyPair(kp1.privateKey, nil, nil)
	if err != nil {
		t.Error(err)
	} else {
		// Assert:
		if !kp2.HasPrivateKey() {
			t.Error("kp2.hasPrivateKey() must be true!")
		}
		if kp2.privateKey != kp1.privateKey {
			t.Error("kp2.privatKey must by equal kp1.privatKey!")
		}
		if kp2.publicKey != kp1.publicKey {
			t.Error("kp2.publicKey must by equal kp1.publicKey!")
		}
	}
}

// @Test
func TestNewKeyPair(t *testing.T) {

	// Arrange:
	kp1, err := NewRandomKeyPair()
	if err != nil {
		t.Fatal(err)
	}
	// Act:
	kp2, err := NewKeyPair(nil, kp1.publicKey, nil)
	if err != nil {
		t.Error(err)
	} else {
		// Assert:
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
// @Test
func TestNewRandomKeyPair(t *testing.T) {

	// Act:
	kp1, err := NewRandomKeyPair()
	if err != nil {
		t.Fatal(err)
	}
	kp2, err := NewRandomKeyPair()
	if err != nil {
		t.Fatal(err)
	}
	// Assert:
	if kp2.privateKey == kp1.privateKey {
		t.Error("kp2.getPrivateKey() and kp1.getPrivateKey() must by not equal !")
	}
	if kp2.publicKey == kp1.publicKey {
		t.Error("kp2.getPublicKey() and kp1.getPublicKey() must by not equal !")
	}
}
