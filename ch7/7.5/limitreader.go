package nio

import "io"

// Reader struct for reading strings
type Reader struct {
	r io.Reader
	l int64 // limit
	i int64 // current index
}

// LimitReader returns reader with a limit on the amount of bytes which can be read
func LimitReader(or io.Reader, l int64) io.Reader {
	return Reader{or, l, 0}
}

func (r Reader) Read(p []byte) (n int, err error) {
	if r.i >= (r.l) {
		return 0, io.EOF
	}

	// only copy up to limit
	n, err = r.r.Read(p[:r.l])
	if err != nil {
		return 0, err
	}

	r.i += int64(n)

	return n, nil
}
