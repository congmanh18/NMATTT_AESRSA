package routes

import (
	"net/http"

	"github.com/congmanh18/NMATTT_AESRSA/database"
	"github.com/congmanh18/NMATTT_AESRSA/handler"
)

// SetupAESRoutes sets up the AES encryption and decryption routes.
func RSARoutes(repo *database.Repository) {
	http.HandleFunc("/RSA/encryption", handler.EncryptionRSAHandler(repo))
	http.HandleFunc("/RSA/decryption", handler.DecryptionRSAHandler(repo))
}
