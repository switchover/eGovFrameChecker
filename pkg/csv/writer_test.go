package csv

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNewWriterWritesHeader(t *testing.T) {
	dir := t.TempDir()
	file := filepath.Join(dir, "test.csv")
	w, err := NewWriter(file, []string{"a", "b"})
	if err != nil {
		t.Fatalf("NewWriter returned error: %v", err)
	}
	// header should be flushed immediately
	content, err := os.ReadFile(file)
	if err != nil {
		t.Fatalf("reading file: %v", err)
	}
	expected := "a,b\n"
	if string(content) != expected {
		t.Fatalf("unexpected content: got %q, want %q", string(content), expected)
	}
	if err := w.Close(); err != nil {
		t.Fatalf("Close returned error: %v", err)
	}
}
