package utils

import "testing"

func TestNewID(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(NewID())
	}
}

func TestHash(t *testing.T) {
	t.Log(Hash("Hello"))
	t.Log(Hash("World"))
}
