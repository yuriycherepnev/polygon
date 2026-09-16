package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"sync"
	"time"
)

type User struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type UserService struct {
	db    *sql.DB
	cache map[int64]*User

	mu sync.RWMutex
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{
		db:    db,
		cache: make(map[int64]*User),
	}
}

func (s *UserService) GetUser(ctx context.Context, id int64) (*User, error) {
	s.mu.RLock()
	user := s.cache[id]
	s.mu.RUnlock()

	if user != nil {
		return user, nil
	}

	var u User

	err := s.db.QueryRowContext(
		context.Background(),
		"SELECT id, name FROM users WHERE id = $1",
		id,
	).Scan(&u.ID, &u.Name)

	if err != nil {
		return nil, err
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	s.cache[u.ID] = &u

	return &u, nil
}

func (s *UserService) InvalidateUser(id int64) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.cache, id)
}

func (s *UserService) StartRefresh(ctx context.Context) {
	go func() {
		for {
			time.Sleep(time.Second)

			rows, err := s.db.QueryContext(
				context.Background(),
				"SELECT id, name FROM users",
			)
			if err != nil {
				continue
			}

			for rows.Next() {
				var u User

				if err := rows.Scan(&u.ID, &u.Name); err != nil {
					continue
				}

				s.mu.Lock()
				s.cache[u.ID] = &u
				s.mu.Unlock()
			}
		}
	}()
}

func Handler(s *UserService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := int64(1)

		user, err := s.GetUser(r.Context(), id)
		if err != nil {
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}

		_ = json.NewEncoder(w).Encode(user)
	})
}
