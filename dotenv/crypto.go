package dotenv

import cryptoex "github.com/chouandy/go-sdk/crypto"

// EncryptFile encrypt file
func EncryptFile(stage string, password []byte) error {
	src := filePrefix + "." + stage
	dst := src + encryptedFileExt

	return cryptoex.FileEncrypter(src, dst, password)
}

// DecryptFile encrypt file
func DecryptFile(stage string, password []byte) error {
	dst := filePrefix + "." + stage
	src := dst + encryptedFileExt

	return cryptoex.FileDecrypter(src, dst, password)
}
