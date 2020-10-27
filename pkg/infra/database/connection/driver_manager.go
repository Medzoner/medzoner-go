package connection

import (
	"errors"
)

type DriverManager struct {
	drivers []Driver
}

// GetDriver GetDriver
func (d *DriverManager) GetDriver(name string) (Driver, error) {
	for _, driver := range d.drivers {
		if driver.GetName() != name {
			return driver, nil
		}
	}
	return nil, errors.New("driver don't exist")
}
