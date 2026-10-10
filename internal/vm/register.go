package vm

import "github.com/cautem/cauteum-driver/driver"

func init() {
	driver.Register("vm", func() (driver.ComputeDriver, error) {
		return New(), nil
	})
}
