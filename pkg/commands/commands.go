package commands

import (
	"fmt"
	"strings"
)

//const (
//   WELCOME     = "WELC"    // WELC <id>
//   START       = "STAR"    // STAR
//   BROADCAST   = "BROA"    // BROA <string message>
//   FIRE        = "FIRE"    // FIRE <col> <row>
//)

const (
    WELCOME = iota
    START
    BROADCAST
    FIRE
    MISSING
)

var CommandMap = map[string]byte {
    "WELCOME": WELCOME,
    "START": START,
    "BROADCAST": BROADCAST,
    "FIRE": FIRE,
    "MISSING": MISSING,
}

var CommandNameLookup = map[byte]string {
    WELCOME: "WELCOME",
    START: "START",
    BROADCAST: "BROADCAST",
    FIRE: "FIRE",
    MISSING: "MISSING",
}

type Command struct {
    Command byte
    Body string
}

func ParseCmd(cmd string) Command {
    fmt.Println(WELCOME)
    strCmd := strings.Split(cmd, " ")[0]

    if c, ok := CommandMap[strCmd]; ok {
        return Command{
            Command: c,
            Body: cmd,
        }
	}
	return Command{
        Command: MISSING,
    }
}

