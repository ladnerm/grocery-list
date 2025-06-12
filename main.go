package main

import (
	"github.com/ladnerm/grocery-list/api"
	"github.com/ladnerm/grocery-list/env"
)

// TODO:
// total cost of all products = sum of each estimated price
// postgres

func main() {
	env.SetEnv()
	api.StartServer()
}
