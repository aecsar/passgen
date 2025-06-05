package main

import (
	"fmt"
	"math/rand"
)

type RandomPassOpts struct {
	excludeNumbers, excludeLowercase, excludeUppercase, excludeSpecialChars bool
}

func GenerateRandomStringBytes(n int, options RandomPassOpts) string {

	numberBytes := "0123456789"
	lowerCaseBytes := "abcdefghijklmnopqrstuvwxyz"
	upperCaseBytes := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharBytes := "!@#$%^&*()_+-=[]{}\\|;':\",./<>?~`"

	finalBytes := []byte{}

	if !options.excludeNumbers {
		finalBytes = append(finalBytes, ([]byte)(numberBytes)...)
	}

	if !options.excludeLowercase {
		finalBytes = append(finalBytes, ([]byte)(lowerCaseBytes)...)
	}

	if !options.excludeUppercase {
		finalBytes = append(finalBytes, ([]byte)(upperCaseBytes)...)
	}

	if !options.excludeSpecialChars {
		finalBytes = append(finalBytes, ([]byte)(specialCharBytes)...)
	}

	b := make([]byte, n)
	for i := range b {
		b[i] = finalBytes[rand.Intn(len(finalBytes))]
	}

	return string(b)
}

func main() {
	// passgen [lenght] --no-nums?[nn] --no-symbols?[ns] --no-upper[nu] --count?c --copy?true-default,
	pswd := GenerateRandomStringBytes(10, RandomPassOpts{excludeSpecialChars: true})

	fmt.Println(pswd)
}
