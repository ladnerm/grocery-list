package types

type Item struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	User     string `json:"user"`
	Location string `json:"location"`
}

func NewItem(name string, user string, location string) *Item {
	return &Item{
		ID:       0,
		Name:     name,
		User:     user,
		Location: location,
	}
}
