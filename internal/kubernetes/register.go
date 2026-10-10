package kubernetes

import "github.com/cautem/cauteum-driver/driver"

func init() {
	driver.Register("kubernetes", func() (driver.ComputeDriver, error) {
		return New(), nil
	})
}
