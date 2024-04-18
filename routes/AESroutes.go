package routes

import (
	"net/http"

	"github.com/congmanh18/NMATTT_AESRSA/database"
	"github.com/congmanh18/NMATTT_AESRSA/handler"
)

// SetupAESRoutes sets up the AES encryption and decryption routes.
func AESRoutes(repo *database.Repository) {
	http.HandleFunc("/AES/encryption", handler.EncryptionAESHandler(repo))
	http.HandleFunc("/AES/decryption", handler.DecryptionAESHandler(repo))
}
