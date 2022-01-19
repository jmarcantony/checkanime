# checkanime
Check if new episodes of you're favourite anime has been released from you're terminal

# Installation
	go install github.com/jmarcantony/checkanime@latest

# Configuiration
You will need to create a json file to store all data needed to scrape data from th website having all episodes of you're anime, where all anime objects are in an array and each object can have the following keys:

`name`: name of anime
<br>
`url`: url of website which lists all episodes of this anime
<br>
`watchedPath`: path to where you store how many episodes of the anime you've watched
<br>
`Splitter`: character which splits marker from episode number from other text
<br>
`episodeSplitter`: character which splits episode number from other text in you're episodes watched file
<br>
`episodeIndex`: index of episode number after being split by `episodeSplitter` (start from 0)
<br>
`addToWatched`: number you might want to add to the number of episodes you've watched during runtime
<br>
`addToEpisode`: number you might want to add to episodes scraped from the website
<br>
`markers`: array of strings where each string contains regex text which every episode has in common of in the websites html
<br>
`irregular`: boolean, set true if multiple occurunces of the same episode exists in the website

here's an example:
```
[
    {
    	"name": "myfavanime",
    	"watchedPath": "/path/to/watchedepisodes.txt",
    	"url": "https://example.com/",
    	"markers": ["episode \\d+", "season \\d+ episode \\d+"],
    	"splitter": " ",
    	"episodeSplitter": " ",
    	"episodeIndex": 1,
    	"addToEpisode": -1,
    	"addToWatched": 26
		"irregular": true
	}
]
```
in my case `/path/to/watchedepisodes.txt` looks like
```
Episode: 7
```
# Screenshot
![screenshot](https://raw.githubusercontent.com/jmarcantony/checkanime/main/images/screenshot.png)