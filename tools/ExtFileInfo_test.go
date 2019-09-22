package tools

import (
	"testing"
)

func BenchmarkNewExtFileInfo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewExtInfo("./")
	}
}

func TestNewExtInfo(t *testing.T) {
	info, err := NewExtInfo("./")
	if err != nil {
		t.Fatalf("NewExtInfo failed: %v", err)
	} else {
		t.Logf("NewExtInfo passed: Name: %s, RealSize: %s", info.BasePath, transSize(info.RealSize))
	}
}
