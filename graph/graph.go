package graph

import (
	"errors"
	"fmt"
	"sync"
)

type vertex[T comparable] struct {
	id    T
	value any
}

type edge[T comparable] struct {
	from   T
	to     T
	weight float64
}

type Graph[T comparable] struct {
	vertices map[T]*vertex[T]
	edges    map[T][]edge[T]
	directed bool
	weighted bool
	lock     sync.RWMutex
}

func New[T comparable](directed bool, weighted bool) *Graph[T] {
	return &Graph[T]{
		vertices: make(map[T]*vertex[T]),
		edges:    make(map[T][]edge[T]),
		directed: directed,
		weighted: weighted,
	}
}

func (g *Graph[T]) AddVertex(id T, value any) {
	g.lock.Lock()
	defer g.lock.Unlock()

	if _, exists := g.vertices[id]; exists {
		return
	}
	g.vertices[id] = &vertex[T]{id: id, value: value}
}

func (g *Graph[T]) RemoveVertex(id T) error {
	g.lock.Lock()
	defer g.lock.Unlock()

	if _, exists := g.vertices[id]; !exists {
		return errors.New("vertex does not exists")
	}

	delete(g.vertices, id)
	delete(g.edges, id)

	for from, edges := range g.edges {
		newEdges := []edge[T]{}
		for _, edge := range edges {
			if edge.to != id {
				newEdges = append(newEdges, edge)
			}
		}
		g.edges[from] = newEdges
	}
	return nil
}

func (g *Graph[T]) AddEdge(from, to T, weight float64) error {
	g.lock.Lock()
	defer g.lock.Unlock()

	if _, exists := g.vertices[from]; !exists {
		return fmt.Errorf("vertex %v does not exists", from)
	}
	if _, exists := g.vertices[to]; !exists {
		return fmt.Errorf("vertex %v does not exists", to)
	}

	if !g.weighted {
		weight = 1.0
	}

	e := edge[T]{from: from, to: to, weight: weight}
	g.edges[from] = append(g.edges[from], e)
	if !g.directed {
		reverseEdge := edge[T]{from: to, to: from, weight: weight}
		g.edges[to] = append(g.edges[to], reverseEdge)
	}
	return nil
}
