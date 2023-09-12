/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// encryptCmd represents the encrypt command
var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		key, err := hex.DecodeString(viper.GetString("crypto.aes_key"))
		cobra.CheckErr(err)
		if len(key) != 32 {
			log.Fatal("The AES_KEY environment variable must be exactly 32 bytes long.")
		}

		toEncrypt := strings.Join(args, " ")
		encrypted, err := encrypt(toEncrypt, key)
		if err != nil {
			log.Fatal("Failed to encrypt: ", err)
		}
		fmt.Printf("Encrypted: %s\n", encrypted)
	},
}

func init() {
	CryptoCmd.AddCommand(encryptCmd)
}

func encrypt(stringToEncrypt string, keyString []byte) (string, error) {
	key := keyString
	plaintext := []byte(stringToEncrypt)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

/*


## Encryption

- Initialization: The aes.NewCipher function initializes a new AES cipher with our 32-byte key. Then, cipher.NewGCM creates a new GCM cipher based on this AES cipher.

```
block, err := aes.NewCipher(key)
aesGCM, err := cipher.NewGCM(block)
```

- Nonce Creation: GCM requires a nonce which is generated cryptographically randomly. The nonce size is determined by aesGCM.NonceSize().
```
nonce := make([]byte, aesGCM.NonceSize())
if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
    return "", err
}
```

- Encryption: The aesGCM.Seal() function is used for encrypting the plaintext. It also includes the nonce and any associated data (which is nil in this case).
```
ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)
```

**Note** The Seal() function returns a slice that consists of the nonce followed by the encrypted data.

*/
