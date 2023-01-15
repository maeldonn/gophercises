package main

import "testing"

func TestCamelcase(t *testing.T) {
	testcases := []struct {
		in   string
		want int32
	}{
		{"saveChangesInTheEditor", 5},
		{"oneTwoThree", 3},
	}
	for _, tc := range testcases {
		c := camelcase(tc.in)
		if c != tc.want {
			t.Errorf("camelcase: %q, want %q", c, tc.want)
		}
	}
}

func TestCaesarCipher(t *testing.T) {
	testcases := []struct {
		in1  string
		in2  int32
		want string
	}{
		{"There's-a-starman-waiting-in-the-sky", 3, "Wkhuh'v-d-vwdupdq-zdlwlqj-lq-wkh-vnb"},
		{"middle-Outz", 2, "okffng-Qwvb"},
	}
	for _, tc := range testcases {
		c := caesarCipher(tc.in1, tc.in2)
		if c != tc.want {
			t.Errorf("caesarCipher: %q, want %q", c, tc.want)
		}
	}
}
