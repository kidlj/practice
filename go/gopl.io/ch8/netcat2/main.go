package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	c, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	// After the user types ^D at the terminal, the program stops,
	// even if the other goroutine still has work to do.
	// If we switch the order of these two goroutines, this won't work.
	go mustCopy(os.Stdout, c)
	mustCopy(c, os.Stdin)
}

func mustCopy(w io.Writer, r io.Reader) {
	// Copy copies from src to dst until either EOF is reached
	// on src or an error occurs.
	_, err := io.Copy(w, r)
	if err != nil {
		log.Fatal(err)
	}
}
