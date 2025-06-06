package main

import (
	"aecsar/password-gen/internal"
	"strings"
	"testing"
)

type test struct {
	name              string
	bytesToNotContain string
	opts              internal.RandomPassOpts
}

func TestGenerateRandomStringBytes(t *testing.T) {
	count := 10
	passwordLength := 10

	tests := []test{
		{
			name:              "exclude numbers",
			opts:              internal.RandomPassOpts{ExcludeNumbers: true},
			bytesToNotContain: internal.NumberBytes,
		},
		{
			name:              "exclude lowercase",
			opts:              internal.RandomPassOpts{ExcludeLowercase: true},
			bytesToNotContain: internal.LowerCaseBytes,
		},
		{
			name:              "exclude uppercase",
			opts:              internal.RandomPassOpts{ExcludeUppercase: true},
			bytesToNotContain: internal.UpperCaseBytes,
		},
		{
			name:              "exclude special chars",
			opts:              internal.RandomPassOpts{ExcludeSpecialChars: true},
			bytesToNotContain: internal.SpecialCharBytes,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			passwords := []string{}

			for range count {
				passwords = append(passwords, internal.GenerateRandomStringBytes(passwordLength, tc.opts))
			}

			for _, pass := range passwords {
				if strings.ContainsAny(pass, tc.bytesToNotContain) {
					t.Error("Expected generated password to not contain excluded characters ", pass)
				}
			}
		})
	}
}
