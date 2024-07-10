package main

import (
	"bufio"
	"fmt"
	"io"
	"log/slog"
	"net"
	"os"
	"strings"

	"github.com/davidlawson7/battleship/pkg/commands"
	"github.com/davidlawson7/battleship/pkg/tcp"
)

func main() {
    conn, err := net.Dial("tcp", ":6446")
    if err != nil {
        slog.Error("cannot connect to server")
        os.Exit(1)
    }
    defer conn.Close()
 
    clientConn := tcp.NewClientConnection(conn, 0)
    cmd, err := clientConn.Reader.ReadString('\n')

    if err != nil {
        slog.Error("unable to read from conn")
        os.Exit(1)
    }
    
    command := commands.ParseCmd(cmd)

    if command.Command != commands.WELCOME {
        slog.Error("We werent welcomes noob")
        os.Exit(1)
    }

    fmt.Println("Body", command.Body)

    //if len(cmd) >= 9 && cmd[:7] == commands.WELCOME {
    //    id, err := strconv.Atoi(cmd[8:len(cmd) - 1])
    //    if err != nil {
    //        fmt.Println("AH SHIT")
    //        os.Exit(1)
    //    } else {
    //        newConn := tcp.NewConnection(conn, id)
    //        go listener(&newConn) 
    //    }
    //}

    reader := bufio.NewReader(os.Stdin)
    // Spin up a routine to handle taking and sending user input, main routine
    // can handle receiving input
    for {
        fmt.Print("->")
        text, _ := reader.ReadString('\n')
        text = strings.Replace(text, "\n", "", -1)
        fmt.Printf("Said %s\n", text)

        conn.Write([]byte(commands.CommandNameLookup[commands.START]))
    }
}

func handleWelcomeMessage() {
    
}

func listener(conn *tcp.Connection) {    
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

		fmt.Printf("client listener[:n] = %d\n", cmdByte[:n])

    }
    slog.Info("Closing connection")
    conn.Conn.Close()
}


