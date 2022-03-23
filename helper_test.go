package main

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	want := "kI3GCqW5JLMJa4iX1lo7X4D6XbYqlLgxIs30+P6tENUV2POR"
	p := &Password{
		Salt:     [4]byte{0x90, 0x8D, 0xC6, 0x0A},
		Password: "test12",
	}

	if p.HashBase64() != want {
		t.Fatal("failed: Hash does not follow algorithm")
	}
}

func TestHashPasswordHelper(t *testing.T) {
	pwd := "p4Ssw0Rd"
	hashPassword(pwd)
	hashPasswordString(pwd)
	hashPasswordBase64(pwd)
}
