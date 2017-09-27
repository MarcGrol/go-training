package tour

//go:generate gen

// +gen slice:"SortBy,Where"
type Cyclist struct {
	Number int    `json:"number"` // json in lower-case // HL
	Name   string `json:"name"`
	Team   string `json:"team"`
}
