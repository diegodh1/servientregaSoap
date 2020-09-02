package connection

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

type ConnectionInterface interface {
	Connect() (string, error)
	Disconnect()
	GetServer() string
	GetPort() int
	GetUser() string
	GetPassword() string
	GetDatabase() string
	GetDBConnection() *sql.DB
	GetContext() context.Context
	SetServer(value string)
	SetPort(value int)
	SetUser(value string)
	SetPassword(value string)
	SetDatabase(value string)
	SetDBConnection(value *sql.DB)
}
type Connection struct {
	server   string
	port     int
	user     string
	password string
	database string
	db       *sql.DB
	context  context.Context
}

func (C *Connection) GetServer() string {
	return C.server
}
func (C *Connection) GetPort() int {
	return C.port
}
func (C *Connection) GetUser() string {
	return C.user
}
func (C *Connection) GetPassword() string {
	return C.database
}
func (C *Connection) GetDBConnection() *sql.DB {
	return C.db
}
func (C *Connection) GetContext() context.Context {
	return C.context
}
func (C *Connection) SetServer(value string) {
	C.server = value
}
func (C *Connection) SetPort(value int) {
	C.port = value
}
func (C *Connection) SetUser(value string) {
	C.user = value
}
func (C *Connection) SetPassword(value string) {
	C.password = value
}
func (C *Connection) SetDBConnection(value *sql.DB) {
	C.db = value
}

func (C *Connection) Disconnect() {
	C.db.Close()
}

func (C *Connection) Connect() (string, error) {
	C.server = "172.16.5.3"
	C.port = 1433
	C.user = "sa"
	C.password = "AdminSQL.2019$"
	C.database = "Integrapps"
	var err error
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		C.server, C.user, C.password, C.port, C.database)
	// Create connection pool
	C.db, err = sql.Open("sqlserver", connString)
	if err != nil {
		return "", errors.New("Error creating connection pool: " + err.Error())
	}
	C.context = context.Background()
	err = C.db.PingContext(C.context)
	if err != nil {
		return "", errors.New("Error creating connection pool: " + err.Error())
	}
	return "Connected!\n", nil
}
