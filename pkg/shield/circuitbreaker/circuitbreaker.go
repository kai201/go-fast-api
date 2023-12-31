// Package circuitbreaker is an adaptive circuit breaker library, support for use in gin middleware and grpc interceptors.
package circuitbreaker

import (
	"errors"
)

// ErrNotAllowed error not allowed.
var ErrNotAllowed = errors.New("circuitbreaker: not allowed for circuit open")

// CircuitBreaker is a circuit breaker.
type CircuitBreaker interface {
	Allow() error
	MarkSuccess()
	MarkFailed()
}
