package internal

import "math/rand"

type RandomPassOpts struct {
	ExcludeNumbers, ExcludeLowercase, ExcludeUppercase, ExcludeSpecialChars bool
}

const (
	NumberBytes      = "0123456789"
	LowerCaseBytes   = "abcdefghijklmnopqrstuvwxyz"
	UpperCaseBytes   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	SpecialCharBytes = "!@#$%^&*()_+-=[]{}\\|;':\",./<>?~`"
)

func GenerateRandomStringBytes(n int, options RandomPassOpts) string {

	finalBytes := []byte{}

	if !options.ExcludeNumbers {
		finalBytes = append(finalBytes, ([]byte)(NumberBytes)...)
	}

	if !options.ExcludeLowercase {
		finalBytes = append(finalBytes, ([]byte)(LowerCaseBytes)...)
	}

	if !options.ExcludeUppercase {
		finalBytes = append(finalBytes, ([]byte)(UpperCaseBytes)...)
	}

	if !options.ExcludeSpecialChars {
		finalBytes = append(finalBytes, ([]byte)(SpecialCharBytes)...)
	}

	b := make([]byte, n)
	for i := range b {
		b[i] = finalBytes[rand.Intn(len(finalBytes))]
	}

	return string(b)
}
