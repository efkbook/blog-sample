package db

import (
	"strings"
	"testing"
)

func TestReadConfig(t *testing.T) {
	r := strings.NewReader(`
development:
  datasource: root@localhost/dev

test:
  datasource: root@localhost/test
`)

	configs, err := NewConfigs(r)
	if err != nil {
		t.Fatalf("read config failed: %s", err)
	}

	c, ok := configs["development"]
	if !ok {
		t.Fatal("cannot read development configuration.")
	}

	if ex := "root@localhost/dev"; c.DSN() != ex {
		t.Fatalf("want %s, got %s", ex, c.DSN())
	}
}
