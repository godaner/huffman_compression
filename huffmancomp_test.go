package huffman_compression

import (
	"os"
	"testing"
)

func TestHuffmanCompression_Encode(t *testing.T) {
	//src := "/home/godaner/Downloads/youkuclient_setup_ywebtop1_7.7.6.4031.exe"
	//src := "/home/godaner/Downloads/template.json"
	src := "/home/godaner/Downloads/a.jpg"
	srcFile, err := os.Open(src)
	if err != nil {
		panic(err)
	}
	defer srcFile.Close()
	//dst := "/home/godaner/Downloads/youkuclient_setup_ywebtop1_7.7.6.4031.hfm"
	//dst := "/home/godaner/Downloads/template.hfm"
	dst := "/home/godaner/Downloads/a.hfm"
	dstFile, err := os.Create(dst)
	if err != nil {
		panic(err)
	}
	defer dstFile.Close()
	hc := HuffmanCompression{}
	err = hc.Encode(srcFile, dstFile)
	if err != nil {
		panic(err)
	}
}

func TestHuffmanCompression_Decode(t *testing.T) {
	//src := "/home/godaner/Downloads/youkuclient_setup_ywebtop1_7.7.6.4031.exe"
	//src := "/home/godaner/Downloads/template.json"
	src := "/home/godaner/Downloads/a.hfm"
	srcFile, err := os.Open(src)
	if err != nil {
		panic(err)
	}
	defer srcFile.Close()
	//dst := "/home/godaner/Downloads/youkuclient_setup_ywebtop1_7.7.6.4031.hfm"
	//dst := "/home/godaner/Downloads/template.hfm"
	dst := "/home/godaner/Downloads/a1.jpg"
	dstFile, err := os.Create(dst)
	if err != nil {
		panic(err)
	}
	defer dstFile.Close()
	hc := HuffmanCompression{}
	err = hc.Decode(srcFile, dstFile)
	if err != nil {
		panic(err)
	}
}