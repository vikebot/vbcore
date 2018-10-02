package vbcore

import (
	"testing"
	"time"
)

func TestCryptoCmp_Timing(t *testing.T) {
	str := FastRandomString(100000000)

	start := str[0:1]
	for start == str[0:1] {
		start = FastRandomString(1)
	}

	end := str[len(str)-1:]
	for end == str[len(str)-1:] {
		end = FastRandomString(1)
	}

	strStart := start + str[1:]
	strEnd := str[:len(str)-1] + end

	strStartTStart := time.Now()
	ok := CryptoCmpStr(str, strStart)
	if ok {
		t.Errorf("CryptoCmpStr() = %v, want %v", ok, false)
	}
	strStartTEnd := time.Now()

	strEndTStart := time.Now()
	ok = CryptoCmpStr(str, strEnd)
	if ok {
		t.Errorf("CryptoCmpStr() = %v, want %v", ok, false)
	}
	strEndTEnd := time.Now()

	startD := strStartTEnd.Sub(strStartTStart)
	endD := strEndTEnd.Sub(strEndTStart)

	if endD > startD*2 {
		t.Errorf("CryptoCmpStr() constant-time-compare issue. Suffix compare took %v, prefix compare took %v", startD, endD)
		return
	}
	t.Logf("CryptoCmpStr() constant-time-compare. Suffix compare took %v, prefix compare took %v", startD, endD)
}

func TestCryptoCmp_Comparison(t *testing.T) {
	dyn1 := FastRandomString(100)
	dyn2 := FastRandomString(10000)
	dyn3 := FastRandomString(1000000)

	tests := []struct {
		name string
		x    string
		y    string
		want bool
	}{
		{"Empty", "", "", true},
		{"Empty x", "", "yIsOK", false},
		{"Empty y", "xIsOK", "", false},

		{"One char", "a", "a", true},
		{"One char big", "B", "B", true},
		{"One char false", "c", "d", false},
		{"One char big false", "E", "F", false},

		{"Dynamic 1 true", dyn1, dyn1, true},
		{"Dynamic 2 true", dyn2, dyn2, true},
		{"Dynamic 3 true", dyn3, dyn3, true},
		{"Dynamic 4 false", FastRandomString(100), FastRandomString(100), false},
		{"Dynamic 5 false", FastRandomString(10000), FastRandomString(10000), false},
		{"Dynamic 6 false", FastRandomString(1000000), FastRandomString(1000000), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CryptoCmpStr(tt.x, tt.y); got != tt.want {
				t.Errorf("CryptoCmpStr() = %v, want %v", got, tt.want)
			}
			if got := CryptoCmpBytes([]byte(tt.x), []byte(tt.y)); got != tt.want {
				t.Errorf("CryptoCmpBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}
