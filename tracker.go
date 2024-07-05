package main

import (
	"sync"
)

type VisitorTracker struct {
	mu     sync.RWMutex
	visits map[string]map[string]struct{}
}

func NewVisitorTracker() *VisitorTracker {

	return &VisitorTracker{
		visits: make(map[string]map[string]struct{}),
	}
}

func (vt *VisitorTracker) AddVisit(url string, visitorName string) {
	vt.mu.Lock()
	defer vt.mu.Unlock()

	if _, exists := vt.visits[url]; !exists {
		vt.visits[url] = make(map[string]struct{})
	}
	vt.visits[url][visitorName] = struct{}{}
}

func (vt *VisitorTracker) GetUniqueVisitors(url string) int {
	vt.mu.RLock()
	defer vt.mu.RUnlock()

	return len(vt.visits[url])
}

func (vt *VisitorTracker) GetAllVisitors() map[string]int {
	vt.mu.RLock()
	defer vt.mu.RUnlock()

	result := make(map[string]int)
	for url, visitors := range vt.visits {
		result[url] = len(visitors)
	}
	return result
}
