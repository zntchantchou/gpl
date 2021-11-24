package main

import (
	"fmt"
	"sort"
)

type Album struct {
	Title  string
	Artist string
	Year   int
}

type comparison int

type Filter func(a, b *Album) comparison

const (
	lt comparison = iota
	eq
	gt
)

type ByColumnSort struct {
	A       []Album
	Filters []Filter
}

func (bs *ByColumnSort) Len() int {
	return len(bs.A)
}

func (bs *ByColumnSort) Swap(i int, j int) {
	bs.A[i], bs.A[j] = bs.A[j], bs.A[i]
}

func (bs *ByColumnSort) Less(i int, j int) bool {
	for _, f := range bs.Filters {
		cmp := f(&bs.A[i], &bs.A[j])
		switch cmp {
		case eq:
			continue
		case lt:
			return true
		case gt:
			return false
		}
	}
	return false
}

func (bs *ByColumnSort) LessArtist(a, b *Album) comparison {
	switch {
	case a.Artist == b.Artist:
		return eq
	case a.Artist < b.Artist:
		return lt
	default:
		return gt
	}
}

func (bs *ByColumnSort) LessTitle(a, b *Album) comparison {
	switch {
	case a.Title == b.Title:
		return eq
	case a.Title < b.Title:
		return lt
	default:
		return gt
	}
}

func (bs *ByColumnSort) LessYear(a, b *Album) comparison {
	switch {
	case a.Year == b.Year:
		return eq
	case a.Year > b.Year:
		return lt
	default:
		return gt
	}
}

func (bs *ByColumnSort) AddFilter(f Filter) {
	bs.Filters = append([]Filter{f}, bs.Filters...)
}

func PrintAlbums(albs []Album) {
	for _, alb := range albs {
		fmt.Println("Artist: ", alb.Artist, "Title: ", alb.Title, "Year: ", alb.Year)
	}
}

func main() {

	albums := []Album{
		{Artist: "MFDoom", Title: "Born Like This", Year: 2009},
		{Artist: "MFDoom", Title: "MM..Food", Year: 2004},
		{Artist: "Teebs", Title: "Ardour", Year: 2017},
		{Artist: "Jean-Marie Bigard", Title: "Born Like This", Year: 2019},
		{Artist: "Zozo le chapo", Title: "Born Like This", Year: 2021},
		{Artist: "Zartist", Title: "A Title", Year: 2002},
		{Artist: "Alicia Keys", Title: "Diary of Alicia Keys", Year: 2002},
	}
	sortable := &ByColumnSort{albums, nil}
	sortable.AddFilter(sortable.LessYear)
	sortable.AddFilter(sortable.LessTitle)

	// last added filter get priority
	fmt.Println(len(sortable.Filters))
	// sortable.AddFilter(sortable.LessTitle)
	sort.Sort(sortable)
	for _, a := range sortable.A {
		fmt.Printf("Artist:  %v \t Title:  %v \t Year:  %v \t \n", a.Artist, a.Title, a.Year)
	}
}
