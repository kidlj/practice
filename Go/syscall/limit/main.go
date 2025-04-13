package main

import (
	"fmt"
	"syscall"
)

// rlimitMaxNumFiles returns hard limit for RLIMIT_NOFILE
func rlimitMaxNumFiles() int {
	var rLimit syscall.Rlimit
	err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	if err != nil {
		fmt.Printf("Error reading system maximum number of open file descriptors (RLIMIT_NOFILE): %v", err)
		return 0
	}
	fmt.Printf("rlimit.max=%v", rLimit.Max)
	return int(rLimit.Max)
}

func main() {
	limit := rlimitMaxNumFiles()
	fmt.Println(limit)
}
