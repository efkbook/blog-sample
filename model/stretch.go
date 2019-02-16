package model

import (
	"crypto/sha256"
	"encoding/base64"
	"math/rand"
	"time"

	"golang.org/x/crypto/pbkdf2"
)

var runes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Stretch makes stretched password using salt.
func Stretch(password, salt string) string {
	b := pbkdf2.Key([]byte(password), []byte(salt), 16384, 32, sha256.New)
	return base64.StdEncoding.EncodeToString(b)
}

// Salt returns random salt string.
func Salt(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = runes[rand.Intn(len(runes))]
	}
	return string(b)
}
