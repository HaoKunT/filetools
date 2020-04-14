package tools

import (
	"testing"
)

func BenchmarkNewExtFileInfo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewExtInfo("../")
	}
}

func TestNewExtInfo(t *testing.T) {
	info, err := NewExtInfo("../")
	if err != nil {
		t.Fatalf("NewExtInfo failed: %v", err)
	} else {
		t.Logf("NewExtInfo passed: %v", info)
	}
}

func TestExtInfo_Size(t *testing.T) {
	info, err := NewExtInfo("../")
	if err != nil {
		t.Fatalf("NewExtInfo failed: %v", err)
	}
	t.Logf("ExtFileInfo size: %s", transSize(RealSizeOfExtFileinfo(info)))
}
