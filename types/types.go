package types

var id = 0

type Item struct {
    ID       int    `json:"id"`
	Name     string `json:"name"`
	User     string `json:"user"`
	Location string `json:"location"`

}

func NewItem(name string, user string, location string) *Item {
    id++
	return &Item{
        ID:       id, 
		Name:     name,
		User:     user,
		Location: location,
	}
}
