package vbcore

import (
	"encoding/hex"
	"reflect"
	"testing"
)

func Test_cryptoPwd_SameOutput(t *testing.T) {
	tests := []struct {
		name     string
		password string
	}{
		{"Dynamic 1", FastRandomString(6)},
		{"Dynamic 2", FastRandomString(8)},
		{"Dynamic 3", FastRandomString(10)},
		{"Dynamic 4", FastRandomString(12)},
		{"Dynamic 5", FastRandomString(16)},
		{"Dynamic 6", FastRandomString(32)},
		{"Dynamic 7", FastRandomString(128)},
		{"Dynamic 8", FastRandomString(512)},
		{"Dynamic 9", FastRandomString(2048)},
		{"Dynamic 10", FastRandomString(1000000)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			salt, err := CryptoGenBytes(32)
			if err != nil {
				t.Errorf("CryptoGenBytes() error = %v", err)
				return
			}

			hash := cryptoPwd(tt.password, salt)
			for i := 0; i < 10; i++ {
				gotHash := cryptoPwd(tt.password, salt)
				if !reflect.DeepEqual(hash, gotHash) {
					t.Errorf("cryptoPwd() = %s on %d time but %s initially", hex.EncodeToString(gotHash), i+1, hex.EncodeToString(hash))
				}
			}
		})
	}
}

func TestCryptoPwd_CryptoPwdVerify_Roundtrip(t *testing.T) {
	tests := []struct {
		name     string
		password string
	}{
		{"Dynamic 1", FastRandomString(1024)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hash, salt, err := CryptoPwd(tt.password)
			if err != nil {
				t.Errorf("CryptoPwd() error = %v", err)
				return
			}

			t.Logf("hash = %s, salt = %s", hash, salt)

			ok, err := CryptoPwdVerify(hash, salt, tt.password)
			if err != nil {
				t.Errorf("CryptoPwdVerify() error = %v", err)
				return
			}
			if !ok {
				t.Errorf("CryptoPwdVerify() = %v", false)
			}
		})
	}
}
