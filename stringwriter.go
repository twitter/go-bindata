// This work is subject to the CC0 1.0 Universal (CC0 1.0) Public Domain Dedication
// license. Its contents can be found at:
// http://creativecommons.org/publicdomain/zero/1.0/

package bindata

import "io"

const lowerHex = "0123456789abcdef"
const hexDigitStart = 2
const hexDigitEnd = 3

type StringWriter struct {
	io.Writer
	c int
}

func (w *StringWriter) Write(p []byte) (n int, err error) {
	if len(p) == 0 {
		return
	}

	buf := []byte{'\\', 'x', '0', '0'}

	for n, b := range p {
		buf[hexDigitStart] = lowerHex[b/16]
		buf[hexDigitEnd] = lowerHex[b%16]
		w.Writer.Write(buf)
		w.c++
	}

	n++

	return
}
