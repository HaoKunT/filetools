package tools

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParWalker(t *testing.T) {
	info, err := os.Stat("../")
	if err != nil {
		t.Fatal(err)
	}
	absPath, err := filepath.Abs("../")
	if err != nil {
		t.Error(err)
	}
	rpath, err := filepath.Rel("../", "../")
	if err != nil {
		t.Error(err)
	}
	walk := ParWalker{
		info,
		absPath,
		rpath,
		err,
	}
	t.Logf("Path: %s\n", walk.AbsPath)
}

func TestParWalk_10(t *testing.T) {
	for walk := range ParWalk("../", 10) {
		if walk.Err != nil {
			t.Errorf("%s\n", walk.Err)
			continue
		}
		t.Logf("Path: %s\n", walk.AbsPath)
	}
	for walk := range ParWalk("../", -10) {
		if walk.Err != nil {
			t.Errorf("%s\n", walk.Err)
			continue
		}
		t.Logf("Path: %s\n", walk.AbsPath)
	}
}

func BenchmarkParWalk_1000(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _ = range ParWalk("../", 1000) {
		}
	}
}

func BenchmarkParWalk_1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _ = range ParWalk("../", 1) {
		}
	}
}
