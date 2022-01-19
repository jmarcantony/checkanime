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

	"github.com/fatih/color"
)

type Anime struct {
	Name, Url, WatchedPath, Splitter, EpisodeSplitter                    string
	Episodes, Watched, Missing, EpisodeIndex, AddToWatched, AddToEpisode int
	Markers, marks                                                       []string
	Irregular                                                            bool
}

func (a *Anime) GetWatched() {
	b, err := os.ReadFile(a.WatchedPath)
	if err != nil {
		log.Fatal(err)
	}
	s := string(b)
	n, err := strconv.Atoi(strings.TrimSpace(strings.Split(s, a.EpisodeSplitter)[a.EpisodeIndex]))
	if err != nil {
		log.Fatal(err)
	}
	a.Watched = n + a.AddToWatched
}

func (a *Anime) GetEpisodes(client *http.Client) {
	req, _ := http.NewRequest("GET", a.Url, nil)
	res, _ := client.Do(req)
	body, _ := ioutil.ReadAll(res.Body)
	for _, marker := range a.Markers {
		r, _ := regexp.Compile(marker)
		marks := r.FindAllString(string(body), -1)
		if !a.Irregular {
			marks = removeDupes(marks)
		}
		a.Episodes += len(marks)
		if a.Irregular {
			a.marks = marks
		}
	}
	a.Episodes += a.AddToEpisode
}

func (a *Anime) GetMissing() {
	var t, prev int
	for _, v := range a.marks {
		i, _ := strconv.Atoi(strings.Split(v, a.Splitter)[1])
		if prev+1 != i {
			t++
		}
		prev = i
	}
	a.Missing = t
}

func (a *Anime) ShowReport(client *http.Client) {
	a.GetWatched()
	a.GetEpisodes(client)
	if a.Irregular {
		a.GetMissing()
	}
	cyan := color.New(color.FgCyan).SprintFunc()
	fmt.Printf("\n\tReport for %s:\n\n", cyan(a.Name))
	fmt.Printf("\tEpisodes watched    : %d\n", a.Watched)
	fmt.Printf("\tEpisodes avaialable : %d\n\n", a.Episodes+a.Missing)
	d := (a.Episodes + a.Missing) - a.Watched
	switch {
	case d == 0:
		color.Yellow("\tNo new episodes are added")
	case d > 0:
		color.Green("\t%d new epsiodes are available\n", d)
	default:
		color.Red("\tThats weird, %d episodes were removed\n", d*-1)
	}
	fmt.Printf("\n%s\n", strings.Repeat("=", 50))
}

func removeDupes(s []string) []string {
	var (
		t    int
		l    []string
		seen = map[string]bool{}
	)
	for _, v := range s {
		if _, ok := seen[v]; !ok {
			l = append(l, v)
			seen[v] = true
			t++
		}
	}
	return l[:t]
}
