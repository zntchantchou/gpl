package types

import "time"

type DataType string

// Enum
const (
	String    DataType = "String"
	Integer   DataType = "Integer"
	Date      DataType = "Date"
	Undefined DataType = ""
	// add date type
)

type TrackType struct {
	Title  string
	Artist string
	Album  string
	Year   int
}

func (tr *TrackType) GetStringValue(colName string) (val string) {
	switch colName {
	case "Title":
		val = tr.Title

	case "Artist":
		val = tr.Artist

	case "Album":
		val = tr.Album
	default:
		val = ""
	}
	return
}

func (tr *TrackType) GetIntValue(colName string) (val int) {
	switch colName {
	case "Year":
		val = tr.Year
	default:
		val = 0
	}
	return
}

type Col struct {
	Title       string
	LastClicked time.Time
	DataType    DataType
}
