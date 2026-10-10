// Package all blank-imports compute backends so they register with driver.Open.
//
//	import _ "github.com/cauteum-haven/cauteum-driver/driver/all"
package all

import (
	_ "github.com/cauteum-haven/cauteum-driver/internal/docker"
	_ "github.com/cauteum-haven/cauteum-driver/internal/kubernetes"
	_ "github.com/cauteum-haven/cauteum-driver/internal/podman"
	_ "github.com/cauteum-haven/cauteum-driver/internal/vm"
)
