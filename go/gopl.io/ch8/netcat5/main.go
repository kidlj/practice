package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan struct{})
	go func() {
		mustCopy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{}
	}()

	mustCopy(conn, os.Stdin)
	// A TCP connection consists of two halves that maybe closed independently using its CloseRead and CloseWrite methods
	// Closing the write half of the connection causes the server to see an end-of-file condition
	// Closing the read half causes the background goroutine’s call to io.Copy to return a "read from closed connection" error
	// You can't close a remote connection
	conn.CloseWrite()
	<-done
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		// No need since we only close the write half of the connection.
		// if err == io.EOF {
		// 	return
		// }
		log.Fatal(err)
	}
}
