package store

import "testing"

func TestStore(t *testing.T) {
    s := New()
    s.Set("k", "v")
    if v, ok := s.Get("k"); !ok || v != "v" {
        t.Fatal("unexpected value")
    }
}
