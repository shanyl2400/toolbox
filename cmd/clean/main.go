package main

import (
	"os"
	"toolbox/internal/config"
	"toolbox/internal/repository"
	"toolbox/internal/repository/boltdb"
)

func main() {
	config.Set(&config.Config{
		BoltDBPath: "./data.db",
	})

	err := boltdb.GetClient().Open()
	if err != nil {
		panic(err)
	}
	defer boltdb.GetClient().Close()

	toolsRepo := repository.GetToolsRepository()
	tools, err := toolsRepo.List()
	if err != nil {
		panic(err)
	}

	profileMap := make(map[string]struct{})
	for _, tool := range tools {
		profileMap[tool.Profile] = struct{}{}
	}

	entries, err := os.ReadDir("./assets")
	if err != nil {
		panic(err)
	}

	deleteFiles := make([]string, 0)
	for _, e := range entries {
		_, ok := profileMap[e.Name()]
		if !ok {
			deleteFiles = append(deleteFiles, "./assets/"+e.Name())
		}
	}

	for _, f := range deleteFiles {
		err := os.Remove(f)
		if err != nil {
			panic(err)
		}
	}
}
