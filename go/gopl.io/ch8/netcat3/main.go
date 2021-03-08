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

	done := make(chan struct{})

	go func() {
		// Ignoring errors on purpose.
		// This blocks.
		_, err := io.Copy(os.Stdout, conn)
		log.Printf("Error: %v", err)
		log.Println("done")
		done <- struct{}{}
	}()
	mustCopy(conn, os.Stdin)
	// Can't be deferred.
	// Can't be the last statement.
	conn.Close()
	<-done
}

func mustCopy(w io.Writer, r io.Reader) {
	_, err := io.Copy(w, r)
	if err != nil {
		log.Fatal(err)
	}
}
