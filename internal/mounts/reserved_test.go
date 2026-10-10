package mounts

import "testing"

func TestValidateContainerMountTarget(t *testing.T) {
	cases := []struct {
		target string
		ok     bool
	}{
		{"/workspace/src", true},
		{"/data", true},
		{"/cauteum", false},
		{"/cauteum/data", false},
		{"/cauteum/policy.yaml", false},
		{"/proc", false},
		{"/sys/fs", false},
		{"/dev/null", false},
		{"/run/cauteum/ssh.sock", false},
		{"/run", false},
		{"/etc/cauteum", false},
		{"/var/run/cauteum", false},
		{"relative", false},
		{"/workspace/../cauteum", false}, // has ..
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
	if err := ValidateUploadDest("/cauteum/data/x"); err == nil {
		t.Fatal("expected refuse /cauteum")
	}
}

func TestPathsOverlap(t *testing.T) {
	if !PathsOverlap("/cauteum", "/cauteum/data") {
		t.Fatal("expected overlap")
	}
	if PathsOverlap("/workspace", "/cauteum") {
		t.Fatal("no overlap")
	}
}
