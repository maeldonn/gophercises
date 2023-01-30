package main

import "testing"

func TestNormalize(t *testing.T) {
	testcases := []struct {
		in, want string
	}{
		{"1234567890", "1234567890"},
		{"123 456 7891", "1234567891"},
		{"(123) 456 7892", "1234567892"},
		{"(123) 456-7893", "1234567893"},
		{"123-456-7894", "1234567894"},
		{"(123)456-7892", "1234567892"},
	}
	for _, tc := range testcases {
		t.Run(tc.in, func(t *testing.T) {
			actual := normalize(tc.in)
			if actual != tc.want {
				t.Errorf("normalize(%s) = %s, want %s", tc.in, tc.want, actual)
			}
		})
	}
}
