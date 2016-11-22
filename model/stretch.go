package model

import (
	"crypto/sha256"
	"encoding/base64"
	"io"
	"math/rand"
	"time"
)

var runes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Stretch makes stretched password using salt.
func Stretch(password, salt string) string {
	var b []byte
	s := sha256.New()
	for i := 0; i < 1000; i++ {
		io.WriteString(s, string(b)+password+salt)
		b = s.Sum(nil)
	}
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
