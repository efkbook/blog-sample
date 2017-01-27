package model

import "testing"

func TestStretch(t *testing.T) {
	password := "test"
	salt := "塩"
	want := "x2SLvhmxOaV2enRmd678M2VFkwZBmYKuHvU369oGoKI="
	s := Stretch(password, salt)
	if exp := want; s != exp {
		t.Fatalf("want %s, got %s", exp, s)
	}
}
