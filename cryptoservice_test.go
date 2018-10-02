package vbcore

import (
	"encoding/hex"
	"testing"
)

func TestCryptoService_EncryptBase64(t *testing.T) {
	tests := []struct {
		name    string
		aeskey  string
		plain   string
		wantErr bool
	}{
		{"EmptyPlain_ShouldReturnError", "d218729125d0a4302e50b950d9c2df824bef52738ee510d7c385f5f7223c3273", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key, _ := hex.DecodeString(tt.aeskey)
			cs, err := NewCryptoService(key)
			if err != nil {
				t.Errorf("NewCryptoService() error = %v", err)
				return
			}
			_, err = cs.EncryptBase64([]byte(tt.plain))
			if (err != nil) != tt.wantErr {
				t.Errorf("CryptoService.EncryptBase64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestCryptoService_EncryptDecryptRoundtrip(t *testing.T) {
	tests := []struct {
		name   string
		aeskey string
		text   string
	}{
		{"Simple 1", "c4cc0dfc4ae0e45c045727f84ffd373127453bc232230bf1386972ac692436c1", "Text"},
		{"Simple 2", "c4cc0dfc4ae0e45c045727f84ffd373127453bc232230bf1386972ac692436c1", "Some random content"},
		{"Simple 3", "c4cc0dfc4ae0e45c045727f84ffd373127453bc232230bf1386972ac692436c1", "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua."},

		{"Dynamic 1", "c4cc0dfc4ae0e45c045727f84ffd373127453bc232230bf1386972ac692436c1", FastRandomString(32)},
		{"Dynamic 2", "c4cc0dfc4ae0e45c045727f84ffd373127453bc232230bf1386972ac692436c1", FastRandomString(1024)},
		{"Dynamic 3", "c4cc0dfc4ae0e45c045727f84ffd373127453bc232230bf1386972ac692436c1", FastRandomString(1024 * 64)},
		{"Dynamic 4", "c4cc0dfc4ae0e45c045727f84ffd373127453bc232230bf1386972ac692436c1", FastRandomString(1024 * 256)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := 0; i < 1000; i++ {
				key, _ := hex.DecodeString(tt.aeskey)
				cs, _ := NewCryptoService(key)
				cipher, err := cs.Encrypt([]byte(tt.text))
				if err != nil {
					t.Errorf("CryptoService.Encrypt() error = %v", err)
					return
				}

				plain, err := cs.Decrypt(cipher)
				if err != nil {
					t.Errorf("CryptoService.Decrypt() error = %v", err)
					return
				}

				plainStr := string(plain)
				if plainStr != tt.text {
					t.Errorf("CryptoService.Decrypt() plain = %v, want = %v", plainStr, tt.text)
					return
				}
			}
		})
	}
}

func TestCryptoService_EncryptBase64DecryptBase64Roundtrip(t *testing.T) {
	tests := []struct {
		name   string
		aeskey string
		text   string
	}{
		{"Simple 1", "c4cc0dfc4ae0e45c045727f84ffd373127453bc232230bf1386972ac692436c1", "Text"},
		{"Simple 2", "c4cc0dfc4ae0e45c045727f84ffd373127453bc232230bf1386972ac692436c1", "Some random content"},
		{"Simple 3", "c4cc0dfc4ae0e45c045727f84ffd373127453bc232230bf1386972ac692436c1", "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua."},

		{"Dynamic 1", "c4cc0dfc4ae0e45c045727f84ffd373127453bc232230bf1386972ac692436c1", FastRandomString(32)},
		{"Dynamic 2", "c4cc0dfc4ae0e45c045727f84ffd373127453bc232230bf1386972ac692436c1", FastRandomString(1024)},
		{"Dynamic 3", "c4cc0dfc4ae0e45c045727f84ffd373127453bc232230bf1386972ac692436c1", FastRandomString(1024 * 64)},
		{"Dynamic 4", "c4cc0dfc4ae0e45c045727f84ffd373127453bc232230bf1386972ac692436c1", FastRandomString(1024 * 256)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := 0; i < 1000; i++ {
				key, _ := hex.DecodeString(tt.aeskey)
				cs, _ := NewCryptoService(key)
				cipher, err := cs.EncryptBase64([]byte(tt.text))
				if err != nil {
					t.Errorf("CryptoService.EncryptBase64() error = %v", err)
					return
				}

				plain, err := cs.DecryptBase64(cipher)
				if err != nil {
					t.Errorf("CryptoService.DecryptBase64() error = %v", err)
					return
				}

				plainStr := string(plain)
				if plainStr != tt.text {
					t.Errorf("CryptoService.DecryptBase64() plain = %v, want = %v", plainStr, tt.text)
					return
				}
			}
		})
	}
}
