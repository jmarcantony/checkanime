package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
)

var (
	client          = &http.Client{}
	jsonPath        = flag.String("p", "", "path to json file")
	animeCollection []*Anime
	wg              sync.WaitGroup
)

func loadJson(path string) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(b, &animeCollection); err != nil {
		log.Fatal(err)
	}
}

func main() {
	flag.Parse()
	if *jsonPath == "" {
		if path, exists := os.LookupEnv("ANIME_JSON"); !exists {
			log.Fatal("No path provided to json file")
		} else {
			loadJson(path)
		}
	} else {
		loadJson(*jsonPath)
	}

	for _, anime := range animeCollection {
		wg.Add(1)
		go func(a *Anime) {
			a.ShowReport(client)
			wg.Done()
		}(anime)
	}
	wg.Wait()
}
