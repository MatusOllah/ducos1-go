package ducos1_test

import (
	"testing"

	. "github.com/MatusOllah/ducos1-go"
)

func TestDUCOS1(t *testing.T) {
	tests := []struct {
		hash         string
		expectedHash string
		difficulty   int
	}{
		{"b9329274bca0970a09977ed64601843ec660c92d", "e504239a8adc31c3fcb3498fa2d76bf139fe773e", 20000},
		{"dc5ca1d73c4208b25d32e2da6b00f7c273ea551e", "31cba56b3245abd5ffceda2f467b28c1ceb2b885", 100000},
		{"f49d02ea0faf1851e3effb299419487fdff9cfa3", "c078292a82c2b6cb8d549cd144963eb9ef8a287f", 200000},
	}

	for _, test := range tests {
		hash := DUCOS1(test.hash, test.expectedHash, test.difficulty)
		if hash != test.expectedHash {
			t.Errorf("expected hash %s, got %s", test.expectedHash, hash)
		}
	}
}
