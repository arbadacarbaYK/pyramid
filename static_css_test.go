package main

import (
	"bytes"
	"testing"
)

func TestDashboardCSSEmbedded(t *testing.T) {
	if len(requiredStyles) < 1000 {
		t.Fatal("static/styles.css missing or tiny; run `just tailwind` before go test")
	}
	if !bytes.Contains(requiredStyles, []byte("--tw")) && !bytes.Contains(requiredStyles, []byte(".flex")) {
		t.Fatal("static/styles.css does not look like Tailwind output")
	}
}
