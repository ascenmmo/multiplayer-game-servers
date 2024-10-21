package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

const dir = "pkg/admin_client/assets/pages/"

var (
	indexFiles = map[string]string{
		"Auth":           dir + "auth/index.html",
		"Docs":           dir + "docs/ru_index.html",
		"GameCollection": dir + "game_collection/index.html",
		"MainPage":       dir + "main/index.html",
		"GameInfo":       dir + "game_info/index.html",
		"GameConfig":     dir + "game_config/index.html",
	}
)

func main() {
	pagesFile := "package pages\n\n"

	keys := make([]string, 0, len(indexFiles))
	for key := range indexFiles {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, name := range keys {
		filepath := indexFiles[name]
		file, err := os.ReadFile(filepath)
		if err != nil {
			panic(err)
		}
		marshal, err := json.Marshal(string(file))
		if err != nil {
			panic(err)
		}
		dataConstant := fmt.Sprintf(`const %s = %s`+"\n\n", name, string(marshal))
		pagesFile += dataConstant
	}

	create, err := os.Create(dir + "pages.go")
	if err != nil {
		panic(err)
	}
	defer create.Close()

	_, err = create.WriteString(pagesFile)
	if err != nil {
		panic(err)
	}
}
