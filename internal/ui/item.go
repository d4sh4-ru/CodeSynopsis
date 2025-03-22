package ui

type item struct {
	Name  string
	Path  string
	IsDir bool
}

func (i item) Title() string       { return i.Name }
func (i item) Description() string { return "" }
func (i item) FilterValue() string { return i.Name }
