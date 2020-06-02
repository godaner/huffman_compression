package huffman_compression

import (
	"bytes"
	"encoding/binary"
	"errors"
	"github.com/godaner/func/tree"
	huffmantree "github.com/godaner/func/tree/huffman"
	"github.com/icza/bitio"
	"io"
	"io/ioutil"
	"strconv"
)

type HuffmanCompression struct {
	huffmanTree tree.HuffmanTree
	codes       map[int]string
	r           io.Reader
	w           io.Writer
}

func (hc *HuffmanCompression) Decode(r io.Reader, w io.Writer) (err error) {
	return nil
}
func (hc *HuffmanCompression) Encode(r io.Reader, w io.Writer) (err error) {
	hc.r = r
	hc.w = w
	datas, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	err = hc.buildHuffmanTree(datas)
	if err != nil {
		return err
	}
	err = hc.writeHuffmanHeader()
	if err != nil {
		return err
	}
	err = hc.writeHuffmanDatas(datas)
	if err != nil {
		return err
	}
	return nil
}

type huffmanTreeHeader struct {
	eLen byte
	elem []huffmanTreeHeaderElem
}
type huffmanTreeHeaderElem struct {
	k     byte
	vsLen byte
	vs    []byte
}

// writeHuffmanHeader
func (hc *HuffmanCompression) writeHuffmanHeader() (err error) {
	header := hc.buildHuffmanTreeHeader()
	buf := new(bytes.Buffer)
	err = binary.Write(buf, binary.BigEndian, header.eLen)
	if err != nil {
		return err
	}
	for _, elem := range header.elem {
		err = binary.Write(buf, binary.BigEndian, elem.k)
		if err != nil {
			return err
		}
		err = binary.Write(buf, binary.BigEndian, elem.vsLen)
		if err != nil {
			return err
		}
		err = binary.Write(buf, binary.BigEndian, elem.vs)
		if err != nil {
			return err
		}
	}
	err = binary.Write(hc.w, binary.BigEndian, buf.Bytes())
	if err != nil {
		return err
	}
	return nil
}

func (hc *HuffmanCompression) buildHuffmanTreeHeader() huffmanTreeHeader {
	elems := make([]huffmanTreeHeaderElem, 0)
	for signal, code := range hc.codes {
		vs := []byte{}
		for _, c := range code {
			vs = append(vs, byte(c))
		}
		elems = append(elems, huffmanTreeHeaderElem{
			k:     byte(signal),
			vsLen: byte(len(vs)),
			vs:    vs,
		})
	}
	return huffmanTreeHeader{
		eLen: byte(len(elems)),
		elem: elems,
	}
}

// buildHuffmanTree
func (hc *HuffmanCompression) buildHuffmanTree(datas []byte) (err error) {
	wdMap := map[int]huffmantree.WeightData{}
	for _, s := range datas {
		d := int(s)
		r, ok := wdMap[d]
		if !ok {
			r = huffmantree.WeightData{
				Data:   d,
				Weight: 0,
			}
		}
		r.Weight++
		wdMap[d] = r
	}
	wds := make([]huffmantree.WeightData, 0)
	for _, wd := range wdMap {
		wds = append(wds, wd)
	}
	if len(wds) == 0 {
		return errors.New("nil data")
	}
	hc.huffmanTree = huffmantree.Build(wds...)
	if hc.huffmanTree == nil {
		return errors.New("nil tree")
	}
	hc.codes = hc.huffmanTree.Codes()
	return nil
}

func (hc *HuffmanCompression) writeHuffmanDatas(datas []byte) error {
	bw := bitio.NewWriter(hc.w)
	defer bw.Close()
	for _, d := range datas {
		code := hc.codes[int(d)]
		r, err := strconv.ParseUint(code, 2, 64)
		if err != nil {
			return err
		}
		err = bw.WriteBits(r, uint8(len(code)))
		if err != nil {
			return err
		}
	}

	return nil
}

