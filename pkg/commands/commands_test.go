package commands

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestParseCommands(t *testing.T) {
    welcomeCommand := "0 1\n"
    cmd := ParseCmd(welcomeCommand)

    assert.Equal(t, Command{Command: WELCOME, Body: welcomeCommand}, cmd, "They should be equal structs")
}
