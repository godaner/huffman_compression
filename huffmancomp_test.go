package huffman_compression

import (
	"os"
	"testing"
)

func TestHuffmanCompression_Encode(t *testing.T) {
	src := "/home/godaner/config.json"
	srcFile, err := os.Open(src)
	if err != nil {
		panic(err)
	}
	dst := "/home/godaner/config.hfm"
	dstFile, err := os.Create(dst)
	if err != nil {
		panic(err)
	}
	hc := HuffmanCompression{}
	err = hc.Encode(srcFile, dstFile)
	if err != nil {
		panic(err)
	}
}
