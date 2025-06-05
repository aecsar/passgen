package internal

import "math/rand"

type RandomPassOpts struct {
	ExcludeNumbers, ExcludeLowercase, ExcludeUppercase, ExcludeSpecialChars bool
}

func GenerateRandomStringBytes(n int, options RandomPassOpts) string {

	numberBytes := "0123456789"
	lowerCaseBytes := "abcdefghijklmnopqrstuvwxyz"
	upperCaseBytes := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharBytes := "!@#$%^&*()_+-=[]{}\\|;':\",./<>?~`"

	finalBytes := []byte{}

	if !options.ExcludeNumbers {
		finalBytes = append(finalBytes, ([]byte)(numberBytes)...)
	}

	if !options.ExcludeLowercase {
		finalBytes = append(finalBytes, ([]byte)(lowerCaseBytes)...)
	}

	if !options.ExcludeUppercase {
		finalBytes = append(finalBytes, ([]byte)(upperCaseBytes)...)
	}

	if !options.ExcludeSpecialChars {
		finalBytes = append(finalBytes, ([]byte)(specialCharBytes)...)
	}

	b := make([]byte, n)
	for i := range b {
		b[i] = finalBytes[rand.Intn(len(finalBytes))]
	}

	return string(b)
}
