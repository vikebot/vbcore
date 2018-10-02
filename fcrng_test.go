package vbcore

import (
	"testing"
)

func TestFCRNG_NextU64(t *testing.T) {
	tests := []struct {
		name string
		arg  uint64
	}{
		{"Small", 5},
		{"Midium", 500},
		{"Large", 10000},

		{"Male", 20},
		{"Female", 15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := NewFCRNG(tt.arg)

			m := make(map[uint64]bool)
			for i := 0; i < int(tt.arg); i++ {
				got := f.NextU64()
				if ok := m[got]; ok {
					t.Errorf("FCRNG.NextU64() = %v, already in map", got)
				}
				m[got] = true
			}

			for i := 0; i < int(tt.arg); i++ {
				if ok := m[uint64(i)]; !ok {
					t.Errorf("FCRNG.NextU64(), resulting sequence set doesn't contain %v", i)
				}
			}
		})
	}
}
