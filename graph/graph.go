package graph

import (
	"fmt"
	"strings"
)

type Graph interface {
	AddVertex(key int) (*vertex, error)
	GetVertex(key int) (*vertex, error)
	DeleteVertex(key int) error
	AddEdge(from, to int) error
	RemoveEdge(from, to int) error
	FindPath(start, end int)
}

type graphImpl struct {
	vertices []*vertex
	size     int
}

func NewGraph() Graph {
	return &graphImpl{
		vertices: make([]*vertex, 0),
		size:     0,
	}
}

func (g *graphImpl) AddVertex(key int) (*vertex, error) {
	if contains(g.vertices, key) {
		return nil, fmt.Errorf("vertex with that key (%d) already exists", key)
	}

	g.size++
	vertex := newVertex(key)
	g.vertices = append(g.vertices, vertex)

	return vertex, nil
}

func (g *graphImpl) GetVertex(key int) (*vertex, error) {
	for _, v := range g.vertices {
		if v.key == key {
			return v, nil
		}
	}
	return nil, fmt.Errorf("vertex with given key (%d) not found", key)
}

func (g *graphImpl) DeleteVertex(key int) error {
	idx := -1
	for i, v := range g.vertices {
		if v.key == key {
			idx = i
		}
	}
	if idx == -1 {
		return fmt.Errorf("vertex with that key (%d) not found", key)
	}

	g.size--
	g.vertices = append(g.vertices[:idx], g.vertices[idx+1:]...)

	return nil
}

func (g *graphImpl) GetOrAddVertex(key int) *vertex {
	v, err := g.GetVertex(key)
	if err == nil {
		return v
	}

	newVertex, _ := g.AddVertex(key)
	return newVertex
}

func (g *graphImpl) AddEdge(from, to int) error {
	fromV, err := g.GetVertex(from)
	if err != nil {
		return err
	}

	toV, err := g.GetVertex(to)
	if err != nil {
		return err
	}

	if contains(fromV.adjVertices, toV.key) {
		return fmt.Errorf("edge between %d -> %d already exists", fromV.key, toV.key)
	}

	fromV.addAdjacentVertex(toV)
	return nil
}

func (g *graphImpl) RemoveEdge(from, to int) error {
	fromV, err := g.GetVertex(from)
	if err != nil {
		return err
	}

	toV, err := g.GetVertex(to)
	if err != nil {
		return err
	}

	fromV.removeAdjacentVertex(toV)
	return nil
}

func (g *graphImpl) FindPath(start, end int) {

}

func contains(list []*vertex, key int) bool {
	for _, v := range list {
		if v.key == key {
			return true
		}
	}
	return false
}

func (g *graphImpl) String() string {
	var buff strings.Builder
	buff.WriteString("\n##### GRAPH #####\n")

	for i, v := range g.vertices {
		buff.WriteString(fmt.Sprintf("%d.[%d] - ", i, v.key))
		for _, adjV := range v.adjVertices {
			buff.WriteString(fmt.Sprintf("[%d]->", adjV.key))
		}
		buff.WriteString("\n")
	}

	buff.WriteString("##### END #####\n")
	return buff.String()
}
