package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(w io.Writer, r io.Reader) {
	// Copy copies from src to dst until either EOF is reached
	// on src or an error occurs.
	_, err := io.Copy(w, r)
	if err != nil {
		log.Fatal(err)
	}
}
