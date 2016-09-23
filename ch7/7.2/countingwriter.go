package countingwriter

import "io"

// CountingWriter returns writer with cummalitive counter
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	cw := countingWriter{w, 0}
	return cw, &cw.Count
}

type countingWriter struct {
	w     io.Writer
	Count int64
}

func (cw countingWriter) Write(p []byte) (int, error) {
	c, err := cw.Write(p)
	cw.Count += int64(c)
	return c, err
}
