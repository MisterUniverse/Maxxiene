/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package crypto

import (
	"github.com/spf13/cobra"
)

// cryptoCmd represents the crypto command
var CryptoCmd = &cobra.Command{
	Use:   "crypto",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {}

/*

**Components**
1. AES (Advanced Encryption Standard): It is a symmetric encryption algorithm that was established by the
U.S. National Institute of Standards and Technology (NIST) in 2001. In symmetric-key encryption, the same key is
used for both encryption and decryption.

2. GCM (Galois/Counter Mode): This is a mode of operation for symmetric-key cryptographic block ciphers.
It provides both encryption and built-in integrity checking. So, AES-GCM is AES encryption in Galois/Counter Mode.

3. Nonce: A nonce ("number used once") is an arbitrary number that should only be used once. It's often used for
initialization in encryption algorithms.

*/
