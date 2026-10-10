package docker

import "testing"

func TestCapabilitiesDescribeBackendFeatures(t *testing.T) {
	if got := (&Driver{}).capabilities(); len(got) != 1 || got[0] != "cdi" {
		t.Fatalf("default capabilities=%v", got)
	}
	d := &Driver{runtimeConfig: RuntimeConfig{Capabilities: []string{"cdi", "libpod-native"}}}
	got := d.capabilities()
	got[0] = "mutated"
	if d.runtimeConfig.Capabilities[0] != "cdi" {
		t.Fatal("capabilities leaked runtime config backing storage")
	}
}

func TestClassifyIsolation(t *testing.T) {
	cases := []struct {
		host, os, osType, want string
	}{
		{"darwin", "Docker Desktop", "linux", "docker-desktop-vm"},
		{"windows", "Docker Desktop", "linux", "docker-desktop-vm"},
		{"linux", "Ubuntu 24.04", "linux", "native-linux"},
		{"linux", "", "linux", "native-linux"},
	}
	for _, tc := range cases {
		got := classifyIsolation(tc.host, tc.os, tc.osType)
		if got != tc.want {
			t.Errorf("classifyIsolation(%q,%q,%q)=%q want %q", tc.host, tc.os, tc.osType, got, tc.want)
		}
	}
}
