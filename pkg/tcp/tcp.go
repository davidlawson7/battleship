package tcp

import (
	"bufio"
	"io"
	"net"
)

type ClientConnection struct {
    Conn    net.Conn
    Id      int
    Reader  bufio.Reader
    Writer  bufio.Writer
}

func NewClientConnection(conn net.Conn, id int) ClientConnection {
    return ClientConnection{
        Conn: conn,
        Id: id,
        Reader: *bufio.NewReader(conn),
        Writer: *bufio.NewWriter(conn),
    }
}

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

