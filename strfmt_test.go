package vbcore

import "testing"

func TestItoNumberUnity(t *testing.T) {
	tests := []struct {
		name string
		arg  int
		want string
	}{
		{"Below zero", -10, "-10"},
		{"Zero", 0, "0"},
		{"Below thousand", 999, "999"},
		{"Exactly thousand", 1000, "1K"},
		{"A few thousand", 7400, "7K"},
		{"A few ten thousand", 49999, "49K"},
		{"A few hundred thousand", 700912, "700K"},
		{"Million", 10700912, "10M"},
		{"Billion", 84200700912, "84B"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ItoNumberUnity(tt.arg); got != tt.want {
				t.Errorf("ItoNumberUnity() = %v, want %v", got, tt.want)
			}
		})
	}
}
