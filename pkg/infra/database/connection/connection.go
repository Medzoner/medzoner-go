package connection

type Connection struct {
	DriverManager DriverManager
}

// Create Create
func (c *Connection) Create(driverName string) (Driver, error) {
	return c.DriverManager.GetDriver(driverName)
}
