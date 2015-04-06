package tour

//go:generate gen

// +gen slice:"SortBy,Where"
type Cyclist struct {
	Number int    `json:"number"`
	Name   string `json:"name"`
	Team   string `json:"team"`
}
