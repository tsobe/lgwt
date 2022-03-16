package context

import (
	"context"
	"fmt"
	"net/http"
)

type Store interface {
	Fetch(ctx context.Context) (string, error)
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fetch, err := store.Fetch(r.Context())
		if err == nil {
			fmt.Fprint(w, fetch)
		}
	}
}
