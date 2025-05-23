package types

type Item struct {
	Name     string `json:"name"`
	User     string `json:"user"`
	Location string `json:"location"`
}

func NewItem(name string, user string, location string) *Item {
	return &Item{
		Name:     name,
		User:     user,
		Location: location,
	}
}
