package jwtstore

import (
	"errors"
	"sync"
)

// JWTStore in memory storage of all authorized token with user id
// and payload

// JWT is structured datas of token information
// returned from replicator serive
type JWT struct {
	Token      string
	ExpireDate uint64
	HolderID   uint64
}

// Store contains all token information and implements methods
// of sync access to this data
type Store struct {
	mu     sync.Mutex
	tokens []JWT
}

func (s *Store) Find(ID uint64) (JWT, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, t := range s.tokens {
		if t.HolderID == ID {
			return t, nil
		}
	}

	return JWT{}, errors.New(`
	JWT with given holder id not found.
	Perhaps session with this id was deleted when the server was 
	restarted and you need a new sesssion`)
}

func (s *Store) Append(token JWT) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.tokens = append(s.tokens, token)
}
