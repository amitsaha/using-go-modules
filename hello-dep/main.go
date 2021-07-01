package main

import (
    "fmt"
    "os"
    "io"

    "github.com/amitsaha/using-go-modules/greetings/hello"
    "github.com/amitsaha/using-go-modules/greetings/world"
)

func displayGreetings(w io.Writer) {
    fmt.Fprintln(w, hello.Greet())
    fmt.Fprintln(w, world.Greet())
}

func main() {   
    displayGreetings(os.Stdout)
}
