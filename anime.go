package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Anime struct {
    Name, Url, Marker, WatchedPath, Splitter, FindUrl string
    Episodes, Watched, Missing, EpisodeIndex int
    Client *http.Client
    marks []string
}

func (a *Anime) GetWatched() {
    b, err :=  os.ReadFile(a.WatchedPath)
    if err != nil {
        log.Fatal(err)
    }
    s := string(b)
    n, err := strconv.Atoi(strings.Split(s, " ")[a.EpisodeIndex])
    if err != nil {
        log.Fatal(err)
    }
    a.Watched = n
}

func (a *Anime) GetEpisodes() {
    req, _ := http.NewRequest("GET", a.Url, nil)
    res, _ := a.Client.Do(req)
    body, _ := ioutil.ReadAll(res.Body)
    r, _ := regexp.Compile(a.Marker)
    marks := r.FindAllString(string(body), -1)
    a.Episodes = len(marks)
    a.marks = marks
}

func (a *Anime) GetMissing() {
    var t, prev int
    for _, v := range a.marks {
        i, _ := strconv.Atoi(strings.Split(v, a.Splitter)[1])
        if prev + 1 != i {
            t++
        }
        prev = i
    }
    a.Missing = t
}

func (a *Anime) ShowReport() {
    a.GetWatched()
    a.GetEpisodes()
    a.GetMissing()
    fmt.Printf("\n\tReport for %s:\n\n", a.Name)
    fmt.Printf("\tEpisodes watched    : %d\n", a.Watched)
    fmt.Printf("\tEpisodes avaialable : %d\n\n", a.Episodes + a.Missing)
    d := (a.Episodes + a.Missing) - a.Watched
    switch {
    case d == 0:
        fmt.Println("\tNo new episodes are added")
    case d > 0:
        fmt.Printf("\t%d new epsiodes are available\n", d)
    default:
        fmt.Printf("\tThats weird, %d episodes were removed\n", d*-1)
    }
    fmt.Printf("\n%s\n", strings.Repeat("=", 50))
}
