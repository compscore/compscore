package helpers

import "golang.org/x/crypto/bcrypt"

// HashPassword hashes a password using bcrypt.
func HashPassword(password string) (string, error) {
	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(passwordBytes), err
}

// CheckPassword checks a password against a hash using bcrypt.
func CheckPassword(password string, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
