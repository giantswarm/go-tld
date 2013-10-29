package tld

import (
	"testing"
)

func Test_Update(t *testing.T) {
	if err := Update(IANA); err != nil {
		t.Error("Failed")
	} else {
		t.Log("Passed")
	}
}

var tldTests = []struct {
	TLD   []byte
	Valid bool
}{
	{[]byte("com"), true},
	{[]byte("123"), false},
	{[]byte("potato"), false},
	{[]byte("xxx"), true},
}

func Test_Valid(t *testing.T) {
	for i, tld := range tldTests {
		if v := Valid(tld.TLD); v != tld.Valid {
			t.Errorf("%d. Valid(\"%s\") returned %v, want %v",
				i, tld.TLD, v, tld.Valid)
		}
	}
}
