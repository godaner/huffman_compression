package huffman_compression

import (
	"io"
)

type HuffmanCompression struct {
}

func (hc *HuffmanCompression) Decode(r io.Reader, w io.Writer) (err error) {
	return (&decoder{
		r: r,
		w: w,
	}).decode()
}
func (hc *HuffmanCompression) Encode(r io.Reader, w io.Writer) (err error) {
	return (&encoder{
		r: r,
		w: w,
	}).encode()
}

type huffmanTreeHeader struct {
	eLen uint16
	elem []huffmanTreeHeaderElem
}
type huffmanTreeHeaderElem struct {
	k     byte
	vsLen byte
	vs    []byte
}
