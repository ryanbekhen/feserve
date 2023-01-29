package loadbalancer

import (
	"sync"
)

type RoundRobin struct {
	sync.Mutex

	current int
	pool    []string
}

func NewRoundRobin(pool []string) *RoundRobin {
	return &RoundRobin{
		current: 0,
		pool:    pool,
	}
}

func (r *RoundRobin) Get() string {
	r.Lock()
	defer r.Unlock()

	if r.current >= len(r.pool) {
		r.current = r.current % len(r.pool)
	}

	result := r.pool[r.current]
	r.current++
	return result
}
