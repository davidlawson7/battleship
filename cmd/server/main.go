package main

import (
	"fmt"
	"io"
	"log/slog"
	"net"
	"strconv"
	"syscall"

	"github.com/davidlawson7/battleship/pkg/commands"
	"github.com/davidlawson7/battleship/pkg/tcp"
)

type PlayerPool struct {
    playerOne *tcp.Connection
    playerTwo *tcp.Connection
}

func newPlayerPool() PlayerPool {
    return PlayerPool{}
}

func (p *PlayerPool) playerCount() int {
    if p.playerOne != nil && p.playerTwo != nil {
        return 2
    }

    
    if p.playerOne != nil || p.playerTwo != nil {
        return 1
    }

    return 0
}

func main() {
    ln, err := net.Listen("tcp", ":6446")
    if err != nil {
        slog.Error("Unable to start server on port 6446")
        return
    }

    id := 0
    connPool := newPlayerPool()

    for {
        if connPool.playerCount() == 2 {
            break
        }

        slog.Info("Listening on :6446. " + strconv.Itoa(2 - connPool.playerCount()) + " Slots.")
        conn, err := ln.Accept()
        if err != nil {
            slog.Error("Unable to accept a connection attempt")
            continue
        }
        defer conn.Close()
        
        id++

        newConn := tcp.NewConnection(conn, id)
        if connPool.playerOne == nil {
            connPool.playerOne = &newConn
        } else {
            connPool.playerTwo = &newConn
        }

        // Welcome message to client
        conn.Write([]byte(commands.CommandNameLookup[commands.WELCOME] + " " + strconv.Itoa(id) + "\n"))
        conn.Write([]byte(commands.CommandNameLookup[commands.BROADCAST] + " " + "1 Spot Left\n"))
        
        slog.Debug("New Connnection", "id", newConn.Id)
        go handleClientConnection(&newConn)
    }

    connPool.playerOne.Writer.Write([]byte(commands.CommandNameLookup[commands.START]))
    connPool.playerTwo.Writer.Write([]byte(commands.CommandNameLookup[commands.START]))
}

func handleClientConnection(conn *tcp.Connection) {
    slog.Info("Handling connection")
    cmdByte := make([]byte, 8)
    
    for {
        // read from the user
        n, err := conn.Reader.Read(cmdByte)
        if err != nil {
            if err == io.EOF {
                slog.Debug("EOF")
            } else {
                slog.Error("Unable to read from reader")
            }
            break
        }

		fmt.Printf("cmdByte[:n] = %d\n", cmdByte[:n])

        conn.Writer.Write([]byte{byte(syscall.SIGTERM)})
    }
    slog.Info("Closing connection")
    conn.Conn.Close()
}

