package main

import (
	"sort"
	"testing"
)

func compare(a []Album, b []Album) bool {
	if len(a) != len(b) {
		return false
	}
	for i, elt := range a {
		if elt != b[i] {
			return false
		}
	}
	return true
}

func TestSortByYear(t *testing.T) {
	albums := []Album{
		{Artist: "MFDoom", Title: "Born Like This", Year: 2009},
		{Artist: "MFDoom", Title: "MM..Food", Year: 2004},
		{Artist: "Teebs", Title: "Ardour", Year: 2017},
		{Artist: "Zartist", Title: "A Title", Year: 2002},
		{Artist: "Alicia Keys", Title: "Diary of Alicia Keys", Year: 2002},
	}
	sorted := []Album{
		{Artist: "Teebs", Title: "Ardour", Year: 2017},
		{Artist: "MFDoom", Title: "Born Like This", Year: 2009},
		{Artist: "MFDoom", Title: "MM..Food", Year: 2004},
		{Artist: "Zartist", Title: "A Title", Year: 2002},
		{Artist: "Alicia Keys", Title: "Diary of Alicia Keys", Year: 2002},
	}
	sortable := &ByColumnSort{albums, nil}
	sortable.AddFilter(sortable.LessYear)
	sort.Sort(sortable)
	// PrintAlbums(sortable.A)
	if !compare(sortable.A, sorted) {
		t.Fail()
	}
}

func TestSortByTitle(t *testing.T) {
	albums := []Album{
		{Artist: "MFDoom", Title: "Born Like This", Year: 2009},
		{Artist: "MFDoom", Title: "MM..Food", Year: 2004},
		{Artist: "Teebs", Title: "Ardour", Year: 2017},
		{Artist: "Someguy", Title: "Bottom", Year: 2002},
		{Artist: "Alicia Keys", Title: "Diary of Alicia Keys", Year: 2002},
	}
	sorted := []Album{
		{Artist: "Teebs", Title: "Ardour", Year: 2017},
		{Artist: "MFDoom", Title: "Born Like This", Year: 2009},
		{Artist: "Someguy", Title: "Bottom", Year: 2002},
		{Artist: "Alicia Keys", Title: "Diary of Alicia Keys", Year: 2002},
		{Artist: "MFDoom", Title: "MM..Food", Year: 2004},
	}
	sortable := &ByColumnSort{albums, nil}
	sortable.AddFilter(sortable.LessTitle)
	sort.Sort(sortable)
	// PrintAlbums(sortable.A)
	if !compare(sortable.A, sorted) {
		t.Fail()
	}
}
func TestSortByYearAndTitle(t *testing.T) {
	albums := []Album{
		{Artist: "MFDoom", Title: "Born Like This", Year: 2009},
		{Artist: "MFDoom", Title: "MM..Food", Year: 2002},
		{Artist: "Someguy", Title: "Bottom", Year: 2002},
		{Artist: "Teebs", Title: "Ardour", Year: 2017},
		{Artist: "Alicia Keys", Title: "Diary of Alicia Keys", Year: 2002},
	}
	sorted := []Album{
		{Artist: "Teebs", Title: "Ardour", Year: 2017},
		{Artist: "MFDoom", Title: "Born Like This", Year: 2009},
		{Artist: "Someguy", Title: "Bottom", Year: 2002},
		{Artist: "Alicia Keys", Title: "Diary of Alicia Keys", Year: 2002},
		{Artist: "MFDoom", Title: "MM..Food", Year: 2002},
	}
	sortable := &ByColumnSort{albums, nil}
	sortable.AddFilter(sortable.LessTitle)
	sortable.AddFilter(sortable.LessYear)
	sort.Sort(sortable)
	// PrintAlbums(sortable.A)
	if !compare(sortable.A, sorted) {
		t.Fail()
	}
}

func TestSortByYearAndTitleAndArtist(t *testing.T) {
	albums := []Album{
		{Artist: "Teebs", Title: "Ardour", Year: 2002},
		{Artist: "MFDoom", Title: "Ardour", Year: 2002},
		{Artist: "Someguy", Title: "Ardour", Year: 2002},
		{Artist: "Someguy", Title: "Ardour", Year: 2019},
	}

	sorted := []Album{
		{Artist: "Someguy", Title: "Ardour", Year: 2019},
		{Artist: "MFDoom", Title: "Ardour", Year: 2002},
		{Artist: "Someguy", Title: "Ardour", Year: 2002},
		{Artist: "Teebs", Title: "Ardour", Year: 2002},
	}
	sortable := &ByColumnSort{albums, nil}
	sortable.AddFilter(sortable.LessArtist)
	sortable.AddFilter(sortable.LessTitle)
	sortable.AddFilter(sortable.LessYear)
	sort.Sort(sortable)
	// PrintAlbums(sortable.A)
	if !compare(sortable.A, sorted) {
		t.Fail()
	}
}
