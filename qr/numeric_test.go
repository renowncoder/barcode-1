package qr

import (
	"bytes"
	"testing"
)

func Test_NumericEncoding(t *testing.T) {
	x, vi, err := Numeric.encode("01234567", H)
	if x == nil || vi == nil || vi.Version != 1 || bytes.Compare(x.GetBytes(), []byte{16, 32, 12, 86, 97, 128, 236, 17, 236}) != 0 {
		t.Error("\"01234567\" failed to encode")
	}
	x, vi, err = Numeric.encode("0123456789012345", H)
	if x == nil || vi == nil || vi.Version != 1 || bytes.Compare(x.GetBytes(), []byte{16, 64, 12, 86, 106, 110, 20, 234, 80}) != 0 {
		t.Error("\"0123456789012345\" failed to encode")
	}
	x, vi, err = Numeric.encode("foo", H)
	if err == nil {
		t.Error("Numeric encoding should not be able to encode \"foo\"")
	}
}
