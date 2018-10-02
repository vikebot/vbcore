package vbcore

import (
	"testing"
)

func TestCryptoGenBytes_Unique(t *testing.T) {
	store := make(map[string]byte)

	n := 200

	for i := 0; i < 1000; i++ {
		buf, err := CryptoGenBytes(n)
		if err != nil {
			t.Fail()
		}

		if len(buf) != 200 {
			t.Fail()
		}

		if _, ok := store[string(buf)]; ok {
			t.Fail()
		} else {
			store[string(buf)] = 0
		}
	}
}

func TestCryptoGen_Unique(t *testing.T) {
	keys := make(map[string]byte)

	for i := 0; i < 1000; i++ {
		k, err := CryptoGen()
		if err != nil {
			t.Fail()
		}

		if _, ok := keys[k]; ok {
			t.Fail()
		} else {
			keys[k] = 0
		}
	}
}

func TestCryptoGenString_Unique(t *testing.T) {
	store := make(map[string]byte)

	n := 200

	for i := 0; i < 1000; i++ {
		str, err := CryptoGenString(n)
		if err != nil {
			t.Fail()
		}

		if len(str) != n {
			t.Fail()
		}

		if _, ok := store[str]; ok {
			t.Fail()
		} else {
			store[str] = 0
		}
	}
}

func TestCryptoGenX_Threshold(t *testing.T) {
	hasCorrectError := func(err error) bool {
		if err != nil {
			if err.Error() != "vbcore.CryptoGenBytes: CryptoGenThreshold mustn't be less than 1 but is 0" {
				return false
			}
		} else {
			return false
		}
		return true
	}

	// Set CryptoGenThreshold to test for custom errors inside CryptoGenX funcs
	CryptoGenThreshold = 0

	key, err := CryptoGen()
	if len(key) > 0 || !hasCorrectError(err) {
		t.Fail()
	}

	buf, err := CryptoGenBytes(100)
	if len(buf) > 0 || !hasCorrectError(err) {
		t.Fail()
	}

	str, err := CryptoGenString(100)
	if len(str) > 0 || !hasCorrectError(err) {
		t.Fail()
	}

	// Reset CryptoGenThreshold
	CryptoGenThreshold = 10
}
