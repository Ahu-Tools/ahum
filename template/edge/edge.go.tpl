package edge

import (
	"context"
	"sync"
	// @ahum: imports
)

type Edge interface {
	Configure()
	Start(context.Context, *sync.WaitGroup)
}

func Start(ctx context.Context, wg *sync.WaitGroup) {
	edges := []Edge{
		// @ahum: edges
	}

	wg.Add(len(edges))
	for _, edge := range edges {
		edge.Configure()
		go edge.Start(ctx, wg)
	}

}
