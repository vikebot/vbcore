package vbcore

import (
	"testing"
)

func TestMinInt(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Positive zero (ASC)", args{0, 1}, 0},
		{"Positive zero (DESC)", args{1, 0}, 0},
		{"Positive one (ASC)", args{1, 2}, 1},
		{"Positive one (DESC)", args{2, 1}, 1},

		{"Mixed zero (ASC)", args{-1, 0}, -1},
		{"Mixed zero (DESC)", args{0, -1}, -1},
		{"Mixed one (ASC)", args{-1, 1}, -1},
		{"Mixed one (DESC)", args{1, -1}, -1},

		{"Negative (ASC)", args{-2, -1}, -2},
		{"Negative (DESC)", args{-1, -2}, -2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinInt(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("MinInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxInt(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Positive zero (ASC)", args{0, 1}, 1},
		{"Positive zero (DESC)", args{1, 0}, 1},
		{"Positive one (ASC)", args{1, 2}, 2},
		{"Positive one (DESC)", args{2, 1}, 2},

		{"Mixed zero (ASC)", args{-1, 0}, 0},
		{"Mixed zero (DESC)", args{0, -1}, 0},
		{"Mixed one (ASC)", args{-1, 1}, 1},
		{"Mixed one (DESC)", args{1, -1}, 1},

		{"Negative (ASC)", args{-2, -1}, -1},
		{"Negative (DESC)", args{-1, -2}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxInt(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("MaxInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
