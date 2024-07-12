package util

import (
	"crypto/rand"
	"fmt"
	"runtime"

	"github.com/alexedwards/argon2id"
)

var argon2Params = &argon2id.Params{
	Memory:      128 * 1024,
	Iterations:  10,
	Parallelism: uint8(runtime.NumCPU()),
	SaltLength:  16,
	KeyLength:   32,
}

func GenerateRandomBytes(length int) ([]byte, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func GenerateRandomHex(length int) (string, error) {
	bytes, err := GenerateRandomBytes(length)
	if err != nil {
		return "", err
	}
	return BytesToHex(bytes), nil
}

func BytesToHex(bytes []byte) string {
	return fmt.Sprintf("%x", bytes)
}

func EncryptPassword(password string) (string, error) {
	return argon2id.CreateHash(password, argon2Params)
}

func VerifyPassword(password, encryptedPassword string) (bool, error) {
	return argon2id.ComparePasswordAndHash(password, encryptedPassword)
}
