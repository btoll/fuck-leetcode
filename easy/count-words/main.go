package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"unicode"

	"github.com/pkg/profile"
)

func readbyte(reader *bufio.Reader) (rune, error) {
	var buf [1]byte
	_, err := reader.Read(buf[:])
	return rune(buf[0]), err
}

func main() {
	defer profile.Start(
		//		profile.CPUProfile,
		profile.MemProfile,
		profile.MemProfileRate(1),
		profile.ProfilePath("."),
	).Stop()

	file, _ := os.Open(os.Args[1])
	var inword bool
	var words int
	reader := bufio.NewReader(file)
	for {
		r, err := readbyte(reader)
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			log.Fatalf("err=%+v\n", err)
		}
		if unicode.IsSpace(r) {
			inword = false
		}
		if !inword && unicode.IsLetter(r) {
			words += 1
			inword = true
		}
	}
	fmt.Println(words)
}
