package graph

import (
	"errors"
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
