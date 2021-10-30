package crypto

import "testing"

func TestCryptography(t *testing.T) {

	var cases = []struct {
		input string
		expected string
	}{
		{
			input:"Texto #3",
			expected:"3# rvzgV",
		},
		{
			input:"abcABC1",
			expected:"1FECedc",
		},
		{
			input:"vxpdylY .ph",
			expected:`ks. \n{frzx`,
		},
		{
			input:"vv.xwfxo.fd",
			expected:"gi.r{hyz-xx",
		},
	}

	for _, scenarios := range cases {
		got := Encrypt(scenarios.input)
		if got != scenarios.expected {
			t.Errorf("Cryptography(v) = %s; Want %s;",got, scenarios.expected)
		}
	}
}

func TestDencrypt(t *testing.T) {
	var cases = []struct {
		input string
		expected string
	}{
		{
			expected:"Texto #3",
			input:"3# rvzgV",
		},
		{
			expected:"abcABC1",
			input:"1FECedc",
		},
		{
			expected:"vxpdylY .ph",
			input:`ks. \n{frzx`,
		},
		{
			expected:"vv.xwfxo.fd",
			input:"gi.r{hyz-xx",
		},
	}

	for _, scenarios := range cases {
		got := Dencrypt(scenarios.input)
		if got != scenarios.expected {
			t.Errorf("Dencrypt(v) = %s; Want %s;",got, scenarios.expected)
		}
	}
}