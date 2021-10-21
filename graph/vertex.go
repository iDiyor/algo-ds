package graph

type vertex struct {
	key         int
	adjVertices []*vertex
}

func newVertex(key int) *vertex {
	return &vertex{
		key:         key,
		adjVertices: make([]*vertex, 0),
	}
}

func (v *vertex) addAdjacentVertex(adjV *vertex) {
	v.adjVertices = append(v.adjVertices, adjV)
}

func (v *vertex) removeAdjacentVertex(adjV *vertex) {
	idx := -1
	for i, v := range v.adjVertices {
		if v.key == adjV.key {
			idx = i
		}
	}

	if idx == -1 {
		return
	}

	v.adjVertices = append(v.adjVertices[:idx], v.adjVertices[idx+1:]...)
}
