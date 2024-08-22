package slices

import (
	"context"
	"golang.org/x/sync/errgroup"
)

func SyncMap[S ~[]V, V any, E any](ctx context.Context, s S, extract func(ctx context.Context, value V) (mapped E, err error)) (mapped []E, err error) {
	mapResults := make([]E, len(s))
	g, ctx := errgroup.WithContext(ctx)
	for idx, item := range s {
		g.Go(func() (err error) {
			mapResults[idx], err = extract(ctx, item)
			return
		})
	}
	if err := g.Wait(); err == nil {
		mapped = mapResults
	}
	return
}
