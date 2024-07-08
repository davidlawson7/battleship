package tcp

import (
	"io"
	"net"
)

type Connection struct {
    Reader io.Reader
    Writer io.Writer
    Id int
    Conn net.Conn
}

func NewConnection(conn net.Conn, id int) Connection {
    return Connection{
        Reader: conn,
        Writer: conn,
        Id: id,
        Conn: conn,
    }
}

