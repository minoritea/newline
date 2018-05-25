package main

import (
	"io"
	"log"
	"os"
)

type lastbyteholder byte

func (l *lastbyteholder) Write(data []byte) (int, error) {
	if l == nil {
		l = new(lastbyteholder)
	}
	*l = lastbyteholder(data[len(data)-1])
	return len(data), nil
}

func (l *lastbyteholder) hasNewlineAtEnd() bool {
	return l == nil || byte(*l) == 0x0a
}

func main() {
	var b *lastbyteholder
	if _, err := io.Copy(b, io.TeeReader(os.Stdin, os.Stdout)); err != nil {
		log.Fatal(err)
	}

	if b.hasNewlineAtEnd() {
		return
	}
	if _, err := os.Stdout.Write([]byte{0x0a}); err != nil {
		log.Fatal(err)
	}
}
