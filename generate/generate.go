package tour

//go:generate gen

// +gen slice:"SortBy,Where"
type Cyclist struct {
	Number int
	Name   string
	Team   string
}
