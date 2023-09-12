/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// decryptCmd represents the decrypt command
var decryptCmd = &cobra.Command{
	Use:   "decrypt",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		key, err := hex.DecodeString(viper.GetString("crypto.aes_key"))
		cobra.CheckErr(err)
		if len(key) != 32 {
			log.Fatal("The AES_KEY environment variable must be exactly 32 bytes long.")
		}

		toDecrypt := strings.Join(args, " ")
		decrypted, err := decrypt(toDecrypt, key)
		if err != nil {
			log.Fatal("Failed to decrypt: ", err)
		}
		fmt.Printf("Decrypted: %s\n", decrypted)
	},
}

func init() {
	CryptoCmd.AddCommand(decryptCmd)
}

func decrypt(encryptedString string, keyString []byte) (string, error) {
	key := keyString
	enc, err := base64.StdEncoding.DecodeString(encryptedString)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := aesGCM.NonceSize()
	if len(enc) < nonceSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

/*



## Decryption
- Initialization: Similar to the encryption step, initialize the AES and GCM ciphers.
```
block, err := aes.NewCipher(key)
aesGCM, err := cipher.NewGCM(block)
```

- Nonce Extraction: The first aesGCM.NonceSize() bytes of the encrypted data are the nonce. This needs to be extracted to decrypt the remaining data.
```
nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]
```

- Decryption: The aesGCM.Open() function decrypts the ciphertext back into plaintext. This function takes the nonce and the ciphertext and returns the decrypted data if successful.
```
plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
```

The code uses base64 encoding to encode the binary data into a string format that can be easily displayed or transmitted.



*/
