package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("%s: file name not given\n", os.Args[0])
		os.Exit(1)
	}

	for i := 1; i < len(os.Args); i++ {
		doCat(os.Args[i])
	}
}

func doCat(path string) {
	const BufferSize = 2048
	var buf = make([]byte, BufferSize)

	fd, err := syscall.Open(path, os.O_RDONLY, 0666)
	if err != nil {
		die(path)
	}
	for {
		n, err := syscall.Read(fd, buf)
		if err != nil {
			die(path)
		}
		if n == 0 {
			break
		}
		m, err := syscall.Write(syscall.Stdout, buf)
		if m < 0 {
			die(path)
		}
	}
	if syscall.Close(fd) != nil {
		die(path)
	}
}

func die(s string) {
	os.Stderr.Write([]byte(s))
	os.Exit(1)
}
