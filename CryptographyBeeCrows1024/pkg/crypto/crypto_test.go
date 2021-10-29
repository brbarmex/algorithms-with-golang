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
			input:"vxpdyly .ph",
			expected:`ks. \n{frzx`,
		},
		{
			input:"vv.xwfxo.fd",
			expected:"gi.r{hyz-xx",
		},
	}

	for _, scenarios := range cases {
		got := Cryptography(scenarios.input)
		if got != scenarios.expected {
			t.Errorf("Cryptography(v) = %s; Want %s;",got, scenarios.expected)
		}
	}

}