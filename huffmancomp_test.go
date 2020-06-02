package huffman_compression

import (
	"fmt"
	"github.com/icza/bitio"
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
	f,err:=os.Create("/home/godaner/Downloads/template.tt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	bw:=bitio.NewWriter(f)
	ds:=[]byte{}
	for i:=0;i<10;i++{
		ds=append(ds,[]byte{1,0,1,1,1,0,1,1}...)
	}
	n, err := bw.Write(ds)
	if err != nil {
		panic(err)
	}
	defer bw.Close()
	fmt.Println(n)
}