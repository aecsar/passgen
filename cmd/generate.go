package cmd

import (
	"aecsar/password-gen/internal"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var (
	ExcludeNumbers      bool
	ExcludeSpecialChars bool
	ExcludeUpper        bool
	ExcludeLower        bool
	Count               int
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(&ExcludeNumbers, "no-nums", "n", false, "exclude numbers")
	rootCmd.PersistentFlags().BoolVarP(&ExcludeSpecialChars, "no-symbols", "s", false, "exclude special characters")
	rootCmd.PersistentFlags().BoolVarP(&ExcludeUpper, "no-upper", "u", false, "exclude uppercase letters")
	rootCmd.PersistentFlags().BoolVarP(&ExcludeLower, "no-lower", "l", false, "exclude lowercase letters")
	rootCmd.PersistentFlags().IntVarP(&Count, "count", "c", 1, "numbers of random passwords to generate")
}

var rootCmd = &cobra.Command{
	Use:   "passgen",
	Short: "Generate secure random passwords",
	Long: `passgen is a CLI tool for generating cryptographically secure random passwords.
By default, passwords include uppercase letters, lowercase letters, numbers, and special characters.
Use flags to customize which character types to exclude from generation. It also copies the generated
password to the clipboard

Examples:
  passgen 16                   Generate 16-character password with all character types
  passgen 12 --no-symbols      Generate 12-character password without special characters
  passgen 20 -ns               Generate 20-character password without numbers or symbols
  passgen 8 --count 5          Generate 5 passwords of 8 characters each`,
	Args: cobra.RangeArgs(0, 1),

	Run: func(cmd *cobra.Command, args []string) {

		length := 10

		if len(args) > 0 {
			var err error

			length, err = strconv.Atoi(args[0])
			if err != nil || length < 1 {
				fmt.Println("Length must be a number higher or equal to 1")
				return
			}
		}

		if ExcludeNumbers && ExcludeLower && ExcludeUpper && ExcludeSpecialChars {
			fmt.Println("You cannot exclude all type of characters")
			return
		}

		if Count < 1 {
			fmt.Println("The count should be equal higher or equal to 1")
			return
		}

		passwords := []string{}

		for range Count {
			passwords = append(passwords,
				internal.GenerateRandomStringBytes(
					length,
					internal.RandomPassOpts{
						ExcludeLowercase:    ExcludeLower,
						ExcludeUppercase:    ExcludeUpper,
						ExcludeNumbers:      ExcludeNumbers,
						ExcludeSpecialChars: ExcludeSpecialChars,
					},
				),
			)
		}

		for _, pswd := range passwords {
			fmt.Println(pswd)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
