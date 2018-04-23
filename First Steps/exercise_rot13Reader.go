package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

// ROT13 Cipher: https://en.wikipedia.org/wiki/ROT13
// Other solutions here: https://gist.github.com/flc/6439105
func (rot13 rot13Reader) Read(b []byte) (int, error) {
	bytes := make([]byte, len(b), cap(b))
	n, err := rot13.r.Read(bytes)

	if n > 0 && err == nil {
		for i := 0; i < n; i++ {
			char := bytes[i]
			if char >= 'a' && char <= 'm' || char >= 'A' && char <= 'M' {
				char += 13
			} else if char > 'm' && char <= 'z' || char > 'M' && char <= 'Z' {
				char -= 13
			}

			b[i] = char
		}
	}

	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
