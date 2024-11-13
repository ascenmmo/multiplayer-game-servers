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
		"RuAuth":           dir + "auth/ru_index.html",
		"RuDocs":           dir + "docs/ru_index.html",
		"RuGameCollection": dir + "game_collection/ru_index.html",
		"RuMainPage":       dir + "main/ru_index.html",
		"RuGameInfo":       dir + "game_info/ru_index.html",
		"RuGameConfig":     dir + "game_config/ru_index.html",
		"RuDeveloperInfo":  dir + "developer_info/ru_index.html",

		"EngAuth":           dir + "auth/eng_index.html",
		"EngDocs":           dir + "docs/eng_index.html",
		"EngGameCollection": dir + "game_collection/eng_index.html",
		"EngMainPage":       dir + "main/eng_index.html",
		"EngGameInfo":       dir + "game_info/eng_index.html",
		"EngGameConfig":     dir + "game_config/eng_index.html",
		"EngDeveloperInfo":  dir + "developer_info/eng_index.html",
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
