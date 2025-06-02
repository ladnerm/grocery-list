package util

import (
	"encoding/json"
	"github.com/ladnerm/grocery-list/types"
	"os"
)

func ItemsFromDB(itemArr *[]types.Item) error {
	db, err := os.OpenFile("db/db.json", os.O_RDWR, 0644)
	defer db.Close()

	if err != nil {
		return err
	}

	err = json.NewDecoder(db).Decode(itemArr)
	if err != nil {
		return err
	}

	return nil
}

func WriteToDB(itemArr []types.Item) error {
	db, err := os.Create("db/db.json")
	defer db.Close()
	if err != nil {
		return err
	}

	err = json.NewEncoder(db).Encode(itemArr)
	if err != nil {
		return err
	}

	return nil
}
