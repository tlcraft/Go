package main

import (
	"fmt"
	"runtime"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

var urlMap = SafeUrlMap{found: make(map[string]bool)}
var numCPU = runtime.NumCPU()

type SafeUrlMap struct {
	found map[string]bool
	mux   sync.Mutex
}

func (m *SafeUrlMap) Add(url string) {
	m.mux.Lock()
	m.found[url] = true
	m.mux.Unlock()
}

func (m *SafeUrlMap) Contains(url string) bool {
	m.mux.Lock()
	defer m.mux.Unlock()
	return m.found[url]
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, c chan string) {
	// Stack Overflow answer: https://stackoverflow.com/a/13223836/8094831
	// Parallelization section of Effective Go: https://golang.org/doc/effective_go.html#parallel
	defer close(c)

	if depth <= 0 || urlMap.Contains(url) {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		c <- err.Error()
		return
	}

	c <- fmt.Sprintf("found: %s %q", url, body)

	urlMap.Add(url)

	result := make([]chan string, len(urls))
	for i, u := range urls {
		result[i] = make(chan string)
		go Crawl(u, depth-1, fetcher, result[i])
	}

	for i := range result {
		for s := range result[i] {
			c <- s
		}
	}

	return
}

func main() {
	//fmt.Println("NumCPU", numCPU)
	fmt.Println("Fetcher len:", len(fetcher))

	c := make([]chan string, len(fetcher))
	i := 0
	for k := range fetcher {
		c[i] = make(chan string)
		go Crawl(k, 4, fetcher, c[i])
		i++
	}

	for i := range c {
		for s := range c[i] {
			fmt.Println(s)
		}
	}
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
			"https://golang.org/test/t",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
			"https://golang.org/test/",
		},
	},
	"https://golang.org/test/": &fakeResult{
		"Test Package",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/test/t": &fakeResult{
		"Test Package 2.0",
		[]string{
			"https://golang.org/",
			"https://golang.org/test/",
		},
	},
}
