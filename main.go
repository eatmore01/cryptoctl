package main

import (
	"fmt"
	"os"

	"github.com/eatmore01/crypto"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "cryptoctl",
		Short: "cryptoctl is a lightweight CLI tool for simple crypto operations with text and files",
		Long:  "CLI encoded and decoded a base64 encoded string and files",
	}

	rootCmd.AddCommand(encryptStringCmd)
	rootCmd.AddCommand(decryptStringCmd)
	rootCmd.AddCommand(encryptFileCmd)
	rootCmd.AddCommand(decryptFileCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func prepareCrypto() (*crypto.Crypto, string,error) {
	passwd := os.Getenv("CRYPTOCTL_PASSWORD")
	if passwd == "" {
		return nil, "", fmt.Errorf("CRYPTOCTL_PASSWORD is not set")
	}

	return crypto.New(passwd), passwd, nil
}	

var encryptStringCmd = &cobra.Command{
	Use:   "encrypt-str",
	Short: "Encrypt a string with base64 encoding",
	Long:  "Encrypt a string to base64 encoded string",
	Run: func(cmd *cobra.Command, args []string) {
		c, passwd, err := prepareCrypto()

		if err != nil {
			fmt.Println(err)
			return
		}

		if len(args) != 1 {
			fmt.Println("Please provide a string to encrypt")
			return
		}

		encodedText, err := c.EncryptText(args[0], passwd); if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("base64 encoded text: ", encodedText)
	},
}

var decryptStringCmd = &cobra.Command{
	Use:   "decrypt-str",
	Short: "Decrypt a string with base64 encoding",
	Long:  "Decrypt a base64 encoded string to original text",
	Run: func(cmd *cobra.Command, args []string) {
		c, passwd, err := prepareCrypto()

		if err != nil {
			fmt.Println(err)
			return
		}

		if len(args) != 1 {
			fmt.Println("Please provide a string to decrypt")
			return
		}

		decodedText, err := c.DecryptText(args[0], passwd); if err != nil {
			fmt.Println(err)
			return
		}	

		fmt.Println("decoded text: ", decodedText)
	},
}


var encryptFileCmd = &cobra.Command{
	Use:   "encrypt-file",
	Short: "Encrypt a file and add *.enc extension",
	Long:  "Encrypt a file and add *.enc extension, for example hello.txt will be encrypted to hello.txt.enc",
	Run: func(cmd *cobra.Command, args []string) {
		c, passwd, err := prepareCrypto()

		if err != nil {
			fmt.Println(err)
			return
		}

		if len(args) != 1 {
			fmt.Println("Please provide a file name to encrypt")
			return
		}

		if err := c.EncryptFile(args[0], passwd); err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("file encrypted successfully and have *.enc extension")
	},
}



var decryptFileCmd = &cobra.Command{
	Use:   "decrypt-file",
	Short: "Decrypt a file and remove *.enc extension",
	Long:  "Decrypt a file and remove *.enc extension, for example hello.txt.enc will be decrypted to hello.txt",
	Run: func(cmd *cobra.Command, args []string) {
		c, passwd, err := prepareCrypto()

		if err != nil {
			fmt.Println(err)
			return
		}

		if len(args) != 1 {
			fmt.Println("Please provide a file name to decrypt")
			return
		}

		if err := c.DecryptFile(args[0], passwd); err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("file decrypted successfully and removed *.enc extension")
	},
}
