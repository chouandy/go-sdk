package dotenv

var filePath = ""

var filePrefix = ".env"

var encryptedFileExt = ".enc"

// SetFilePath set file prefix
func SetFilePath(s string) {
	filePath = s
}

// SetFilePrefix set file prefix
func SetFilePrefix(s string) {
	filePrefix = s
}

// SetEncryptedFileExt set encrypted file ext
func SetEncryptedFileExt(s string) {
	encryptedFileExt = s
}
