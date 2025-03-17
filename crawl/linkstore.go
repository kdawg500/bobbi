package crawl

import "sync"

type LinkCount map[string]int

type LinkStore struct {
	Links map[string]LinkCount
	mu    sync.Mutex
}

func (l *LinkStore) Add(parentLink, link string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.Links[parentLink][link]++
}
