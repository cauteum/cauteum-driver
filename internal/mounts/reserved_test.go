package mounts

import "testing"

func TestValidateContainerMountTarget(t *testing.T) {
	cases := []struct {
		target string
		ok     bool
	}{
		{"/workspace/src", true},
		{"/data", true},
		{"/whaleshell", false},
		{"/whaleshell/data", false},
		{"/whaleshell/policy.yaml", false},
		{"/proc", false},
		{"/sys/fs", false},
		{"/dev/null", false},
		{"/run/whaleshell/ssh.sock", false},
		{"/run", false},
		{"/etc/whaleshell", false},
		{"/var/run/whaleshell", false},
		{"relative", false},
		{"/workspace/../whaleshell", false}, // has ..
	}
	for _, tc := range cases {
		err := ValidateContainerMountTarget(tc.target)
		if tc.ok && err != nil {
			t.Fatalf("%s: unexpected err %v", tc.target, err)
		}
		if !tc.ok && err == nil {
			t.Fatalf("%s: expected error", tc.target)
		}
	}
}

func TestValidateUploadDest(t *testing.T) {
	if err := ValidateUploadDest("/workspace/file.txt"); err != nil {
		t.Fatal(err)
	}
	if err := ValidateUploadDest("/workspace"); err != nil {
		t.Fatal(err)
	}
	if err := ValidateUploadDest("/whaleshell/data/x"); err == nil {
		t.Fatal("expected refuse /whaleshell")
	}
}

func TestPathsOverlap(t *testing.T) {
	if !PathsOverlap("/whaleshell", "/whaleshell/data") {
		t.Fatal("expected overlap")
	}
	if PathsOverlap("/workspace", "/whaleshell") {
		t.Fatal("no overlap")
	}
}
