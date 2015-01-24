package config

import (
	"sync"
)

var (
	secretKey   string
	secretMutex = &sync.RWMutex{}
)

// SetSecret sets a new secret key (thread-safe)
func SetSecret(s string) {
	secretMutex.Lock()
	secretKey = s
	secretMutex.Unlock()
}

// Returns the secret key that has been set (thread-safe))
func Secret() string {
	secretMutex.RLock()
	defer secretMutex.RUnlock()
	return secretKey
}
