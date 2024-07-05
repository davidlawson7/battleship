package main

import (
	"fmt"
	"log/slog"
	"net"
	"os"
)

func main() {
    conn, err := net.Dial("tcp", ":6446")
    if err != nil {
        slog.Error("cannot connect to server")
        os.Exit(1)
    }
 
    cmdByte := make([]byte, 8)
    n, err:= conn.Read(cmdByte)

    if err != nil {
        slog.Error("unable to read byte")
        conn.Close()
        os.Exit(1)
    }

	fmt.Printf("client cmdByte[:n] = %q\n", cmdByte[:n])
    conn.Close()
}

