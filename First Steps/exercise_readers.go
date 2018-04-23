package main

// Other resources:
// io.Reader in depth: https://medium.com/@matryer/golang-advent-calendar-day-seventeen-io-reader-in-depth-6f744bb4320b
// Demystifying Golang's io.Reader and io.Writer Interfaces: https://nathanleclaire.com/blog/2014/07/19/demystifying-golangs-io-dot-reader-and-io-dot-writer-interfaces/
// Exercise comments: https://www.reddit.com/r/golang/comments/2pa8mx/exercise_readers/

import (
	"golang.org/x/tour/reader"
)

type MyReader struct{}

func (mr MyReader) Read(b []byte) (int, error) {
	for i, _ := range b {
		b[i] = 'A'
		//fmt.Printf("b = %v\n", b)
	}

	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}
