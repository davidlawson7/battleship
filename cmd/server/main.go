package main

import (
	"fmt"
	"io"
	"log/slog"
	"net"
)

type Connection struct {
    Reader io.Reader
    Writer io.Writer
    Id int
    conn net.Conn
}

func NewConnection(conn net.Conn, id int) Connection {
    return Connection{
        Reader: conn,
        Writer: conn,
        Id: id,
        conn: conn,
    }
}

func main() {
    ln, err := net.Listen("tcp", ":6446")
    if err != nil {
        slog.Error("Unable to start server on port 6446")
        return
    }

    id := 0
    for {
        conn, err := ln.Accept()
        if err != nil {
            slog.Error("Unable to accept a connection attempt")
            continue
        }
        defer conn.Close()
        
        id++

        newConn := NewConnection(conn, id)
        // Welcome message to client
        conn.Write([]byte{'w'})
        
        slog.Debug("New Connnection", "id", newConn.Id)
        go handleClientConnection(&newConn)
    }
}

func handleClientConnection(conn *Connection) {
    fmt.Println("Handling conn")
    
    // read from the user
    cmdByte, err := conn.Reader.Read
    if err != nil {
        slog.Error("Unable to read from reader")
    }

    fmt.Print(cmdByte)
    conn.conn.Close()
}

