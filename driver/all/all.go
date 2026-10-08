// Package all blank-imports compute backends so they register with driver.Open.
//
//	import _ "github.com/cauteum/cauteum-driver/driver/all"
package all

import (
	_ "github.com/cauteum/cauteum-driver/internal/docker"
	_ "github.com/cauteum/cauteum-driver/internal/kubernetes"
	_ "github.com/cauteum/cauteum-driver/internal/podman"
	_ "github.com/cauteum/cauteum-driver/internal/vm"
)
