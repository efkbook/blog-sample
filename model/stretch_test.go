package model

import "testing"

func TestStretch(t *testing.T) {
	password := "test"
	salt := "塩"
	s := Stretch(password, salt)
	if exp := "Y9SA/lgzzns9Tjt474at6mAC7YXhnSeYt5tjxd4BPrM="; s != exp {
		t.Fatalf("want %s, got %s", exp, s)
	}
}
