package element

type Els struct {
	Nodes []uint32
}

func NewEls() Els {
	return Els{Nodes: make([]uint32, 0)}
}

func (els *Els) AddElement(e Element) {
	els.Nodes = append(els.Nodes, e.ToMat())
}
