package connection

type Driver interface {
	GetName() string
	Connect(dsn string, databaseName string)
	ExecuteQuery(query string, body interface{})
}
