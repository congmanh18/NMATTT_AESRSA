package handler

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/congmanh18/NMATTT_AESRSA/database"
	"github.com/congmanh18/NMATTT_AESRSA/model"
)

// EncryptAES mã hóa thông điệp bằng AES
func EncryptAES(plainText []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plainText))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plainText)

	return ciphertext, nil
}

// DecryptAES giải mã thông điệp đã được mã hóa bằng AES
func DecryptAES(ciphertext []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < aes.BlockSize {
		return nil, fmt.Errorf("ciphertext is too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	return ciphertext, nil
}

// EncryptionAESHandler handles requests for AES encryption.
func EncryptionAESHandler(repo *database.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse request body
		var data model.Data
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Encrypt the message
		encryptedMessage, err := EncryptAES([]byte(*data.Content), []byte(*data.Key))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Encode the encrypted message as base64
		encodedMessage := base64.StdEncoding.EncodeToString(encryptedMessage)

		// Send the encrypted message in the response
		response := struct {
			EncryptedMessage string `json:"encrypted_message"`
		}{
			EncryptedMessage: encodedMessage,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

// DecryptionAESHandler handles requests for AES decryption.
func DecryptionAESHandler(repo *database.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse request body
		var data model.Data
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Decode the encrypted message from base64
		encryptedMessage, err := base64.StdEncoding.DecodeString(*data.Content)
		if err != nil {
			http.Error(w, "Invalid base64 encoded message", http.StatusBadRequest)
			return
		}

		// Decrypt the message
		decryptedMessage, err := DecryptAES(encryptedMessage, []byte(*data.Key))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Send the decrypted message in the response
		response := struct {
			DecryptedMessage string `json:"decrypted_message"`
		}{
			DecryptedMessage: string(decryptedMessage),
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}