package nstrings

import "io"

// Reader struct for reading strings
type Reader struct {
	s string
	i int64 // current index
}

// NewReader creates Reader
func NewReader(s string) *Reader {
	return &Reader{s, 0}
}

func (r *Reader) Read(p []byte) (n int, err error) {
	if r.i >= int64(len(r.s)-1) {
		return 0, io.EOF
	}
	n = copy(p, []byte(r.s[r.i:]))
	r.i += int64(n)
	return n, nil
}
