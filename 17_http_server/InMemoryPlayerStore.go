package main

import "sync"

type InMemoryPlayerStore struct {
	store map[string]int
	mux   sync.RWMutex
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}, sync.RWMutex{}}
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) (int, bool) {
	i.mux.Lock()
	defer i.mux.Unlock()

	score, ok := i.store[name]

	return score, ok
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.mux.Lock()
	defer i.mux.Unlock()

	i.store[name]++
}
