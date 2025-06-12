package types

type Item struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	User     string `json:"user"`
	Location string `json:"location"`
}

func NewItem(id int, name string, user string, location string) *Item {
	return &Item{
		ID:       id,
		Name:     name,
		User:     user,
		Location: location,
	}
}

type Config struct {
	Host    string
	Port    string
	User    string
	DBName  string
	SSLMODE string
}
