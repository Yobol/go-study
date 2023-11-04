package reverse

import "testing"

func TestString(t *testing.T) {
	for _, c := range []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	} {
		got := String(c.in)
		if got != c.want {
			t.Errorf("String(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
