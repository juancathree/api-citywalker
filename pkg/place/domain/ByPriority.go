package domain

type ByPriority []Place

func (b ByPriority) Len() int {
	return len(b)
}

func (b ByPriority) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b ByPriority) Less(i, j int) bool {
	return b[i].Priority < b[j].Priority
}
