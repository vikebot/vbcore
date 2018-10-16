package vbcore

import "testing"

func TestFastRandomString(t *testing.T) {
	tests := []struct {
		name string
		arg  int
	}{
		{"Zero", 0},
		{"Small", 100},
		{"Large", 10000},
		{"Very large", 1000000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FastRandomString(tt.arg); len(got) != tt.arg {
				t.Errorf("FastRandomString() = %v, len = %v want %v", got, len(got), tt.arg)
			}
		})
	}
}

func TestFastRandomString_Unique(t *testing.T) {
	tests := []struct {
		name string
		arg  int
	}{
		{"Small", 100},
		{"Large", 10000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := make(map[string]bool)

			for i := 0; i < 1000; i++ {
				got := FastRandomString(tt.arg)
				if ok := u[got]; ok {
					t.Errorf("FastRandomString() = %v, already exists in map", got)
				}
				u[got] = true
			}
		})
	}
}
