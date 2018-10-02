package vbcore

import (
	"testing"
)

func TestTernaryOperatorA(t *testing.T) {
	type args struct {
		condition  bool
		trueValue  string
		falseValue string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"TrueTest", args{true, "yes", "no"}, "yes"},
		{"FalseTest", args{false, "yes", "no"}, "no"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TernaryOperatorA(tt.args.condition, tt.args.trueValue, tt.args.falseValue); got != tt.want {
				t.Errorf("TernaryOperatorA() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTernaryOperatorI(t *testing.T) {
	type args struct {
		condition  bool
		trueValue  int
		falseValue int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"TrueTest", args{true, 0, 1}, 0},
		{"FalseTest", args{false, 0, 1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TernaryOperatorI(tt.args.condition, tt.args.trueValue, tt.args.falseValue); got != tt.want {
				t.Errorf("TernaryOperatorI() = %v, want %v", got, tt.want)
			}
		})
	}
}
