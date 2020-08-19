package part1

import "testing"

func TestUnpack(t *testing.T) {
	a := "a4bc2d5e"
	b := "abcd"
	c := "45"
	e := "aaaabccddddde"
	f := "abcd"
	g := ""
	if ct := unpack(a); ct != e {
		t.Fatalf("wrong unpacking for %s: got %s expected %s", a, ct, e)
	}
	if ct := unpack(b); ct != f {
		t.Fatalf("wrong unpacking for %s: got %s expected %s", b, ct, f)
	}
	if ct := unpack(c); ct != g {
		t.Fatalf("wrong unpacking for %s: got %s expected %s", c, ct, g)
	}
}
