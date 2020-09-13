package crypto_utils

import "golang.org/x/crypto/bcrypt"

// Hash takes a string and return it hashed
func Hash(input string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(input), 14)
	return string(bytes), err
}

// CheckHash return true if the input and hash are equals, if not, return false
func CheckHash(input, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(input))
	return err == nil
}
