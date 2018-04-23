package main

import (
	//"fmt"
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
	bytes := make([]byte, 1)
	var ind int = 0

loop:
	for {
		n, err := rot13.r.Read(bytes)
		//fmt.Printf("err=%v n=%v io.EOF=%v\n", err, n, io.EOF)
		//fmt.Printf("b[]=%v", b)
		if n == 0 || err == io.EOF {
			break loop
		}

		if n > 0 && err == nil {
			char := bytes[0]
			if char >= 'a' && char <= 'm' || char >= 'A' && char <= 'M' {
				char += 13
			} else if char > 'm' && char <= 'z' || char > 'M' && char <= 'Z' {
				char -= 13
			}

			b[ind] = char
			ind++
		}
	}

	return ind, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
