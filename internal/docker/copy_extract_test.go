package docker

import (
	"archive/tar"
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func testCopyArchive(t *testing.T, name string, mode int64) *tar.Reader {
	t.Helper()
	var buf bytes.Buffer
	w := tar.NewWriter(&buf)
	data := []byte("contents")
	if err := w.WriteHeader(&tar.Header{Name: name, Typeflag: tar.TypeReg, Mode: mode, Size: int64(len(data))}); err != nil {
		t.Fatal(err)
	}
	if _, err := w.Write(data); err != nil {
		t.Fatal(err)
	}
	if err := w.Close(); err != nil {
		t.Fatal(err)
	}
	return tar.NewReader(&buf)
}

func TestCopyExtractRejectsTraversalAndSymlinks(t *testing.T) {
	root := t.TempDir()
	for _, name := range []string{"../outside", "/tmp/outside", "sub/../../outside"} {
		if err := extractCopyTar(testCopyArchive(t, name, 0777), root, true, "source"); err == nil {
			t.Fatalf("accepted archive path %q", name)
		}
	}
	outside := t.TempDir()
	if err := os.Symlink(outside, filepath.Join(root, "link")); err != nil {
		t.Skipf("symlink unavailable: %v", err)
	}
	if err := extractCopyTar(testCopyArchive(t, "link/file", 0644), root, true, "source"); err == nil {
		t.Fatal("followed existing destination symlink")
	}
}

func TestCopyExtractCapsFileModes(t *testing.T) {
	root := t.TempDir()
	for _, tc := range []struct {
		name string
		mode int64
		want os.FileMode
	}{
		{"data", 0666, 0644},
		{"script", 0777, 0755},
	} {
		if err := extractCopyTar(testCopyArchive(t, tc.name, tc.mode), root, true, "source"); err != nil {
			t.Fatal(err)
		}
		info, err := os.Stat(filepath.Join(root, tc.name))
		if err != nil || info.Mode().Perm() != tc.want {
			t.Fatalf("%s mode: %v, %v", tc.name, info, err)
		}
	}
}
