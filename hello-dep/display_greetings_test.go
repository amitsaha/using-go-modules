package main

import (
        "testing"
        "bytes"
)

func TestDisplayGreetings(t *testing.T) {
    buf := new(bytes.Buffer)

    displayGreetings(buf)

    expectedOutput := "Hello!\nWorld!\n"
    gotOutput := buf.String()
    if gotOutput != expectedOutput {
        t.Fatalf("Expected: %s, Got: %s\n", expectedOutput, gotOutput )
    }
}

