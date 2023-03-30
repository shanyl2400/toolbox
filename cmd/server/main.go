package main

import (
	"toolbox/internal/api"
	"toolbox/internal/repository/boltdb"
)

func main() {
	err := boltdb.GetClient().Open()
	if err != nil {
		panic(err)
	}
	defer boltdb.GetClient().Close()

	api.Start()
}
