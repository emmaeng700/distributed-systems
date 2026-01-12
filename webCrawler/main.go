package main

import "fmt"
import "sync"


type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

func Crawl(url string, depth int, fetcher Fetcher) {
	var mut sync.Mutex
	var wg sync.WaitGroup
	visited := make(map[string]bool)

	var crawl func (url string, depth int, fetcher Fetcher)
	crawl = func (url string, depth int, fetcher Fetcher)  {
		if depth <= 0 {
			return
		}
		
		mut.Lock()
		if visited[url] {
			mut.Unlock()
			return
		}

		mut.Unlock()

		body, urls, err := fetcher.Fetch(url)
		if err != nil {
			return
		}

		fmt.Printf("found: %s %q\n", url, body)

		for _, u := range urls {
			wg.Add(1)
			do := func ()  {
				defer wg.Done()
				crawl(u, depth-1, fetcher)
			}
			go do()
		}
	}
	crawl(url, depth, fetcher)
	wg.Wait()
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
}

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
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
