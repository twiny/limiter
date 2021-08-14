package limiter

import (
	"sync"

	"golang.org/x/time/rate"
)

// RateLimiter
type Limiter struct {
	mu      *sync.RWMutex
	rate    rate.Limit
	bursts  int
	clients map[string]*rate.Limiter
}

// NewRateLimiter
func NewLimiter(r float64, b int) *Limiter {
	return &Limiter{
		mu:      &sync.RWMutex{},
		rate:    rate.Limit(r),
		bursts:  b,
		clients: map[string]*rate.Limiter{},
	}
}

// add
func (l *Limiter) Set(ip string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	//
	l.clients[ip] = rate.NewLimiter(l.rate, l.bursts)
}

// get
func (l *Limiter) Get(ip string) (*rate.Limiter, bool) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	//
	rate, found := l.clients[ip]

	return rate, found
}

// purge
func (l *Limiter) Purge() {
	l.mu.RLock()
	defer l.mu.RUnlock()

	l.clients = map[string]*rate.Limiter{}
}
