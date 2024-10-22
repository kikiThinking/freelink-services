package encryption

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
)

func DecryptPassword(password, salt string) string {
	hash := sha256.Sum256([]byte(password + salt))
	return hex.EncodeToString(hash[:])
}

func EncryptionPassword(password string) (string, string) {
	saltBytes := make([]byte, 16)
	for i := 0; i < 16; i++ {
		saltBytes[i] = byte(rand.Intn(256))
	}

	return DecryptPassword(password, hex.EncodeToString(saltBytes)), hex.EncodeToString(saltBytes)
}
