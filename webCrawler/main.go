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