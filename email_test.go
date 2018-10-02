package vbcore

import (
	"testing"
)

func TestEmailItoA(t *testing.T) {
	if EmailItoA(EmailLinked) != EmailLinkedString {
		t.Fail()
	}
	if EmailItoA(EmailVerified) != EmailVerifiedString {
		t.Fail()
	}
	if EmailItoA(EmailPrimary) != EmailPrimaryString {
		t.Fail()
	}
}

func TestEmailItoA_Lesser(t *testing.T) {
	if EmailItoA(EmailLinked-1) != EmailLinkedString {
		t.Fail()
	}
}

func TestEmailItoA_Bigger(t *testing.T) {
	if EmailItoA(EmailPrimary+1) != EmailPrimaryString {
		t.Fail()
	}
}

func TestEmailAtoI(t *testing.T) {
	if EmailAtoI(EmailLinkedString) != EmailLinked {
		t.Fail()
	}
	if EmailAtoI(EmailVerifiedString) != EmailVerified {
		t.Fail()
	}
	if EmailAtoI(EmailPrimaryString) != EmailPrimary {
		t.Fail()
	}
}

func TestEmailAtoI_LinkedFallback(t *testing.T) {
	if EmailAtoI("TRY_FALLBACK") != EmailLinked {
		t.Fail()
	}
}

func TestEmail_IsLinked(t *testing.T) {
	tests := []struct {
		name        string
		emailStatus int
		want        bool
	}{
		{"Linked", EmailLinked, true},
		{"Verified", EmailVerified, false},
		{"Primary", EmailPrimary, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := Email{
				Email:  "",
				Status: tt.emailStatus,
				Public: false,
			}
			if got := e.IsLinked(); got != tt.want {
				t.Errorf("Email.IsLinked() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEmail_IsVerified(t *testing.T) {
	tests := []struct {
		name        string
		emailStatus int
		want        bool
	}{
		{"Linked", EmailLinked, false},
		{"Verified", EmailVerified, true},
		{"Primary", EmailPrimary, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := Email{
				Email:  "",
				Status: tt.emailStatus,
				Public: false,
			}
			if got := e.IsVerified(); got != tt.want {
				t.Errorf("Email.IsVerified() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEmail_IsPrimary(t *testing.T) {
	tests := []struct {
		name        string
		emailStatus int
		want        bool
	}{
		{"Linked", EmailLinked, false},
		{"Verified", EmailVerified, false},
		{"Primary", EmailPrimary, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := Email{
				Email:  "",
				Status: tt.emailStatus,
				Public: false,
			}
			if got := e.IsPrimary(); got != tt.want {
				t.Errorf("Email.IsPrimary() = %v, want %v", got, tt.want)
			}
		})
	}
}
