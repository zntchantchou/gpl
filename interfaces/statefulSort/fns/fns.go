package fns

import (
	"time"

	"github.com/zclmk/statefulsort/types"
)

func SortDates(a time.Time, b time.Time) bool {
	if !a.Equal(b) {
		return a.Before(b)
	}
	return true
}

func SortStrings(a string, b string) bool {
	if a != b {
		return a < b
	}
	return true
}

func SortInts(a int, b int) bool {
	if a != b {
		return a < b
	}
	return true
}

func SortTracks(a *types.TrackType, b *types.TrackType, t types.DataType, colName string) bool {
	switch t {
	case "String":
		return SortStrings(a.GetStringValue(colName), b.GetStringValue(colName))
	case "Integer":
		return SortInts(a.GetIntValue(colName), b.GetIntValue(colName))
	default:
		return false
	}
}
