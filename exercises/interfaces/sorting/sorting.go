package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

// a sort interface that does not assume anything about the sequence to sort (data structure) or its elements

// needs
/*
- length
- how to compare
- how to swap
*/

var ints = []int{10, 7, 12, 9, 10}

func main() {

	intSlice := sort.IntSlice(ints)

	// why sort the reverse to get the reverse ?
	// because sort.Reverse only returns reverses the indexes used by the sort function of our data structure
	sort.Sort(sort.Reverse(intSlice))
	fmt.Println(intSlice)
	fmt.Println("before --------------------------")
	// printTracks(tracks)
	// mutates tracks
	// sort.Sort(byArtist(tracks))
	sort.Sort(customSort{tracks, func(x, y *TrackType) bool {
		if x.Title != y.Title {
			return x.Title < y.Title
		}

		if x.Year != y.Year {
			return x.Year < y.Year
		}

		if x.Length != y.Length {
			return x.Length < y.Length
		}
		return false
	}})
	printTracks(tracks)
}

// one type per element type and per property used
type byArtist []*TrackType

func (tracks byArtist) Len() int {
	return len(tracks)
}

func (tracks byArtist) Less(i int, j int) bool {
	return tracks[i].Artist < tracks[j].Artist
}

func (tracks byArtist) Swap(i int, j int) {
	tracks[i], tracks[j] = tracks[j], tracks[i]
}

type TrackType struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*TrackType{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(d string) time.Duration {
	p, err := time.ParseDuration(d)
	if err != nil {
		panic(err)
	}
	return p
}

func printTracks(tracks []*TrackType) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "------", "------", "------", "------", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush()
}

// A more flexible type

type customSort struct {
	t    []*TrackType
	less func(x, y *TrackType) bool
}

func (cs customSort) Less(i int, j int) bool {
	return cs.less(cs.t[i], cs.t[j])
}

func (cs customSort) Swap(i int, j int) {
	cs.t[i], cs.t[j] = cs.t[j], cs.t[i]
}

func (cs customSort) Len() int {
	return len(cs.t)
}
