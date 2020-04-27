package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())
	for i := 0; i < 10; i++ {
		g.Go(func() error {
			return run(ctx, time.Second)
		})
	}

	if err := g.Wait(); err != nil {
		fmt.Println(err)
	}
}

func run(ctx context.Context, d time.Duration) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(d):
		return nil
	}
}
