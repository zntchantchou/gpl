package countingwriter

import (
	"io"

	"github.com/zclmk/interfaces/counters"
)

var TotalBytes = int64(0)

func Write(p []byte) (int, error) {
	return len(p), nil
}

func Countingwriter(w io.Writer) (io.Writer, *int64) {
	c := &counters.ByteCounter{W: w, Written: 0}
	return c, &c.Written
}
