package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	path    = "/mnt/c/Users/HP/Desktop/Folders/BorutoEpisode.txt"
	url     = "https://ww4.narutowatchonline.com/tvshows/boruto-subbed-english-online-free/"
	findUrl = "https://ww4.narutowatchonline.com/episodes/boruto-episode-%d-subbed-english-free-online/"
)

var (
	client       = &http.Client{}
	f      *bool = flag.Bool("f", false, "find hidden episodes from snooping url's")
)

func main() {
	flag.Parse()
	watched := getWatched()
	if *f {
		find(&watched)
		return
	}
	fmt.Printf("Episodes watched    : %d\n", watched)
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := client.Do(req)
	body, _ := ioutil.ReadAll(res.Body)
	episodes := getEpisodes(string(body))
	fmt.Printf("Episodes avaialable : %d\n\n", episodes)
	d := episodes - watched
	switch {
	case d == 0:
		fmt.Println("No new episodes are added")
	case d > 0:
		fmt.Printf("%d new epsiodes are available\n", d)
	default:
		fmt.Printf("Thats weird, %d episodes were removed\n", d*-1)
	}
}

func find(watched *int) {
	i := *watched + 1
	for {
		req, _ := http.NewRequest("GET", fmt.Sprintf(findUrl, i), nil)
		res, _ := client.Do(req)
		if res.StatusCode != 200 {
			break
		}
		fmt.Printf("Bruh... episode %d is hiding\n", i)
		i++
	}
	if i == *watched+1 {
		fmt.Println("Looks like no episodes be hinding")
	} else {
		fmt.Printf("\n%d episodes where hiding", i-(*watched+1))
	}
}

func getWatched() int {
	f, _ := os.ReadFile(path)
	i, _ := strconv.Atoi(strings.Split(string(f), " ")[1])
	return i
}

func getEpisodes(b string) int {
	r, _ := regexp.Compile(`mark-\d+`)
	marks := r.FindAllString(b, -1)
	return len(marks) + getMissing(marks)
}

func getMissing(marks []string) int {
	var (
		t    int
		prev int
	)
	for _, v := range marks {
		i, _ := strconv.Atoi(strings.Split(v, "-")[1])
		if prev+1 != i {
			t++
		}
		prev = i
	}
	return t
}
