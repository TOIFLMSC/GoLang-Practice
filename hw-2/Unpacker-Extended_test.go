package part1

import "testing"

func TestUnpackExtended(t *testing.T) {
	a1 := "a4bc2d5e"
	b1 := "abcd"
	a2 := `qwe\4\5`
	b2 := `qwe\45`
	c2 := `qwe\\5`
	e1 := "aaaabccddddde"
	f1 := "abcd"
	e2 := `qwe45`
	f2 := `qwe44444`
	g2 := `qwe\\\\\`
	if c := unpackEx(a1); c != e1 {
		t.Fatalf("wrong unpacking for %s: got %s expected %s", a1, c, e1)
	}
	if c := unpackEx(b1); c != f1 {
		t.Fatalf("wrong unpacking for %s: got %s expected %s", b1, c, f1)
	}
	if c := unpackEx(a2); c != e2 {
		t.Fatalf("wrong unpacking for %s: got %s expected %s", a2, c, e2)
	}
	if c := unpackEx(b2); c != f2 {
		t.Fatalf("wrong unpacking for %s: got %s expected %s", b2, c, f2)
	}
	if c := unpackEx(c2); c != g2 {
		t.Fatalf("wrong unpacking for %s: got %s expected %s", c2, c, g2)
	}
}
