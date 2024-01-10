package cmd

import (
	"fmt"
	"math/rand"

	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a random password",
	Long: `Generate a random password of custom length and choose to include numbers and symbols.
  Example: pwdg generate -l 8 -d -s`,
	Run: generatePassword,
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().IntP("length", "l", 8, "Length for your password")
	generateCmd.Flags().BoolP("numbers", "n", false, "Include numbers in your password")
	generateCmd.Flags().BoolP("symbols", "s", false, "Include symobls in your password")
}

func generatePassword(cmd *cobra.Command, args []string) {
	length, _ := cmd.Flags().GetInt("length")
	includeNumbers, _ := cmd.Flags().GetBool("numbers")
	includeSymbols, _ := cmd.Flags().GetBool("symbols")

	charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	if includeNumbers {
		charset += "0123456789"
	}
	if includeSymbols {
		charset += "!@#$%^&*()_-+={}[]?<>;:"
	}

	password := make([]byte, length)

	for i := range password {
		password[i] = charset[rand.Intn(len(charset))]
	}

	fmt.Println("Generating password...")
	fmt.Println(string(password))
}
