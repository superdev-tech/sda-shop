package database

import (
	"errors"
	"strings"
	"testing"
)

func TestHint(t *testing.T) {
	cases := []struct {
		err  string
		want string
	}{
		{"FATAL: password authentication failed for user", "รหัสผ่าน"},
		{"dial tcp [2406:da18::1]:5432: connect: network is unreachable", "Session pooler"},
		{"something else", "3 ข้อ"},
	}
	for _, c := range cases {
		if got := Hint(errors.New(c.err)); !strings.Contains(got, c.want) {
			t.Errorf("Hint(%q) = %q, want it to mention %q", c.err, got, c.want)
		}
	}
}
