package main

import (
	"net/http"
	"sync"
)

var (
	client          = &http.Client{}
	animeCollection = []*Anime{
		{
			Name:            "Boruto",
			WatchedPath:     "/mnt/c/Users/HP/Desktop/Folders/BorutoEpisode.txt",
			Url:             "https://ww4.narutowatchonline.com/tvshows/boruto-subbed-english-online-free/",
			FindUrl:         "https://ww4.narutowatchonline.com/episodes/boruto-episode-%d-subbed-english-free-online/",
			Splitter:        "-",
			Markers:         []string{`mark-\d+`},
			Client:          client,
			EpisodeSplitter: " ",
			EpisodeIndex:    1,
			Irregular:       true,
		},
		{
			Name:            "Demon Slayer",
			WatchedPath:     "/mnt/c/Users/HP/Desktop/Folders/demonslayerepisodes.txt",
			Url:             "https://ww2.demonslayerepisodes.com/demon-slayer-english-subbed/",
			Markers:         []string{`Demon Slayer Episode \d+ English Subbed`, `Demon Slayer Season \d+ Episode \d+ English Subbed`},
			Splitter:        " ",
			Client:          client,
			EpisodeSplitter: "\n",
			EpisodeIndex:    0,
			AddToEpisode:    -1,
			AddToWatched:    26,
		},
                {
                    Name: "Jujutsu Kaisen",
                    WatchedPath: "/mnt/c/Users/HP/Desktop/Folders/JujutsuKaisenEpisodes.txt",
                    Url: "https://watchjujutsukaisen4freeonline.blogspot.com/p/english-subbed_30.html",
                    Markers: []string{`Jujutsu Kaisen - Episode \d+`},
                    Client: client,
                    EpisodeSplitter: " ",
                    EpisodeIndex: 1,
                },
	}
	wg sync.WaitGroup
)

func main() {
	for _, anime := range animeCollection {
		wg.Add(1)
		go func(a *Anime) {
			a.ShowReport()
			wg.Done()
		}(anime)
	}
	wg.Wait()
}
