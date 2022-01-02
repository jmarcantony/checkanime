package main

import "net/http"

var (
    client = &http.Client{}
    boruto = &Anime{
        Name: "Boruto",
        WatchedPath: "/mnt/c/Users/HP/Desktop/Folders/BorutoEpisode.txt",
        Url: "https://ww4.narutowatchonline.com/tvshows/boruto-subbed-english-online-free/",
        FindUrl: "https://ww4.narutowatchonline.com/episodes/boruto-episode-%d-subbed-english-free-online/",
        Splitter: "-",
        Marker: `mark-\d+`,
        Client: client,
        EpisodeIndex: 1,
    }
    animeCollection = []*Anime{boruto}
)

func main() {
    for _, anime := range animeCollection  {
        anime.ShowReport()
    }
}
