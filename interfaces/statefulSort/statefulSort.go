package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"

	"github.com/zclmk/statefulsort/fns"
	"github.com/zclmk/statefulsort/types"
)

type byLastClicked []*types.Col

type lastClickedSort struct {
	tracks []*types.TrackType
	less   func(x, y *types.TrackType) bool
}

func (lcs lastClickedSort) Len() int {
	return len(lcs.tracks)
}

func (lcs lastClickedSort) Less(i int, j int) bool {
	return less(lcs.tracks[i], lcs.tracks[j])
}

func (lcs lastClickedSort) Swap(i int, j int) {
	lcs.tracks[i], lcs.tracks[j] = lcs.tracks[j], lcs.tracks[i]
}

func less(x, y *types.TrackType) bool {
	for _, col := range cols {
		// skip if values are identical
		if types.DataType(col.DataType) == "String" {
			if x.GetStringValue(col.Title) != y.GetStringValue(col.Title) {
				return fns.SortTracks(x, y, types.DataType(col.DataType), col.Title)
			}
		}
		if types.DataType(col.DataType) == "Integer" {
			if x.GetIntValue(col.Title) != y.GetIntValue(col.Title) {
				return fns.SortTracks(x, y, types.DataType(col.DataType), col.Title)
			}
		}
	}
	return false
}

var tracks = []*types.TrackType{
	{Title: "Go", Artist: "Delilah", Album: "From the Roots Up", Year: 2012},
	{Title: "Go", Artist: "Moby", Album: "Moby", Year: 1992},
	{Title: "Beat it", Artist: "Michael Jackson", Album: "Thriller", Year: 2007},
	{Title: "Amazin", Artist: "Michael Jackson", Album: "Lolilol", Year: 2009},
	{Title: "Go Ahead", Artist: "Alicia Keys", Album: "ZAs I Am", Year: 2007},
	{Title: "Ready 2 Go", Artist: "Martin Solveig", Album: "Smash", Year: 2011},
}

var cols = []*types.Col{
	{Title: "Title", LastClicked: time.Date(2021, time.Month(2), 21, 5, 5, 0, 0, time.UTC), DataType: "String"},
	{Title: "Artist", LastClicked: time.Date(2021, time.Month(2), 21, 15, 10, 0, 0, time.UTC), DataType: "String"},
	{Title: "Album", LastClicked: time.Date(2021, time.Month(2), 21, 20, 15, 0, 0, time.UTC), DataType: "String"},
	{Title: "Year", LastClicked: time.Date(2021, time.Month(2), 21, 10, 20, 0, 0, time.UTC), DataType: "Integer"},
}

// ---------------------------------------- Methods ----------------------------------------

func (bd byLastClicked) Len() int {
	return len(bd)
}

func (bd byLastClicked) Swap(i int, j int) {
	bd[i], bd[j] = bd[j], bd[i]
}

func (bd byLastClicked) Less(i int, j int) bool {
	return bd[j].LastClicked.Before(bd[i].LastClicked)
}

func printTable(tracks []*types.TrackType) {
	const format = "%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year")
	fmt.Fprintf(tw, format, "-------", "-------", "-------", "-------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year)
	}
	tw.Flush()
}

func main() {
	// sort columns by last clicked
	sort.Sort(byLastClicked(cols))
	sort.Sort(lastClickedSort{tracks: tracks, less: less})
	printTable(tracks)
}
